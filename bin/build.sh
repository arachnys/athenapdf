#!/bin/sh
#
# Builds athenapdf go assemblies
# Arachnys <techteam@arachnys.com>
# https://github.com/arachnys/athenapdf/
#

set -exo pipefail

if [ -z "$1" ] || [ -z "$2" ]; then
    echo "Build output or package not specified"
    exit 1
fi

CGO_ENABLED=0 \
    go build \
    -ldflags "-s" -a -installsuffix cgo \
    -o "$1" \
    "$2"
