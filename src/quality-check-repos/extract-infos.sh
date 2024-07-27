#!/usr/bin/env bash

REPOSITORY=$1
WITH_BUILD=${2:-true}
echo "B WITH_BUILD: $WITH_BUILD"

source ./build-and-scan.sh $REPOSITORY $WITH_BUILD
# . ./build-and-scan.sh $REPOSITORY $WITH_BUILD

echo "B projectType: $projectType"
echo "B projectTypeVersion: $projectTypeVersion"
echo "B analysisSuccess: $analysisSuccess"
echo "B filesAtRootDir: $filesAtRootDir"

export projectType=$projectType
export projectTypeVersion=$projectTypeVersion
export analysisSuccess=$analysisSuccess
export filesAtRootDir=$filesAtRootDir

echo "AAA Path:"
pwd

cd ../../repos
cd $REPOSITORY
# git submodule update --init --recursive
mkdir -p ../../out/quality-check-repos

echo "BBB Path:"
pwd

# SONAR_USERNAME="admin"
# SONAR_PASSWORD="sonar"
# SONAR_CREDENTIALS="${SONAR_USERNAME}:${SONAR_PASSWORD}"
# ENCODED_CREDENTIALS=$(echo -n "$SONAR_CREDENTIALS" | base64)

# echo "ENCODED_CREDENTIALS: $ENCODED_CREDENTIALS"

# SONAR_URL="http://localhost:9000"
pageSize=500

# currentPage=1
# curl -s "${SONAR_URL}/api/issues/search?projects=${REPOSITORY}&statuses=OPEN&ps=${pageSize}&p=1" --header "Authorization: Basic $ENCODED_CREDENTIALS" \
# --header "Authorization: Basic $ENCODED_CREDENTIALS" \
# > ../../out/quality-check-repos/$REPOSITORY-ISSUES-page-${currentPage}.json

currentPage=1
totalPages=2 # apenas para executar a primeira vez
while [ $currentPage -le $totalPages ]
do
    response=$(curl -s "${SONAR_URL}/api/issues/search?projects=${REPOSITORY}&statuses=OPEN&ps=${pageSize}&p=$currentPage" --header "Authorization: Basic $ENCODED_CREDENTIALS")
    if [ $(echo $response | jq '.errors') != "null" ]; then
        echo "Erro ao buscar dados: $(echo $response | jq '.errors')"
        exit 1
    fi
    echo $response | jq '.' > ../../out/quality-check-repos/$REPOSITORY-ISSUES-page-${currentPage}.json
    totalPages=$(echo $response | jq '.paging.total')
    currentPage=$((currentPage + 1))
done
# echo "############################"


# echo "######### Geting Code Smells ###################"
currentPage=1
curl -s "${SONAR_URL}/api/issues/search?projects=${REPOSITORY}&statuses=OPEN&types=CODE_SMELL&ps=${pageSize}&p=1" --header "Authorization: Basic $ENCODED_CREDENTIALS" \
--header "Authorization: Basic $ENCODED_CREDENTIALS" \
> ../../out/quality-check-repos/$REPOSITORY-CODE_SMELL-page-${currentPage}.json

# currentPage=1
# totalPages=2 # apenas para executar a primeira vez
# while [ $currentPage -le $totalPages ]
# do
#     response=$(curl -s "${SONAR_URL}/api/issues/search?projects=${REPOSITORY}&statuses=OPEN&types=CODE_SMELL&ps=${pageSize}&p=$currentPage" --header "Authorization: Basic $ENCODED_CREDENTIALS")
#     if [ $(echo $response | jq '.errors') != "null" ]; then
#         echo "Erro ao buscar dados: $(echo $response | jq '.errors')"
#         exit 1
#     fi
#     echo $response | jq '.' > ../../out/quality-check-repos/$REPOSITORY-CODE_SMELL-page-${currentPage}.json
#     totalPages=$(echo $response | jq '.paging.total')
#     currentPage=$((currentPage + 1))
# done
# echo "############################"

requestSonarMeasures() {
  METRIC=$1
  curl -s "${SONAR_URL}/api/measures/component?component=${REPOSITORY}&metricKeys=$METRIC" \
--header "Authorization: Basic $ENCODED_CREDENTIALS" \
> ../../out/quality-check-repos/$REPOSITORY-$METRIC.json
}

requestSonarMeasures "sqale_rating"
#A=0-0.05, B=0.06-0.1, C=0.11-0.20, D=0.21-0.5, E=0.51-1

requestSonarMeasures "reliability_rating"
# A = 0 Bugs
# B = at least 1 Minor Bug
# C = at least 1 Major Bug
# D = at least 1 Critical Bug
# E = at least 1 Blocker Bug

requestSonarMeasures "complexity"
requestSonarMeasures "cognitive_complexity"
requestSonarMeasures "duplicated_blocks"
requestSonarMeasures "duplicated_files"
requestSonarMeasures "duplicated_lines"
requestSonarMeasures "code_smells"
requestSonarMeasures "ncloc"
requestSonarMeasures "sqale_index"
requestSonarMeasures "sqale_debt_ratio"
requestSonarMeasures "quality_gate_details"
requestSonarMeasures "bugs"
requestSonarMeasures "vulnerabilities"
requestSonarMeasures "security_rating"
requestSonarMeasures "classes"
requestSonarMeasures "comment_lines"
requestSonarMeasures "coverage"
requestSonarMeasures "tests"

json=$(jq -n \
  --arg projectType "$projectType" \
  --arg projectTypeVersion "$projectTypeVersion" \
  --arg REPOSITORY "$REPOSITORY" \
  --arg analysisSuccess "$analysisSuccess" \
  --arg filesAtRootDir "$filesAtRootDir" \
  '{projectType: $projectType, projectTypeVersion: $projectTypeVersion, repository: $REPOSITORY, analysisSuccess, $analysisSuccess, filesAtRootDir: $filesAtRootDir'})
echo  $json > ../../out/quality-check-repos/$REPOSITORY.json

export extractInfosSuccess=1

sleep 5