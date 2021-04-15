# syntax=docker/dockerfile:1.2
FROM alpine as edts
ENV EDTS_COMMIT=fce284e2d6973f65937c586f29393d2323766366

WORKDIR /edts
RUN apk add --no-cache git
RUN git clone --depth=1 https://bitbucket.org/Esvandiary/edts.git .
RUN git checkout ${EDTS_COMMIT}
RUN rm .gitignore README.md
RUN rm -rf .git doc test

# ---
FROM node:14-alpine AS deps-common

WORKDIR /app
COPY ./package.json ./yarn.lock ./

# ---
FROM deps-common AS deps-dev
RUN yarn install --no-optional --frozen-lockfile && \
  yarn cache clean

# ---
FROM deps-common AS deps-prod
RUN yarn install --production=true --frozen-lockfile && \
  yarn cache clean

# ---
FROM node:14-alpine AS builder
WORKDIR /app

COPY . .
COPY --from=deps-dev /app/node_modules ./node_modules
RUN yarn build

# ---
FROM node:14-alpine AS runner

WORKDIR /app
ENV NODE_ENV production

RUN apk add --no-cache tini python3

COPY --from=edts /edts ./edts
COPY --from=deps-prod /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/build ./build

RUN addgroup -g 1001 -S nodejs && \
  adduser -S nodejs -u 1001 && \
  chown -R nodejs:nodejs /app/build && \
  chown -R nodejs:nodejs /app/edts

USER nodejs
EXPOSE 3000
VOLUME ["/app/edts/edtslib/data"]

ENTRYPOINT ["/sbin/tini", "--"]
CMD ["node", "."]
