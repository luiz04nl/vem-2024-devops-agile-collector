# Direct Dependencies
- docker
- docker-compose

# Other dependencies inside docker
- pydriller
- jq
- sonar-cli 6.1
- sonar
- go
- python

# Research steps

## First Step - create the initial dataset sqlite

- Clone the research collector

```bash
git clone https://github.com/luiz04nl/vem-2024-devops-agile-collector.git
cd vem-2024-devops-agile-collector
```

- Set your github access token as GITHUB_ACCESS_TOKEN inside the .env file following the .env.example file
- Set your absolute folder location as DOCKER_VOLUMES_BASE_PATH inside the .env file, it is necessary to mount a docker volume inside docker container

```bash
cp .env.example .env
```

## For Windows
- Install WSL2
https://learn.microsoft.com/pt-br/windows/wsl/install

-  Install Docker
https://docs.docker.com/desktop/install/windows-install/

## For Ubunutu / Linux
-  Install Docker
https://docs.docker.com/engine/install/ubuntu/

## For MAC
-  Install Docker
https://docs.docker.com/desktop/install/mac-install/

## For any OS
- Install Vscode
https://code.visualstudio.com/download

- Install Vscode Devcontainers
https://code.visualstudio.com/docs/devcontainers/containers

- It is recomends to use dev containers to speed up the research process, but the dependencies are declared in .devcontainer/debug.Dockerfile if other approach was choised

## First Step - Create the Initial Dataset

- Get repositories from github via github graphql api requests

```bash
sh ./scripts/run-create-dataset.sh
```

That command got the repositories from github via github graphql api requests filtering by
repositories with more or equal 3224 star, witch was the average got from the paper "Understanding the Factors that Impact the Popularity of GitHubRepositories", in addition was not included in the graphql filter the keys "android, jvm, spring, platform_frameworks_base and hbase", this projects was considered out of scope due to the complexity of build or analyze, future works will include more projects.

## Second Step - Clone Each Repository

```bash
sh ./scripts/run-clone-repos.sh
```

That command tries to clone each repository present in the dataset, marking wasCloned to 1 in case of success and keeping the wasCloned as 0 in case of some network or other eventual failure.

## Third Step - Check the presence of devops configuration files

```bash
sh ./scripts/run-check-devops-and-tools.sh
```

That command goes to each repository market like wasCloned and for each one check the presence of files that suggest the use of devops like ".github/workflows/*", ".cicleci/config.yaml", "Jenkinsfile", ".gitlab-ci.yml", "azure-pipelines.yml", ".travis.yml", ".harness/*ya*" e "bitbucket-pipelines.yml", in case of any of these files the project was marketed as useDevOps.

## Third Step - Check the frequency of commits integrated in the default branch and the amount of contributors

```bash
sh ./scripts/run-check-agile-and-behaviors.sh
```

```bash
sh ./scripts/run-contributors.sh
```

That command goes to each repository market like wasCloned and for each one checks the frequency of commits integrated in the default branch using the project pydriller. Repositories with average frequency more than 15 days were considered and markets like not agile useAgile = 0, and the other like useAgile = 1.

## Fourth Step - Build and extract data using sonarqube and docker
Access the sonar url on http://localhost:9000/ in your web browser.
Wait when SonarQube is starting.
Log in to SonarQube with username (login) admin and password admin (default username and password).
You will need to change the admin password, change to sonar because it is previously configured in scripts.
You can close the sonar webpage.

Run the quality check
```bash
sh ./scripts/run-quality-check.sh
```

That command goes to each repository market like wasCloned and for each one builds the project and runs the sonar to get information, that step was the bottleneck due the complexity and time demand.

Some repositories, but not all, where successfully built using 'maven', 'gradle' ou 'ant' and analyzed with sonar, future works will explore a larger dataset and other analytics tools.
