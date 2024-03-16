# #!/usr/bin/env bash
# set -e

# SOURCE_ENV="${PWD}/.env" &&

# . $SOURCE_ENV &&
# export $(cut -d= -f1 $SOURCE_ENV) &&

# export CGO_ENABLED=1

# cd src/jdeodorant-analysis

# docker build -t jdeodorant-analysis -f jdeodorant-analysis.Dockerfile .

# REPOSITORY="zfile-a8cd9760f60fe91b5a28435645bbf43a8e12ecf30a30f1bdf9d8dc6eb271f126"
# MSYS_NO_PATHCONV=1 docker run -it --rm -v $(pwd)/../../repos/:/repos/ jdeodorant-analysis $REPOSITORY

# # go run . > ../../out/out-jdeodorant-analysis.json 2> ../../out/out-jdeodorant-analysis.json


