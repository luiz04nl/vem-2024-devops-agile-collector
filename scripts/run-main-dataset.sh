#!/usr/bin/env bash

SOURCE_ENV="${PWD}/.env" &&

. $SOURCE_ENV &&
export $(cut -d= -f1 $SOURCE_ENV) &&

export CGO_ENABLED=1

#gcc --version - tdm gcc https://jmeubank.github.io/tdm-gcc/download/
#sudo apt install build-essential
#brew install gcc

cd src/main/
go run .

#docker pull sonarqube
#docker run -d --name sonarqube -p 9000:9000 sonarqube
#http://localhost:9000
#docker-compose up -d

#Jdeodorant

# docker run -d --name sonarqube \
#   -p 9000:9000 \
#   -v sonarqube_data:/opt/sonarqube/data \
#   sonarqube
