#!/usr/bin/env bash

REPOSITORY=$1
echo "C REPOSITORY: $REPOSITORY"

WITH_BUILD=${2:-false}
echo "C WITH_BUILD: $WITH_BUILD"

source ./extract-infos.sh $REPOSITORY $WITH_BUILD
# . ./extract-infos.sh $REPOSITORY $WITH_BUILD

echo "C projectType: $projectType"
echo "C projectTypeVersion: $projectTypeVersion"
echo "C analysisSuccess: $analysisSuccess"
echo "C filesAtRootDir: $filesAtRootDir"
echo "C extractInfosSuccess: $extractInfosSuccess"