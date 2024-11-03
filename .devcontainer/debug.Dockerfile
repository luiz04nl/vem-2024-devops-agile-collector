# Docker-outside-of-Docker (DoD)
FROM mcr.microsoft.com/vscode/devcontainers/base:ubuntu AS debug

WORKDIR /usr/src/app

ENV LANG en_US.utf8

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && DEBIAN_FRONTEND=noninteractive \
  && apt-get install -y locales tzdata \
  && rm -rf /var/lib/apt/lists/* \
  && localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8

RUN apt-get update \
  && apt-get upgrade -y \
  && apt-get install git bash -y  \
  && apt-get install jq wget unzip curl ca-certificates -y \
  && apt-get install python3 python-is-python3 pip python3.10-venv -y

RUN git config --global core.autocrlf false

# Instalar dependências para Docker
RUN apt-get install -y apt-transport-https ca-certificates curl software-properties-common \
    && curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - \
    && add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" \
    && apt-get update \
    && apt-get install -y docker-ce-cli

# Install docker compose
RUN curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
RUN chmod +x /usr/local/bin/docker-compose

### Install sonarqube cli
WORKDIR /tmp
RUN mkdir /tmp/sonarqube -p
RUN cd /tmp/sonarqube
RUN wget -O sonar-scanner-cli-6.1.0.4477-linux.zip https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-6.1.0.4477-linux-x64.zip?_gl=1*y4q9fp*_gcl_au*NDAzMzQ4NDgzLjE3MjI1NTk5MTY.*_ga*OTkzOTYxNjAxLjE3MjI1NTk5MTc.*_ga_9JZ0GZ5TC6*MTcyMjU1OTkxNi4xLjEuMTcyMjU2MTY2My40Ni4wLjA.
RUN unzip sonar-scanner-cli-6.1.0.4477-linux.zip
RUN mv sonar-scanner-6.1.0.4477-linux-x64/ /opt/sonar-scanner/
RUN ln -sf /opt/sonar-scanner/bin/sonar-scanner /usr/local/bin/sonar-scanner
WORKDIR /usr/src/app

# Install nodejs nvm
WORKDIR /tmp
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash
RUN . ~/.nvm/nvm.sh \
  && nvm install 20 \
  && nvm alias default 20 \
  && nvm use default
WORKDIR /usr/src/app

# Initial go
WORKDIR /tmp
RUN wget https://go.dev/dl/go1.23.1.linux-amd64.tar.gz
RUN rm -rf /usr/local/go
RUN tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go
WORKDIR /usr/src/app

# Prepare
RUN mkdir -p repos

# Adicionar o usuário padrão ao grupo docker (opcional, se necessário)
RUN groupadd docker \
    && usermod -aG docker vscode

# RUN apt-get install sudo -y
# RUN useradd -ms /bin/bash vscode
# RUN usermod -aG docker vscode
# RUN echo "vscode ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers
# RUN chown -R vscode:vscode /home/vscode
# WORKDIR /usr/src/app
# USER vscode

ENV SHELL /bin/bash