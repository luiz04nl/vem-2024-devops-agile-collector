# Dependencies
- docker
- docker-compose
- pydriller
- jq
- sonar-cli 6.1

## Ubuntu/Linux
```bash
sudo apt-get install python3-full golang-go python-is-python3 pip jq
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
sudo apt-get install unzip wget nodejs

### Install docker compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

### Install sonarqube cli
mkdir /tmp/sonarqube -p
cd /tmp/sonarqube
wget -O sonar-scanner-cli-6.1.0.4477-linux.zip https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-6.1.0.4477-linux-x64.zip?_gl=1*y4q9fp*_gcl_au*NDAzMzQ4NDgzLjE3MjI1NTk5MTY.*_ga*OTkzOTYxNjAxLjE3MjI1NTk5MTc.*_ga_9JZ0GZ5TC6*MTcyMjU1OTkxNi4xLjEuMTcyMjU2MTY2My40Ni4wLjA.
unzip sonar-scanner-cli-6.1.0.4477-linux.zip
mv sonar-scanner-6.1.0.4477-linux-x64/ /opt/sonar-scanner/
sudo ln -s /opt/sonar-scanner/sonar-scanner-6.1.0.4477-linux-x64/bin/sonar-scanner /usr/local/bin/sonar-scanner
```

# Research steps

## First Step - create the initial dataset sqlite

- Clone the research collector

```bash
git clone https://github.com/luiz04nl/devops-ic-collector.git
cd devops-ic-collector
```

- Set the GITHUB_ACCESS_TOKEN in .env file following the .env.example file

```bash
cp .env.example .env
```

- Get repositories from github via github graphql api requests

```bash
sh ./scripts/run-create-dataset.sh
```

That command got the repositories from github via github graphql api requests filtering by
repositories with more or equal 3224 star, witch was the average got from the paper "Understanding the Factors that Impact the Popularity of GitHubRepositories", in addition was not included in the graphql filter the keys "android, jvm, spring, platform_frameworks_base and hbase", this projects was considered out of scope due to the complexity of build or analyze, future works will include more projects.

Where found 500 repositories from graphql api response, which was used to create the initial dataset.

## Second Step - Clone Each Repository

```bash
sh ./scripts/run-clone-repos.sh
```

That command tries to clone each repository present in the dataset, marking wasCloned to 1 in case of success and keeping the wasCloned as 0 in case of some network or other eventual failure.

490 repositories were marketed like cloned and was considered in the next steps

## Third Step - Check the presence of devops configuration files

```bash
sh ./scripts/run-check-devops-and-tools.sh
```

That command goes to each repository market like wasCloned and for each one check the presence of files that suggest the use of devops like ".github/workflows/*", ".cicleci/config.yaml", "Jenkinsfile", ".gitlab-ci.yml", "azure-pipelines.yml", ".travis.yml", ".harness/*ya*" e "bitbucket-pipelines.yml", in case of any of these files the project was marketed as useDevOps.

Where found 331 with the suggestion of use of devops

## Third Step - Check the frequency of commits integrated in the default branch and the amount of contributors

```bash
sh ./scripts/run-check-agile-and-behaviors.sh
```

```bash
sh ./scripts/run-contributors.sh 
```

That command goes to each repository market like wasCloned and for each one checks the frequency of commits integrated in the default branch using the project pydriller. Repositories with average frequency more than 15 days were considered and markets like not agile useAgile = 0, and the other like useAgile = 1.

Where found 317 with the suggestion of use of agile

## Fourth Step - Build and extract data using sonarqube and docker

Prepare the docker container with sonar and postgres sql
```bash
sh ./scripts/run-prepare-quality-check.sh
```

Access the sonar url on http://localhost:9000/ with username and password admin
you will need to change the admin password, change to sonar because it is configured in scripts

Run the quality check
```bash
sh ./scripts/run-quality-check.sh
```

That command goes to each repository market like wasCloned and for each one builds the project and runs the sonar to get information, that step was the bottleneck due the complexity and time demand.

56 repositories where successfully built using 'maven', 'gradle' ou 'ant' and analyzed with sonar, future works will explore a larger dataset.
