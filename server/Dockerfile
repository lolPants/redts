# syntax=docker/dockerfile:1.4
FROM node:16-alpine as base
FROM alpine as edts
ENV EDTS_COMMIT=e65543b8779e2c164dc1bff6e03f0286cc89e611

WORKDIR /edts
RUN apk add --no-cache git
RUN git clone --depth=1 https://bitbucket.org/Esvandiary/edts.git .
RUN git checkout ${EDTS_COMMIT}
RUN rm .gitignore README.md
RUN rm -rf .git doc test

# ---
FROM base AS deps

WORKDIR /app
COPY ./.yarn ./.yarn
COPY ./package.json ./yarn.lock ./.yarnrc.yml ./

RUN apk add --no-cache --virtual \
  build-deps \
  python3 \
  alpine-sdk \
  autoconf \
  libtool \
  automake

RUN yarn install --immutable --inline-builds

# ---
FROM base AS builder
WORKDIR /app

COPY . .
COPY --from=deps /app/.yarn ./.yarn
COPY --from=deps /app/node_modules ./node_modules
RUN yarn build

# ---
FROM base AS runner

WORKDIR /app
ENV NODE_ENV production

RUN apk add --no-cache tini python3

COPY --from=edts /edts ./edts
COPY --from=deps /app/.yarn ./.yarn
COPY --from=deps /app/node_modules ./node_modules
COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json ./package.json

RUN mkdir -p /app/edts/edtslib/data && \
  addgroup -g 1001 -S nodejs && \
  adduser -S nodejs -u 1001 && \
  chown -R nodejs:nodejs /app/build && \
  chown -R nodejs:nodejs /app/edts

USER nodejs
EXPOSE 3000
VOLUME ["/app/edts/edtslib/data"]

ARG GIT_REPO
LABEL org.opencontainers.image.source=${GIT_REPO}

ENTRYPOINT ["/sbin/tini", "--"]
CMD ["node", "."]
