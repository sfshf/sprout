#!/usr/bin/env bash

set -eo pipefail

# Build a govern app image using the specific Dockerfile.
docker build -t govern:latest -f app/govern/Dockerfile .
# Remove the old container named govern, if has.
docker rm -f govern
# Run a new container named govern.
docker run -itdp 8080:8080 --network sprout --name govern govern