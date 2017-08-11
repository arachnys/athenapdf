#!/bin/sh
#
# Tests athenapdf go packages
# Arachnys <techteam@arachnys.com>
# https://github.com/arachnys/athenapdf/
#

set -exo pipefail


if [ -z "$COVERALLS_TOKEN" ]; then
    go test \
        $(go list ./... | grep -v /vendor/) \
        -v -tags non_integration
else
    goveralls \
        -v -flags="-tags=non_integration" \
        -service=travis-ci
fi
