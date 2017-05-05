#!/bin/bash

set -e

autoDesc=`git log -1 --pretty=oneline --abbrev-commit | cut -d " " -f2- | awk '{printf ("%s-%s-%s", $1,$2,$3)}' | sed 's/[^a-zA-Z0-9]/-/g'`
if [ -z "$1" ]
  then
    echo "Custom description not provided. Using auto description ${autoDesc}"
    desc=${autoDesc}
  else
    desc=$1
fi
br=`git rev-parse --abbrev-ref HEAD | sed 's/[^a-zA-Z0-9]/-/g'`
hash=`git log -1 --pretty=oneline --abbrev-commit | awk '{ print $1 }'`
dt=`date +"%m-%d-%Y-%H%M"`
tag=${br}-${dt}-${desc}-${hash}
if [ "$?" != "0" ]; then
  exit 1
fi

docker tag simplycredit/athenapdf:latest simplycredit/athenapdf:${tag}
if [ "$?" != "0" ]; then
  exit 1
fi
docker push simplycredit/athenapdf:${tag}
 
docker tag simplycredit/athenapdf-service:latest simplycredit/athenapdf-service:${tag}
if [ "$?" != "0" ]; then
  exit 1
fi
docker push simplycredit/athenapdf-service:${tag}
