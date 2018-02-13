FROM golang:1.9-alpine3.7
MAINTAINER Charles Paulet <charles.paulet@epitech.eu>

COPY . /tmp/api/

WORKDIR /tmp/api

RUN set -ex \
      && apk --no-cache add musl-dev git \
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
      && apk del git

ENTRYPOINT ["/api"]
