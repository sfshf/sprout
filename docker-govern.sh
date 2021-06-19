#!/usr/bin/env bash

#set -eo pipefail

echo '### Checking "docker" command ...'
docker info &> /dev/null
if [ $? -ne 0 ] ; then
  echo "bash: docker: command not found"
  exit 1
fi

echo '### Creating a new docker network named "sprout" ...'
docker network inspect sprout &> /dev/null
if [ $? -ne 0 ] ; then
  echo "create one docker network named sprout"
  docker network create sprout
fi

echo '### Run a new docker container named "govern-mongo" ...'
docker container inspect govern-mongo &> /dev/null
if [ $? -ne 0 ] ; then
  docker run -itdp 27017:27017 \
--name govern-mongo \
-v govern-mongo-data-db:/data/db \
-v govern-mongo-data-configdb:/data/configdb \
--network sprout \
mongo:latest
else
  docker start govern-mongo
fi

echo '### Run a new docker container named "govern-redis" ...'
docker container inspect govern-redis &> /dev/null
if [ $? -ne 0 ] ; then
  docker run -itdp 6379:6379 \
--name govern-redis \
-v govern-redis-data-db:/data/db \
-v govern-redis-data-configdb:/data/configdb \
--network sprout \
redis:latest
else
  docker start govern-redis
fi

echo '### Build a new docker image named "govern-builder:latest" ...'
docker image inspect govern-builder:latest &> /dev/null
if [ $? -eq 0 ] ; then
  docker rmi govern-builder:latest
fi
docker build \
--target builder \
-t govern-builder:latest \
-f app/govern/Dockerfile .

echo '### Build a new docker image named "govern-binary:latest" ...'
docker image inspect govern-binary:latest &> /dev/null
if [ $? -eq 0 ] ; then
  docker rmi govern-binary:latest
fi
docker build \
--target binary \
-t govern-binary:latest \
-f app/govern/Dockerfile .

echo '### Build a new docker container named "govern" ...'
docker container inspect govern &> /dev/null
if [ $? -eq 0 ] ; then
  docker rm -f govern
fi
docker run -itd -p 8080:8080 \
-p 8000:8000 \
-p 8010:8010 \
-p 9000:9000 \
--network sprout \
--name govern \
govern-binary:latest

