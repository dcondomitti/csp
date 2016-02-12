#!/bin/bash

TAG=`docker build . | tail -n1 | awk '{ print $3 }'`
docker tag -f ${TAG} quay.io/${ORGANIZATION}/csp:latest
docker push quay.io/${ORGANIZATION}/csp:latest
