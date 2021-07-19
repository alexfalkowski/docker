FROM circleci/golang:1.16-buster

USER root

# Install helm
RUN curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash

# Install yamllint
RUN apt-get update && apt-get install --no-install-recommends -y \
  python3 \
  python3-setuptools \
  python3-pip \
  yamllint \
  shellcheck \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/local/bin

# Install dockerize
ENV DOCKERIZE_VERSION v0.6.1
RUN curl -OL https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C . -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Install hadolint
RUN curl -OL https://github.com/hadolint/hadolint/releases/download/v2.5.0/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint

# Install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
  chmod +x kubectl

# Install helm chart-testing
RUN curl -OL https://github.com/helm/chart-testing/releases/download/v3.4.0/chart-testing_3.4.0_linux_amd64.tar.gz \
    && tar -C . -xzvf chart-testing_3.4.0_linux_amd64.tar.gz \
    && rm chart-testing_3.4.0_linux_amd64.tar.gz

# Add helm chart-testing files
RUN mkdir -p /usr/local/etc/ct
COPY chart_schema.yaml /usr/local/etc/ct
COPY lintconf.yaml /usr/local/etc/ct

USER circleci
WORKDIR /home/circleci/

# Install kind
RUN go get sigs.k8s.io/kind@v0.11.1

# Install diagrams
RUN pip3 install --no-cache-dir -U yamale==3.0.2 --user
