FROM alexfalkowski/root:2.3

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
RUN set -eu; \
  arch="$(uname -m)"; \
  case "$arch" in \
  x86_64|amd64) hadolint_arch="x86_64" ;; \
  aarch64|arm64) hadolint_arch="arm64" ;; \
  *) echo "unsupported arch: $arch" >&2; exit 1 ;; \
  esac; \
  curl -fLsSLo hadolint "https://github.com/hadolint/hadolint/releases/download/v2.14.0/hadolint-linux-${hadolint_arch}" \
  && chmod +x hadolint

# Install buf.
RUN export machine="$(uname -m)"; curl -sSOL https://github.com/bufbuild/buf/releases/download/v1.68.1/buf-Linux-$machine \
  && mv buf-Linux-$machine buf \
  && chmod +x buf

# Install trivy.
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.70.0

# Install mkcert.
RUN export machine="$(dpkg --print-architecture)"; curl -sJLO "https://dl.filippo.io/mkcert/latest?for=linux/$machine" \
  && mv mkcert-v*-linux-$machine mkcert \
  && chmod +x mkcert

# Install codecov.
RUN set -eux; \
  arch="$(uname -m)"; \
  suffix=""; \
  case "$arch" in \
  x86_64|amd64) suffix="" ;; \
  aarch64|arm64) suffix="_arm64" ;; \
  *) echo "unsupported arch: $arch" >&2; exit 1 ;; \
  esac; \
  curl -fLsSLo codecovcli \
  "https://github.com/getsentry/prevent-cli/releases/download/v11.2.8/codecovcli_linux${suffix}"; \
  chmod +x codecovcli

# Install golangci-lint.
ENV GO_LINT_VERSION=2.11.4
RUN export machine="$(dpkg --print-architecture)"; curl -sSOL https://github.com/golangci/golangci-lint/releases/download/v$GO_LINT_VERSION/golangci-lint-$GO_LINT_VERSION-linux-$machine.tar.gz \
  && tar -C . -xzvf golangci-lint-$GO_LINT_VERSION-linux-$machine.tar.gz \
  && mv golangci-lint-$GO_LINT_VERSION-linux-$machine/golangci-lint . \
  && rm -rf golangci-lint-$GO_LINT_VERSION-linux-$machine/ golangci-lint-$GO_LINT_VERSION-linux-$machine.tar.gz

USER circleci
WORKDIR /home/circleci/

# Setup go lint.
RUN echo "$GO_LINT_VERSION" > .go-lint-version

# Install go tools.
RUN go install github.com/alexfalkowski/gocovmerge/v2@v2.26.0 \
  && go install golang.org/x/vuln/cmd/govulncheck@v1.1.4 \
  && go install gotest.tools/gotestsum@v1.13.0 \
  && go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@v0.44.0 \
  && go install github.com/boyter/scc/v3@v3.7.0 \
  && go install github.com/Zxilly/go-size-analyzer/cmd/gsa@v1.12.1 \
  && go clean --cache -testcache -fuzzcache -modcache \
  && rm -rf go/pkg \
  && rm -rf .cache
