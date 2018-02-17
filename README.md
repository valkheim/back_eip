# Backend Groomshop

Rest API in Golang

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/a1a7593063044011995d27c53372625c)](https://www.codacy.com/app/valkheim/back_eip?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=valkheim/back_eip&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/valkheim/back_eip)](https://goreportcard.com/report/github.com/valkheim/back_eip)

## Features
* SSLv2/TLS
* gorilla mux http router
* redis integration
* curl-like tests
* redis/certificates tools

## Getting started

You must install golang and redis at least ofc. You also must
generate server certificate and key. Retrieve also certificate (cacert.pem) to
run curl tests (or use -k option). You might be able to retrieve where it
belong on auth/server.crt or remotely by using :
```
$ openssl s_client -connect HOST:443 -showcerts
```

### Dockerize

Start docker deamon and use the compose file.
```
$ docker-compose up
```

### Local

You can start the service without docker.
To do so, add `api` and `store` as local hosts via `/etc/hosts`. Then, start
the `local.sh` script. It will init GroomShop api and start logging in files.
```
# local.sh && tail -f *.dump
```
Be sure to stop background process when needed:
```
# pkill eip && pkill redis-server
```
