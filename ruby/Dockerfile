FROM alexfalkowski/root:1.61

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
RUN export machine="$(uname -m)"; curl -sSOL https://github.com/bufbuild/buf/releases/download/v1.64.0/buf-Linux-$machine \
  && mv buf-Linux-$machine buf \
  && chmod +x buf

# Install trivy.
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.69.0

# Install mkcert.
RUN export machine="$(dpkg --print-architecture)"; curl -sJLO "https://dl.filippo.io/mkcert/latest?for=linux/$machine" \
  && mv mkcert-v*-linux-$machine mkcert \
  && chmod +x mkcert

# Install codecov.
RUN curl -sSOL https://cli.codecov.io/v11.2.6/linux/codecov \
  && mv codecov codecovcli \
  && chmod +x codecovcli

USER circleci
WORKDIR /home/circleci/
