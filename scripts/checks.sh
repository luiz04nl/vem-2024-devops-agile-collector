#!/usr/bin/env bash

cat <<EOF > sonar-project.properties
sonar.projectKey=devops-ic-check
sonar.projectName=Devops IC Check
sonar.projectVersion=1.0
sonar.sources=.
sonar.sourceEncoding=UTF-8
EOF

# https://docs.sonarsource.com/sonarqube/9.9/analyzing-source-code/scanners/sonarscanner/

sonar-scanner --version

sonar-scanner