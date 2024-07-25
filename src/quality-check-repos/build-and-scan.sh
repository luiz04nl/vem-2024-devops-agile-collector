#!/usr/bin/env bash

echo "The first argument is: $1"
REPOSITORY=$1
echo "REPOSITORY: $1"

cd ../../repos
cd $REPOSITORY
git submodule update --init --recursive

SONAR_USERNAME="admin"
SONAR_PASSWORD="sonar"
SONA_CREDENTIALS="${SONAR_USERNAME}:${SONAR_PASSWORD}"
ENCODED_CREDENTIALS=$(echo -n "$SONA_CREDENTIALS" | base64)

echo "ENCODED_CREDENTIALS: $ENCODED_CREDENTIALS"

cat <<EOF > sonar-project.properties
sonar.projectKey=$REPOSITORY
sonar.projectName=$REPOSITORY
sonar.projectVersion=1.0
sonar.sourceEncoding=UTF-8
sonar.host.url=http://localhost:9000
sonar.login=$SONAR_USERNAME
sonar.password=$SONAR_PASSWORD
EOF

projectTypeVersion=""
projectType=""
analysisSuccess=undefined
filesAtRootDir=$(ls .)

echo "####### Building project #####################"

mkdir -p target/classes
mkdir -p target/test-classes

if [ -f "pom.xml" ]; then
  cat <<EOF >> sonar-project.properties
