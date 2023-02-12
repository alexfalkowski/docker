FROM cimg/go:1.20.0

USER root

RUN apt-get update && apt-get install --no-install-recommends -y \
    gnupg \
    python3-pip \
    software-properties-common \
    && rm -rf /var/lib/apt/lists/*

RUN curl -sSL https://apt.releases.hashicorp.com/gpg | gpg --dearmor | tee /usr/share/keyrings/hashicorp-archive-keyring.gpg && \
    gpg --no-default-keyring --keyring /usr/share/keyrings/hashicorp-archive-keyring.gpg --fingerprint && \
    echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/hashicorp.list && \
    apt-get update && apt-get install --no-install-recommends -y \
      terraform \
      && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/local/bin

# Install pips.
RUN pip3 install --no-cache-dir -U checkov==2.3.3 terraform-local==0.6 --user

# Install terragrunt.
RUN curl -sSOL https://github.com/gruntwork-io/terragrunt/releases/download/v0.43.2/terragrunt_linux_amd64 && \
  mv terragrunt_linux_amd64 terragrunt && \
  chmod +x terragrunt

# Install driftctl.
RUN curl -sSOL https://github.com/snyk/driftctl/releases/latest/download/driftctl_linux_amd64 && \
  mv driftctl_linux_amd64 driftctl && \
  chmod +x driftctl

# Install infracost.
RUN curl -fsSL https://raw.githubusercontent.com/infracost/infracost/master/scripts/install.sh | sh

# Install dockerize.
ENV DOCKERIZE_VERSION v0.6.1
RUN curl -sSOL https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C . -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Install hadolint.
RUN curl -sSOL https://github.com/hadolint/hadolint/releases/download/v2.12.0/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint

# Install jq.
RUN curl -sSOL https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && \
  mv jq-linux64 jq && \
  chmod +x jq

# Install yq.
RUN curl -sSOL https://github.com/mikefarah/yq/releases/download/4.30.8/yq_linux_amd64 && \
  mv yq_linux_amd64 yq && \
  chmod +x yq

USER circleci
WORKDIR /home/circleci/

# Install ruby.
RUN curl -sSL https://rvm.io/mpapis.asc | gpg --import - && \
  curl -sSL https://rvm.io/pkuczynski.asc | gpg --import - && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && \
  rvm requirements && \
  rvm install 3.1.3 && rvm use 3.1.3"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.1.3/bin:$PATH"

# Install gems.
RUN gem update bundler --force
