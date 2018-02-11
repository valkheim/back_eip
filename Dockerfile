FROM alpine:latest
MAINTAINER Charles Paulet <charles.paulet@epitech.eu>

COPY . /tmp/api/

RUN set -ex \
      && echo "@edge http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
      && apk --no-cache add musl-dev git go@edge \
      && cd /tmp/api \
      && export GOPATH=/tmp/api \
      && export GOBIN=/ \
      && git config --global http.https://gopkg.in.followRedirects true \
      && go get \
      && go build \
      && go install \
      && mkdir /auth \
      && mv /tmp/api/auth/server.rsa.crt /auth/server.crt \
      && mv /tmp/api/auth/server.rsa.key /auth/server.key \
      && rm -rf /tmp/api \
      && apk del git go

WORKDIR /

ENTRYPOINT ["/api"]
