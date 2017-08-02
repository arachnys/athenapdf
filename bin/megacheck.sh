#!/bin/sh
#
# Runs linters, and static analysis on athenapdf go assembles, and packages
# Arachnys <techteam@arachnys.com>
# https://github.com/arachnys/athenapdf/
#

set -exo pipefail

megacheck \
    $(go list ./... | grep -v /vendor/)
