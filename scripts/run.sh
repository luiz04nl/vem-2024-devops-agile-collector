#!/usr/bin/env bash

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

cd src/
go run .