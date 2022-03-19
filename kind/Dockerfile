FROM cimg/go:1.18

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
RUN curl -OL https://github.com/hadolint/hadolint/releases/download/v2.8.0/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint

# Install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
  chmod +x kubectl

# Install helm chart-testing
RUN curl -OL https://github.com/helm/chart-testing/releases/download/v3.5.1/chart-testing_3.5.1_linux_amd64.tar.gz \
    && tar -C . -xzvf chart-testing_3.5.1_linux_amd64.tar.gz \
    && rm chart-testing_3.5.1_linux_amd64.tar.gz

# Add helm chart-testing files
RUN mkdir -p /usr/local/etc/ct
COPY chart_schema.yaml /usr/local/etc/ct
COPY lintconf.yaml /usr/local/etc/ct

USER circleci
WORKDIR /home/circleci/

# Install deps
RUN go install sigs.k8s.io/kind@v0.12.0 && \
    pip3 install --no-cache-dir -U yamale==4.0.3 --user
