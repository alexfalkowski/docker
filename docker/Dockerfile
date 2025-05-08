FROM alexfalkowski/root:1.33

USER root
WORKDIR /usr/local/bin

# Install hadolint.
RUN curl -sSOL https://github.com/hadolint/hadolint/releases/download/v2.12.0/hadolint-Linux-x86_64 \
  && mv hadolint-Linux-x86_64 hadolint \
  && chmod +x hadolint

# Install shellcheck.
ENV SHELLCHECK_VERSION=v0.10.0
RUN curl -sSOL https://github.com/koalaman/shellcheck/releases/download/$SHELLCHECK_VERSION/shellcheck-$SHELLCHECK_VERSION.linux.x86_64.tar.xz \
  && tar -C . -xvf shellcheck-$SHELLCHECK_VERSION.linux.x86_64.tar.xz \
  && mv shellcheck-$SHELLCHECK_VERSION/shellcheck . \
  && rm -rf shellcheck-$SHELLCHECK_VERSION/ shellcheck-$SHELLCHECK_VERSION.linux.x86_64.tar.xz
