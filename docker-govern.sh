#!/usr/bin/env bash

set -eo pipefail

# Remove the old container named govern-mongo, if has.
docker rm -f govern-mongo
# Run a new container named govern-mongo.
docker run -itdp 27017:27017 \
--name govern-mongo \
-v govern-mongo-data-db:/data/db \
-v govern-mongo-data-configdb:/data/configdb \
--network sprout \
mongo

# Build a govern app builder image using the specific Dockerfile.
docker build \
--target builder \
-t builder:latest \
-f app/govern/Dockerfile .
# Build a govern app binary image using the specific Dockerfile.
docker build \
--target binary \
-t govern-binary:latest \
-f app/govern/Dockerfile .

# Remove the old container named govern, if has.
docker rm -f govern
# Run a new container named govern.
docker run -itd -p 8080:8080 \
-p 8000:8000 \
-p 8010:8010 \
-p 9000:9000 \
--network sprout \
--name govern \
govern-binary

