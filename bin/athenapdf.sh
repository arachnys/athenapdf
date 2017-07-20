#!/bin/sh
#
# Run athenapdf cli in a container
# Arachnys <techteam@arachnys.com>
# https://github.com/arachnys/athenapdf/
#

set -e

VERSION="3"
IMAGE="arachnysdocker/athenapdf:$VERSION"

if [ "$(pwd)" != '/' ]; then
    VOLUMES="-v $(pwd):/conversions/"
fi

exec docker run --rm $VOLUMES -w "/conversions/" $IMAGE "$@"
