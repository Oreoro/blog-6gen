# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements...

FROM golang:1.23-alpine AS golang-builder
LABEL maintainer="linkinstar@apache.org"

ARG GOPROXY

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PACKAGE github.com/oreoro/blog-6gen
ENV BUILD_DIR ${GOPATH}/src/${PACKAGE}
ENV ANSWER_MODULE ${BUILD_DIR}

ARG TAGS="sqlite sqlite_unlock_notify"
ENV TAGS "bindata timetzdata $TAGS"
ARG CGO_EXTRA_CFLAGS

# === Ensure your forked code is used ===
COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}

# Optional: Verification step to confirm local code use
RUN echo "==== Folder Listing ====" && ls -la ${BUILD_DIR} \
    && if [ -f "${BUILD_DIR}/test-fork.txt" ]; then cat ${BUILD_DIR}/test-fork.txt; fi

RUN apk --no-cache add build-base git bash nodejs npm && npm install -g pnpm@9.7.0 \
    && make clean build

RUN chmod 755 answer
RUN ["/bin/bash","-c","script/build_plugin.sh"]
RUN cp answer /usr/bin/answer

RUN mkdir -p /data/uploads && chmod 777 /data/uploads \
    && mkdir -p /data/i18n && cp -r i18n/*.yaml /data/i18n

FROM alpine
LABEL maintainer="linkinstar@apache.org"

ARG TIMEZONE
ENV TIMEZONE=${TIMEZONE:-"Asia/Shanghai"}

RUN apk update \
    && apk --no-cache add \
        bash \
        ca-certificates \
        curl \
        dumb-init \
        gettext \
        openssh \
        sqlite \
        gnupg \
        tzdata \
    && ln -sf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime \
    && echo "${TIMEZONE}" > /etc/timezone

COPY --from=golang-builder /usr/bin/answer /usr/bin/answer
COPY --from=golang-builder /data /data
COPY /script/entrypoint.sh /entrypoint.sh
RUN chmod 755 /entrypoint.sh

VOLUME /data
EXPOSE 80
ENTRYPOINT ["/entrypoint.sh"]
