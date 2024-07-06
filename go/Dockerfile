FROM alexfalkowski/root:1.2

USER root
WORKDIR /usr/local/bin

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
RUN curl -sSOL https://github.com/bufbuild/buf/releases/download/v1.34.0/buf-Linux-x86_64 && \
  mv buf-Linux-x86_64 buf && \
  chmod +x buf

# Install trivy.
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.53.0

# Install mkcert.
RUN curl -sJLO "https://dl.filippo.io/mkcert/latest?for=linux/amd64" && \
  mv mkcert-v*-linux-amd64 mkcert && \
  chmod +x mkcert

USER circleci
WORKDIR /home/circleci/

# Install gosec.
RUN curl -sSL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" v2.20.0

# Install go tools.
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1 && \
  go install github.com/alexfalkowski/gocovmerge@latest && \
  go install golang.org/x/vuln/cmd/govulncheck@v1.1.2 && \
  go install gotest.tools/gotestsum@v1.12.0 && \
  go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@v0.22.0 && \
  go install github.com/Zxilly/go-size-analyzer/cmd/gsa@v1.4.0 \
  && go clean --cache -testcache -fuzzcache -modcache && rm -rf "${GOPATH}/pkg"
