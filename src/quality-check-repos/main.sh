#!/usr/bin/env bash

echo "The first argument is: $1"
REPOSITORY=$1
echo "REPOSITORY: $1"

sh ./build-and-scan.sh $REPOSITORY

sh ./retry-with-cache.sh $REPOSITORY