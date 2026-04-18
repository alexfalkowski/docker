FROM alexfalkowski/root:2.4

USER root
WORKDIR /usr/local/bin

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

# Install shellcheck.
ENV SHELLCHECK_VERSION=v0.11.0
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
