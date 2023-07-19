FROM cimg/go:1.20.6

USER root

# Install deps.
RUN apt-get update && apt-get install --no-install-recommends -y \
    gnupg \
    python3-pip \
    software-properties-common \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/local/bin

# Install terragrunt.
RUN curl -sSOL https://github.com/gruntwork-io/terragrunt/releases/download/v0.45.15/terragrunt_linux_amd64 && \
  mv terragrunt_linux_amd64 terragrunt && \
  chmod +x terragrunt

# Install driftctl.
RUN curl -sSOL https://github.com/snyk/driftctl/releases/latest/download/driftctl_linux_amd64 && \
  mv driftctl_linux_amd64 driftctl && \
  chmod +x driftctl

# Install infracost.
RUN curl -fsSL https://raw.githubusercontent.com/infracost/infracost/master/scripts/install.sh | sh

# Install regula.
RUN curl -sSOL https://github.com/fugue/regula/releases/download/v3.2.1/regula_3.2.1_Linux_x86_64.tar.gz \
  && tar -C . -xzvf regula_3.2.1_Linux_x86_64.tar.gz \
  && rm regula_3.2.1_Linux_x86_64.tar.gz

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
RUN curl -sSOL https://github.com/mikefarah/yq/releases/download/4.34.2/yq_linux_amd64 && \
  mv yq_linux_amd64 yq && \
  chmod +x yq

USER circleci
WORKDIR /home/circleci/

# Install tfenv.
RUN git clone --depth 1 --branch v3.0.0 https://github.com/tfutils/tfenv.git /home/circleci/.tfenv
ENV PATH "/home/circleci/.tfenv/bin:$PATH"

# Install terraform.
RUN tfenv use 1.4.6 && \
    pip3 install --no-cache-dir -U checkov==2.3.247 terraform-local==0.8 --user

# Install ruby.
RUN curl -sSL https://rvm.io/mpapis.asc | gpg --import - && \
  curl -sSL https://rvm.io/pkuczynski.asc | gpg --import - && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && \
  rvm requirements && \
  rvm install 3.1.4 && rvm use 3.1.4"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.1.4/bin:$PATH"

# Install gems.
RUN gem update bundler --force
