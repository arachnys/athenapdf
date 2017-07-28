#!/bin/sh
#
# Tests athenapdf go packages
# Arachnys <techteam@arachnys.com>
# https://github.com/arachnys/athenapdf/
#

set -ex

go test \
    -v -tags non_integration \
    $(go list ./... | grep -v /vendor/)
