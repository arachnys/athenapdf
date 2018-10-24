#!/bin/bash
set -e

# Block bad hosts, if /etc/hosts is writeable
if [ -w /etc/hosts ]; then
   cat conf/hosts >> /etc/hosts
fi

rm -f /tmp/.X99-lock
export DISPLAY=:99

exec "$@"
