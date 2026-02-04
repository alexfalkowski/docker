FROM alexfalkowski/root:1.63

USER root
WORKDIR /usr/local/bin

# Install hadolint.
RUN export machine="$(uname -m)"; curl -sSOL https://github.com/hadolint/hadolint/releases/download/v2.14.0/hadolint-Linux-$machine \
  && mv hadolint-Linux-$machine hadolint \
  && chmod +x hadolint

# Install shellcheck.
ENV SHELLCHECK_VERSION=v0.11.0
RUN export machine="$(uname -m)"; curl -sSOL https://github.com/koalaman/shellcheck/releases/download/$SHELLCHECK_VERSION/shellcheck-$SHELLCHECK_VERSION.linux.$machine.tar.xz \
  && tar -C . -xvf shellcheck-$SHELLCHECK_VERSION.linux.$machine.tar.xz \
  && mv shellcheck-$SHELLCHECK_VERSION/shellcheck . \
  && rm -rf shellcheck-$SHELLCHECK_VERSION/ shellcheck-$SHELLCHECK_VERSION.linux.$machine.tar.xz
