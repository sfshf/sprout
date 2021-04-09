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

# Build a govern app image using the specific Dockerfile.
docker build -t govern:latest \
-f app/govern/Dockerfile .
# Remove the old container named govern, if has.
docker rm -f govern
# Run a new container named govern.
docker run -itd -p 8080:8080 \
-p 8090:8090 \
--network sprout \
--name govern \
govern

