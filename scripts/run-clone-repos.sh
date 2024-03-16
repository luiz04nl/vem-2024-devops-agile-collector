#!/usr/bin/env bash
set -e

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

cd src/clone-repos

git config --global core.longpaths true
git config --global http.postBuffer 524288000
git config --global http.lowSpeedLimit 0
git config --global http.lowSpeedTime 999999
git config --global core.protectNTFS false
git config --global http.version HTTP/2

go run . > ../../out/out-clone-repos.json 2> ../../out/out-clone-repos.json

