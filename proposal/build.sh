#!/usr/bin/env bash
# To be run as a bootstrap for the microservice on each machine.

# build docker container
cd demo-app
bash build-and-run.sh || (\
    echo "============="
    echo "Build failed!"
    echo "============="
    exit 1)