sonar.language=java
sonar.java.binaries=target/classes,target/test-classes
EOF

  projectType="maven"
  distributionUrl=`cat .mvn/wrapper/maven-wrapper.properties | grep distributionUrl`
  splitA=(${distributionUrl//"distributionUrl=https://repo.maven.apache.org/maven2/org/apache/maven/apache-maven/"/ })
  splitB=(${splitA//"-all.zip"/ })
  splitC=(${splitB//"-bin.zip"/ })

  projectTypeVersion=(${splitC//"/apache-maven-"/ })

  defaultProjectTypeVersion="3.9.6-eclipse-temurin-11"

  if [ -z "$projectTypeVersion" ]; then
      projectTypeVersion=$defaultProjectTypeVersion
  fi
  image_exists() {
    docker image inspect "$1" > /dev/null 2>&1
  }
  if ! image_exists "$projectTypeVersion"
  then
    projectTypeVersion=$defaultProjectTypeVersion
  fi
  echo "projectTypeVersion: "$projectTypeVersion

  MSYS_NO_PATHCONV=1 docker run --rm -v "$(pwd)":/usr/src/mymaven -w /usr/src/mymaven maven:$projectTypeVersion /bin/bash -c "git config --global --add safe.directory /usr/src/mymaven && mvn compile" 2>&1

elif [ -f "build.gradle" ]; then
  mkdir -p build/classes/java/main
  mkdir -p build/classes/java/test

  cat <<EOF >> sonar-project.properties
sonar.language=java
sonar.java.binaries=build/classes/java/main,build/classes/java/test
EOF

  projectType="gradle"
  if [ -f "app/src/main/AndroidManifest.xml" ]; then
    projectType="gradle-android"
  fi

  distributionUrl=`cat gradle/wrapper/gradle-wrapper.properties  | grep distributionUrl`
  splitA=(${distributionUrl//"distributionUrl=https\://services.gradle.org/distributions/gradle-"/ })
  splitB=(${splitA//"-all.zip"/ })
  projectTypeVersion=(${splitB//"-bin.zip"/ })

  defaultProjectTypeVersion="8.7.0-jdk8"
  # defaultProjectTypeVersion="7.6.4-jdk8"
  # defaultProjectTypeVersion="6.9.4-jdk8"
  # defaultProjectTypeVersion="latest"

  if [ -z "$projectTypeVersion" ]; then
      projectTypeVersion=$defaultProjectTypeVersion
  fi
  image_exists() {
    docker image inspect "$1" > /dev/null 2>&1
  }
  if ! image_exists "$projectTypeVersion"
  then
    projectTypeVersion=$defaultProjectTypeVersion
  fi
  echo "projectTypeVersion: "$projectTypeVersion

  MSYS_NO_PATHCONV=1 docker run --rm -u gradle -v `pwd`:/home/gradle/project -w /home/gradle/project gradle:$projectTypeVersion /bin/bash -c "git config --global --add safe.directory /home/gradle/project && gradle build" 2>&1

elif  [ -f "build.xml" ]; then
  cat <<EOF >> sonar-project.properties
sonar.language=java
sonar.java.binaries=**/**/classes
EOF

  projectType="ant"
  distributionUrl=`cat build.xml | grep antversion | xargs`
  splitA=(${distributionUrl/"<antversion atleast=\""/ })
  projectTypeVersion=(${splitA//"\""/ })

  defaultProjectTypeVersion="latest"

  if [ -z "$projectTypeVersion" ]; then
      projectTypeVersion=$defaultProjectTypeVersion
  fi
  image_exists() {
    docker image inspect "$1" > /dev/null 2>&1
  }
  if ! image_exists "$projectTypeVersion"
  then
    projectTypeVersion=$defaultProjectTypeVersion
  fi
  echo "projectTypeVersion: "$projectTypeVersion

  MSYS_NO_PATHCONV=1 docker run --rm -v "$(pwd)":/workspace -w /workspace bitnami/ant:$projectTypeVersion /bin/bash -c "git config --global --add safe.directory/workspace && ant build" 2>&1

elif  [ -f "package.json" ]; then
  if [ -f "tsconfig.json" ]; then
    cat <<EOF >> sonar-project.properties
sonar.language=TypeScript
EOF
  else
    cat <<EOF >> sonar-project.properties
sonar.language=JavaScript
EOF
  fi

  projectType="nodejs"
  projectTypeVersion=""
else
  projectType="other"
  projectTypeVersion=""

  if find . -name "*.java" -not -name "test.java" -print -quit | grep -q '.'; then
    echo "Arquivos .java encontrados."

    cat <<EOF >> sonar-project.properties
sonar.language=java
sonar.java.binaries=target/classes/
EOF
#sonar.java.binaries=target/classes,target/test-classes

#########################################################
    projectType="undefined"
    projectTypeVersion="undefined"

    cat <<EOF > custom-pom.xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.devops</groupId>
    <artifactId>test</artifactId>
    <version>1.0-SNAPSHOT</version>

    <properties>
        <maven.compiler.source>1.8</maven.compiler.source>
        <maven.compiler.target>1.8</maven.compiler.target>
    </properties>

    <dependencies>
    </dependencies>

    <build>
        <plugins>
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-compiler-plugin</artifactId>
                <version>3.8.1</version>
                <configuration>
                    <source>\${maven.compiler.source}</source>
                    <target>\${maven.compiler.target}</target>
                </configuration>
            </plugin>
            <plugin>
                <groupId>org.sonarsource.scanner.maven</groupId>
                <artifactId>sonar-maven-plugin</artifactId>
                <version>3.7.0.1746</version>
            </plugin>
        </plugins>
    </build>
</project>
EOF

    srcDir="src/main/java"
    if [ -d "$srcDir" ]; then
        rm -r $srcDir
    fi
    mkdir -p $srcDir

    find . -name "*.java" -exec cp {} $srcDir/ \;

  mavenVersion="3.8.1"
  MSYS_NO_PATHCONV=1 docker run --rm -v "$(pwd)":/usr/src/mymaven -w /usr/src/mymaven maven:$mavenVersion /bin/bash -c "git config --global --add safe.directory /usr/src/mymaven && mvn -f custom-pom.xml compile" 2>&1

##########################################################
    # projectType="javac"
    # projectTypeVersion=""

    # mkdir -p target/classes
    # find . -name "*.java" > sources.txt
    # # javac -d target/classes @sources.txt
    ## MSYS_NO_PATHCONV=1 docker run --rm -v "$(pwd)":/usr/src/app -w /usr/src/app eclipse-temurin:18 javac -d target/classes @sources.txt 2>&1

  else
    projectType="other"
    projectTypeVersion=""
  fi
fi

echo "projectType: " $projectType
echo "projectTypeVersion: "$projectTypeVersion

sleep 10

echo "####### Scaning with sonar #####################"
sonar-scanner.bat --version
sonar-scanner.bat -X  2>&1

# echo "File .scannerwork/report-task.txt"
# cat .scannerwork/report-task.txt

# echo "####### Geting Sonar Execution Success #####################"
# analysisSuccess="undefined"
# if grep -q "EXECUTION SUCCESS" ../../out/quality-check-repos/$REPOSITORY.out.txt; then
#     analysisSuccess=true
# else
#     analysisSuccess=false
# fi
# echo "############################"
# echo "analysisSuccess: "$analysisSuccess

# echo "############################"

# mkdir -p ../../out/quality-check-repos

# SONAR_URL="http://localhost:9000"
# pageSize=500

# sleep 10

# echo "####### Geting Issues #####################"
# # currentPage=1
# # curl -s "${SONAR_URL}/api/issues/search?projects=${REPOSITORY}&statuses=OPEN&ps=${pageSize}&p=1" --header "Authorization: Basic $ENCODED_CREDENTIALS" \
# # --header "Authorization: Basic $ENCODED_CREDENTIALS" \
# # > ../../out/quality-check-repos/$REPOSITORY-ISSUES-page-${currentPage}.json

# currentPage=1
# totalPages=2 # apenas para executar a primeira vez
# while [ $currentPage -le $totalPages ]
# do
#     response=$(curl -s "${SONAR_URL}/api/issues/search?projects=${REPOSITORY}&statuses=OPEN&ps=${pageSize}&p=$currentPage" --header "Authorization: Basic $ENCODED_CREDENTIALS")
#     if [ $(echo $response | ../../jq '.errors') != "null" ]; then
#         echo "Erro ao buscar dados: $(echo $response | ../../jq '.errors')"
#         exit 1
#     fi
#     echo $response | ../../jq '.' > ../../out/quality-check-repos/$REPOSITORY-ISSUES-page-${currentPage}.json
#     totalPages=$(echo $response | ../../jq '.paging.total')
#     currentPage=$((currentPage + 1))
# done
# echo "############################"


# echo "######### Geting Code Smells ###################"
# currentPage=1
# curl -s "${SONAR_URL}/api/issues/search?projects=${REPOSITORY}&statuses=OPEN&types=CODE_SMELL&ps=${pageSize}&p=1" --header "Authorization: Basic $ENCODED_CREDENTIALS" \
# --header "Authorization: Basic $ENCODED_CREDENTIALS" \
# > ../../out/quality-check-repos/$REPOSITORY-CODE_SMELL-page-${currentPage}.json

# # currentPage=1
# # totalPages=2 # apenas para executar a primeira vez
# # while [ $currentPage -le $totalPages ]
# # do
# #     response=$(curl -s "${SONAR_URL}/api/issues/search?projects=${REPOSITORY}&statuses=OPEN&types=CODE_SMELL&ps=${pageSize}&p=$currentPage" --header "Authorization: Basic $ENCODED_CREDENTIALS")
# #     if [ $(echo $response | ../../jq '.errors') != "null" ]; then
# #         echo "Erro ao buscar dados: $(echo $response | ../../jq '.errors')"
# #         exit 1
# #     fi
# #     echo $response | ../../jq '.' > ../../out/quality-check-repos/$REPOSITORY-CODE_SMELL-page-${currentPage}.json
# #     totalPages=$(echo $response | ../../jq '.paging.total')
# #     currentPage=$((currentPage + 1))
# # done
# echo "############################"

# requestSonarMeasures() {
#   METRIC=$1
#   curl -s "${SONAR_URL}/api/measures/component?component=${REPOSITORY}&metricKeys=$METRIC" \
# --header "Authorization: Basic $ENCODED_CREDENTIALS" \
# > ../../out/quality-check-repos/$REPOSITORY-$METRIC.json
# }

# requestSonarMeasures "sqale_rating"
# #A=0-0.05, B=0.06-0.1, C=0.11-0.20, D=0.21-0.5, E=0.51-1

# requestSonarMeasures "reliability_rating"
# # A = 0 Bugs
# # B = at least 1 Minor Bug
# # C = at least 1 Major Bug
# # D = at least 1 Critical Bug
# # E = at least 1 Blocker Bug

# requestSonarMeasures "complexity"
# requestSonarMeasures "cognitive_complexity"
# requestSonarMeasures "duplicated_blocks"
# requestSonarMeasures "duplicated_files"
# requestSonarMeasures "duplicated_lines"
# requestSonarMeasures "code_smells"
# requestSonarMeasures "ncloc"
# requestSonarMeasures "sqale_index"
# requestSonarMeasures "sqale_debt_ratio"
# requestSonarMeasures "quality_gate_details"
# requestSonarMeasures "bugs"
# requestSonarMeasures "vulnerabilities"
# requestSonarMeasures "security_rating"
# requestSonarMeasures "classes"
# requestSonarMeasures "comment_lines"
# requestSonarMeasures "coverage"
# requestSonarMeasures "tests"

# echo "######### Geting Other Infos ###################"
# json=$(../../jq -n \
#   --arg projectType "$projectType" \
#   --arg projectTypeVersion "$projectTypeVersion" \
#   --arg REPOSITORY "$REPOSITORY" \
#   --arg analysisSuccess "$analysisSuccess" \
#   --arg filesAtRootDir "$filesAtRootDir" \
#   '{projectType: $projectType, projectTypeVersion: $projectTypeVersion, repository: $REPOSITORY, analysisSuccess, $analysisSuccess, filesAtRootDir: $filesAtRootDir'})
# echo  $json > ../../out/quality-check-repos/$REPOSITORY.json
# echo "############################"

# sleep 10

# echo "FINISHED"