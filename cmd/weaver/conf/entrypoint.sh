#!/bin/bash
set -e

# Block bad hosts
cat conf/hosts >> /etc/hosts

exec "$@"
