#!/usr/bin/env bash

set -e

# Version control
commit_hash=$(git rev-parse HEAD)
commit_hash_short=$(git rev-parse --short HEAD)

# Set arch variable based on the architecture
architecture=$(uname -m)
if [ "$architecture" == "x86_64" ]; then
    export arch="x86_64" # Linux, Windows
else
    export arch="aarch64" # Mac
fi

docker build \
       --build-arg ARCH="$arch" \
       --build-arg BUILD_DATE="$(git show -s --format=%ci "$commit_hash")"\
       --build-arg SERVICE=mpcfusionmodel \
       --build-arg GIT_SHA="$commit_hash" \
       -t "${ECR}"mpcfusionmodel:latest  \
       -t "${ECR}"mpcfusionmodel:"$commit_hash_short"  \
       -f Dockerfile-mpcfusionmodel ..