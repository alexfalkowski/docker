FROM alexfalkowski/root:1.30

USER root
WORKDIR /usr/local/bin

# Install dockerize.
ENV DOCKERIZE_VERSION=v0.6.1
RUN curl -sSOL https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C . -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Install shellcheck.
ENV SHELLCHECK_VERSION=v0.10.0
RUN curl -sSOL https://github.com/koalaman/shellcheck/releases/download/$SHELLCHECK_VERSION/shellcheck-$SHELLCHECK_VERSION.linux.x86_64.tar.xz \
  && tar -C . -xvf shellcheck-$SHELLCHECK_VERSION.linux.x86_64.tar.xz \
  && mv shellcheck-$SHELLCHECK_VERSION/shellcheck . \
  && rm -rf shellcheck-$SHELLCHECK_VERSION/ shellcheck-$SHELLCHECK_VERSION.linux.x86_64.tar.xz

# Install hadolint.
RUN curl -sSOL https://github.com/hadolint/hadolint/releases/download/v2.12.0/hadolint-Linux-x86_64 \
  && mv hadolint-Linux-x86_64 hadolint \
  && chmod +x hadolint

# Install buf.
RUN curl -sSOL https://github.com/bufbuild/buf/releases/download/v1.52.1/buf-Linux-x86_64 \
  && mv buf-Linux-x86_64 buf \
  && chmod +x buf

# Install trivy.
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.61.1

# Install mkcert.
RUN curl -sJLO "https://dl.filippo.io/mkcert/latest?for=linux/amd64" \
  && mv mkcert-v*-linux-amd64 mkcert \
  && chmod +x mkcert

# Install codecov.
RUN curl -sSOL https://cli.codecov.io/v10.4.0/linux/codecov \
  && mv codecov codecovcli \
  && chmod +x codecovcli

USER circleci
WORKDIR /home/circleci/
