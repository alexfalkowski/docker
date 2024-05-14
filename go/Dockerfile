FROM cimg/go:1.22.3

USER root
WORKDIR /usr/local/bin

# Install deps
RUN apt-get update && apt-get install --no-install-recommends -y \
    iproute2 \
    && rm -rf /var/lib/apt/lists/*

# Install dockerize.
ENV DOCKERIZE_VERSION v0.6.1
RUN curl -sSOL https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C . -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Install hadolint.
RUN curl -sSOL https://github.com/hadolint/hadolint/releases/download/v2.12.0/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint

# Install buf.
RUN curl -sSOL https://github.com/bufbuild/buf/releases/download/v1.31.0/buf-Linux-x86_64 && \
  mv buf-Linux-x86_64 buf && \
  chmod +x buf

# Install trivy.
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.51.1

# Install mkcert.
RUN curl -sJLO "https://dl.filippo.io/mkcert/latest?for=linux/amd64" && \
  mv mkcert-v*-linux-amd64 mkcert && \
  chmod +x mkcert

# Install doctl.
ENV DOCTL_VERSION 1.104.0
RUN curl -sSOL https://github.com/digitalocean/doctl/releases/download/v$DOCTL_VERSION/doctl-$DOCTL_VERSION-linux-amd64.tar.gz \
  && tar C . -xf doctl-$DOCTL_VERSION-linux-amd64.tar.gz \
  && rm doctl-$DOCTL_VERSION-linux-amd64.tar.gz

USER circleci
WORKDIR /home/circleci/

# Install gosec.
RUN curl -sSL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" v2.20.0

# Install go deps.
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1 && \
  go install github.com/alexfalkowski/gocovmerge@latest && \
  go install golang.org/x/vuln/cmd/govulncheck@latest && \
  go install gotest.tools/gotestsum@v1.11.0 && \
  go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest \
  && go clean --cache -testcache -fuzzcache -modcache

# Install ruby.
RUN curl -sSL https://rvm.io/mpapis.asc | gpg --import - && \
  curl -sSL https://rvm.io/pkuczynski.asc | gpg --import - && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && \
  rvm requirements && \
  rvm install 3.3.1 && rvm use 3.3.1"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.3.1/bin:$PATH"

# Install gems.
RUN gem update bundler --force
