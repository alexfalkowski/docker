FROM alexfalkowski/root:1.65

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
  curl -fLsSLo hadolint "https://github.com/hadolint/hadolint/releases/download/v2.14.0/hadolint-Linux-${hadolint_arch}" \
  && chmod +x hadolint

# Install shellcheck.
ENV SHELLCHECK_VERSION=v0.11.0
RUN export machine="$(uname -m)"; curl -sSOL https://github.com/koalaman/shellcheck/releases/download/$SHELLCHECK_VERSION/shellcheck-$SHELLCHECK_VERSION.linux.$machine.tar.xz \
  && tar -C . -xvf shellcheck-$SHELLCHECK_VERSION.linux.$machine.tar.xz \
  && mv shellcheck-$SHELLCHECK_VERSION/shellcheck . \
  && rm -rf shellcheck-$SHELLCHECK_VERSION/ shellcheck-$SHELLCHECK_VERSION.linux.$machine.tar.xz
