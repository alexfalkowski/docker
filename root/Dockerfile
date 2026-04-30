FROM cimg/base:2026.04-24.04

USER root

# Install packages.
RUN apt-get update && apt-get install --no-install-recommends -y \
  clang \
  libc6-dev \
  libpq-dev \
  libyaml-dev \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENV GO_VERSION=1.26.2
ENV RUBY_VERSION=ruby-4.0.3

WORKDIR /usr/local/bin
COPY scripts/install-image-tool /usr/local/bin/install-image-tool
COPY root/scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
RUN chmod +x /usr/local/bin/install-image-tool \
  && install-image-tool go "$GO_VERSION" \
  && install-image-tool ruby "$RUBY_VERSION"

# Install gems, https://rubygems.org/gems/rubygems-update.
RUN gem update --system 4.0.8 \
  && gem install bundler -v 4.0.8 \
  && gem install executable-hooks -v 1.7.1

RUN install-image-tool bazelisk v1.29.0

USER circleci
WORKDIR /home/circleci/

# Make flags.
ENV MAKEFLAGS=--no-print-directory

# Setup ruby
RUN echo "$RUBY_VERSION" > .ruby-version

# Setup go.
RUN mkdir -p go/bin
RUN echo "$GO_VERSION" > .go-version
ENV GOPATH=/home/circleci/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

# Setup bazelisk.
ENV USE_BAZEL_VERSION=9.0.0
RUN echo "$USE_BAZEL_VERSION" > .bazelversion
RUN bazelisk version
