#!/bin/bash
set -e

# Block bad hosts
cat conf/hosts >> /etc/hosts

rm -f /tmp/.X99-lock
export DISPLAY=:99

exec "$@"
