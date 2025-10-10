FROM alexfalkowski/root:1.51

USER root
WORKDIR /usr/local/bin

# Install dockerize.
ENV DOCKERIZE_VERSION=v0.9.3
RUN export machine="$(dpkg --print-architecture)"; curl -sSOL https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-$machine-$DOCKERIZE_VERSION.tar.gz \
  && tar -C . -xzvf dockerize-linux-$machine-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-linux-$machine-$DOCKERIZE_VERSION.tar.gz

# Install shellcheck.
ENV SHELLCHECK_VERSION=v0.10.0
RUN export machine="$(uname -m)"; curl -sSOL https://github.com/koalaman/shellcheck/releases/download/$SHELLCHECK_VERSION/shellcheck-$SHELLCHECK_VERSION.linux.$machine.tar.xz \
  && tar -C . -xvf shellcheck-$SHELLCHECK_VERSION.linux.$machine.tar.xz \
  && mv shellcheck-$SHELLCHECK_VERSION/shellcheck . \
  && rm -rf shellcheck-$SHELLCHECK_VERSION/ shellcheck-$SHELLCHECK_VERSION.linux.$machine.tar.xz

# Install hadolint.
RUN export machine="$(uname -m)"; curl -sSOL https://github.com/hadolint/hadolint/releases/download/v2.14.0/hadolint-Linux-$machine \
  && mv hadolint-Linux-$machine hadolint \
  && chmod +x hadolint

# Install buf.
RUN export machine="$(uname -m)"; curl -sSOL https://github.com/bufbuild/buf/releases/download/v1.58.0/buf-Linux-$machine \
  && mv buf-Linux-$machine buf \
  && chmod +x buf

# Install trivy.
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.67.1

# Install mkcert.
RUN export machine="$(dpkg --print-architecture)"; curl -sJLO "https://dl.filippo.io/mkcert/latest?for=linux/$machine" \
  && mv mkcert-v*-linux-$machine mkcert \
  && chmod +x mkcert

# Install codecov.
RUN curl -sSOL https://cli.codecov.io/v11.2.3/linux/codecov \
  && mv codecov codecovcli \
  && chmod +x codecovcli

# Install golangci-lint.
ENV GO_LINT_VERSION=2.5.0
RUN export machine="$(dpkg --print-architecture)"; curl -sSOL https://github.com/golangci/golangci-lint/releases/download/v$GO_LINT_VERSION/golangci-lint-$GO_LINT_VERSION-linux-$machine.tar.gz \
  && tar -C . -xzvf golangci-lint-$GO_LINT_VERSION-linux-$machine.tar.gz \
  && mv golangci-lint-$GO_LINT_VERSION-linux-$machine/golangci-lint . \
  && rm -rf golangci-lint-$GO_LINT_VERSION-linux-$machine/ golangci-lint-$GO_LINT_VERSION-linux-$machine.tar.gz

USER circleci
WORKDIR /home/circleci/

# Setup go lint.
RUN echo "$GO_LINT_VERSION" > .go-lint-version

# Install go tools.
RUN go install github.com/alexfalkowski/gocovmerge/v2@v2.15.0 \
  && go install golang.org/x/vuln/cmd/govulncheck@v1.1.4 \
  && go install gotest.tools/gotestsum@v1.13.0 \
  && go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@v0.37.0 \
  && go install github.com/boyter/scc/v3@v3.5.0 \
  && go install github.com/Zxilly/go-size-analyzer/cmd/gsa@v1.10.0 \
  && go clean --cache -testcache -fuzzcache -modcache && rm -rf "${GOPATH}/pkg"
