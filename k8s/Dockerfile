FROM cimg/go:1.18.4

USER root

# Install deps
RUN curl -s https://packagecloud.io/install/repositories/github/git-lfs/script.deb.sh | bash && \
    curl -s https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash && \
    apt-get update && apt-get install --no-install-recommends -y \
    git-lfs \
    python3 \
    python3-setuptools \
    python3-pip \
    yamllint \
    shellcheck \
    && rm -rf /var/lib/apt/lists/*

# Install lfs
RUN git lfs install

WORKDIR /usr/local/bin

# Install dockerize
ENV DOCKERIZE_VERSION v0.6.1
RUN curl -sSOL https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C . -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Install hadolint
RUN curl -sSOL https://github.com/hadolint/hadolint/releases/download/v2.10.0/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint

# Install kubectl
RUN curl -sSOL "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
  chmod +x kubectl

# Install minikube
RUN curl -sSOL https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && \
  mv minikube-linux-amd64 minikube && \
  chmod +x minikube

# Install helm chart-testing
RUN curl -sSOL https://github.com/helm/chart-testing/releases/download/v3.5.1/chart-testing_3.5.1_linux_amd64.tar.gz \
    && tar -C . -xzvf chart-testing_3.5.1_linux_amd64.tar.gz \
    && rm chart-testing_3.5.1_linux_amd64.tar.gz

# Install kubeaudit
RUN curl -sSOL https://github.com/Shopify/kubeaudit/releases/download/0.16.0/kubeaudit_0.16.0_linux_amd64.tar.gz \
    && tar -C . -xzvf kubeaudit_0.16.0_linux_amd64.tar.gz \
    && rm kubeaudit_0.16.0_linux_amd64.tar.gz

# Add helm chart-testing files
RUN mkdir -p /usr/local/etc/ct
COPY chart_schema.yaml /usr/local/etc/ct
COPY lintconf.yaml /usr/local/etc/ct

USER circleci
WORKDIR /home/circleci/

# Install deps
RUN go install sigs.k8s.io/kind@v0.14.0 && \
    go install github.com/controlplaneio/kubesec/v2@v2.11.4 && \
    pip3 install --no-cache-dir -U yamale==4.0.3 --user

# Install ruby.
RUN curl -sSL https://rvm.io/mpapis.asc | gpg --import - && \
  curl -sSL https://rvm.io/pkuczynski.asc | gpg --import - && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && \
  rvm requirements && \
  rvm install 3.0.4 && rvm use 3.0.4"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.0.4/bin:$PATH"

# Install gems.
RUN gem update bundler --force
