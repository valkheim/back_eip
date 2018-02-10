#!/bin/sh

# docker stop groomshop-ctnr
# docker container rm groomshop-ctnr
# docker rmi groomshop-img
# #create docker image
# docker build -t groomshop-img -f ./Dockerfile .
# #listing docker images
# docker images
# #run docker image as a container
# docker run -d --name groomshop-ctnr groomshop-img
# #listing docker containers
# docker ps
docker-compose up
