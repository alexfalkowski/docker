FROM alexfalkowski/root:2.3

USER root
WORKDIR /usr/local/bin

# Install dockerize.
ENV DOCKERIZE_VERSION=v0.9.3
RUN set -eu; \
  machine="$(dpkg --print-architecture)"; \
  asset="dockerize-linux-${machine}-${DOCKERIZE_VERSION}.tar.gz"; \
  curl -fsSLO "https://github.com/jwilder/dockerize/releases/download/${DOCKERIZE_VERSION}/${asset}"; \
  tar -C . -xzvf "$asset"; \
  rm -f "$asset"

# Install shellcheck.
ENV SHELLCHECK_VERSION=v0.10.0
RUN set -eu; \
  arch="$(uname -m)"; \
  case "$arch" in \
  x86_64|amd64) shellcheck_arch="x86_64" ;; \
  aarch64|arm64) shellcheck_arch="aarch64" ;; \
  *) echo "unsupported arch: $arch" >&2; exit 1 ;; \
  esac; \
  asset="shellcheck-${SHELLCHECK_VERSION}.linux.${shellcheck_arch}.tar.xz"; \
  curl -fsSLO "https://github.com/koalaman/shellcheck/releases/download/${SHELLCHECK_VERSION}/${asset}"; \
  tar -C . -xJf "$asset"; \
  mv shellcheck-${SHELLCHECK_VERSION}/shellcheck .; \
  rm -rf "shellcheck-${SHELLCHECK_VERSION}" "$asset"

# Install hadolint.
RUN set -eu; \
  arch="$(uname -m)"; \
  case "$arch" in \
  x86_64|amd64) hadolint_arch="x86_64" ;; \
  aarch64|arm64) hadolint_arch="arm64" ;; \
  *) echo "unsupported arch: $arch" >&2; exit 1 ;; \
  esac; \
  asset="hadolint-linux-${hadolint_arch}"; \
  base_url="https://github.com/hadolint/hadolint/releases/download/v2.14.0"; \
  curl -fsSLo "$asset" "$base_url/$asset"; \
  curl -fsSLo "$asset.sha256" "$base_url/$asset.sha256"; \
  sha256sum -c "$asset.sha256"; \
  mv "$asset" hadolint; \
  chmod +x hadolint; \
  rm -f "$asset.sha256"

# Install buf.
RUN set -eu; \
  arch="$(uname -m)"; \
  case "$arch" in \
  x86_64|amd64) buf_arch="x86_64" ;; \
  aarch64|arm64) buf_arch="aarch64" ;; \
  *) echo "unsupported arch: $arch" >&2; exit 1 ;; \
  esac; \
  asset="buf-Linux-${buf_arch}"; \
  base_url="https://github.com/bufbuild/buf/releases/download/v1.68.1"; \
  curl -fsSLo "$asset" "$base_url/$asset"; \
  curl -fsSLo sha256.txt "$base_url/sha256.txt"; \
  grep "  ${asset}$" sha256.txt | sha256sum -c -; \
  mv "$asset" buf; \
  chmod +x buf; \
  rm -f sha256.txt

# Install trivy.
RUN set -eu; \
  version="0.70.0"; \
  machine="$(dpkg --print-architecture)"; \
  case "$machine" in \
  amd64) trivy_arch="64bit" ;; \
  arm64) trivy_arch="ARM64" ;; \
  *) echo "unsupported arch: $machine" >&2; exit 1 ;; \
  esac; \
  asset="trivy_${version}_Linux-${trivy_arch}.tar.gz"; \
  base_url="https://github.com/aquasecurity/trivy/releases/download/v${version}"; \
  curl -fsSLO "$base_url/$asset"; \
  curl -fsSLO "$base_url/trivy_${version}_checksums.txt"; \
  grep "  ${asset}$" "trivy_${version}_checksums.txt" | sha256sum -c -; \
  tar -C . -xzf "$asset" trivy; \
  rm -f "$asset" "trivy_${version}_checksums.txt"

# Install mkcert.
ENV MKCERT_VERSION=v1.4.4
RUN set -eu; \
  machine="$(dpkg --print-architecture)"; \
  asset="mkcert-${MKCERT_VERSION}-linux-${machine}"; \
  curl -fsSLo "$asset" "https://github.com/FiloSottile/mkcert/releases/download/${MKCERT_VERSION}/${asset}"; \
  mv "$asset" mkcert; \
  chmod +x mkcert

# Install codecov.
RUN set -eux; \
  arch="$(uname -m)"; \
  case "$arch" in \
  x86_64|amd64) asset="codecovcli_linux" ;; \
  aarch64|arm64) asset="codecovcli_linux_arm64" ;; \
  *) echo "unsupported arch: $arch" >&2; exit 1 ;; \
  esac; \
  curl -fLsSLo codecovcli \
  "https://github.com/getsentry/prevent-cli/releases/download/v11.2.8/${asset}"; \
  chmod +x codecovcli

# Install golangci-lint.
ENV GO_LINT_VERSION=2.11.4
RUN set -eu; \
  machine="$(dpkg --print-architecture)"; \
  asset="golangci-lint-${GO_LINT_VERSION}-linux-${machine}.tar.gz"; \
  base_url="https://github.com/golangci/golangci-lint/releases/download/v${GO_LINT_VERSION}"; \
  curl -fsSLO "$base_url/$asset"; \
  curl -fsSLO "$base_url/golangci-lint-${GO_LINT_VERSION}-checksums.txt"; \
  grep "  ${asset}$" "golangci-lint-${GO_LINT_VERSION}-checksums.txt" | sha256sum -c -; \
  tar -C . -xzvf "$asset"; \
  mv "golangci-lint-${GO_LINT_VERSION}-linux-${machine}/golangci-lint" .; \
  rm -rf "golangci-lint-${GO_LINT_VERSION}-linux-${machine}" "$asset" "golangci-lint-${GO_LINT_VERSION}-checksums.txt"

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
