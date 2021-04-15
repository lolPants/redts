# syntax=docker/dockerfile:1.2
FROM alpine as edts-clone
ENV EDTS_COMMIT=fce284e2d6973f65937c586f29393d2323766366

WORKDIR /edts
RUN apk add --no-cache git
RUN git clone --depth=1 https://bitbucket.org/Esvandiary/edts.git .
RUN git checkout ${EDTS_COMMIT}
RUN rm .gitignore README.md
RUN rm -rf .git doc test
