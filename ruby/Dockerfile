FROM alexfalkowski/root:2.7

USER root
WORKDIR /usr/local/bin

ENV DOCKERIZE_VERSION=v0.9.3
ENV SHELLCHECK_VERSION=v0.11.0
ENV MKCERT_VERSION=v1.4.4
ARG HADOLINT_VERSION=v2.14.0
ARG BUF_VERSION=v1.69.0
ARG TRIVY_VERSION=0.70.0
ARG CODECOV_VERSION=v11.2.8

COPY scripts/install-image-tool /usr/local/bin/install-image-tool
COPY scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
COPY ruby/scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
RUN chmod +x /usr/local/bin/install-image-tool \
  && install-image-tool dockerize \
  && install-image-tool shellcheck \
  && install-image-tool hadolint \
  && install-image-tool buf \
  && install-image-tool trivy \
  && install-image-tool mkcert \
  && install-image-tool codecov

USER circleci
WORKDIR /home/circleci/

# Install go tools.
RUN go install github.com/boyter/scc/v3@v3.7.0 \
  && go clean --cache -testcache -fuzzcache -modcache \
  && rm -rf go/pkg \
  && rm -rf .cache
