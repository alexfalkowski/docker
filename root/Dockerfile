FROM cimg/base:2026.04-24.04

USER root

# Install packages.
RUN apt-get update && apt-get install --no-install-recommends -y \
  clang \
  libc6-dev \
  libpq-dev \
  libyaml-dev \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install go
WORKDIR /usr/local
ENV GO_VERSION=1.26.2
RUN set -eu; \
  machine="$(dpkg --print-architecture)"; \
  asset="go${GO_VERSION}.linux-${machine}.tar.gz"; \
  curl -fsSLO "https://go.dev/dl/${asset}"; \
  tar -C . -xf "$asset"; \
  rm -f "$asset"

# Install ruby.
WORKDIR /home/circleci/
ENV RUBY_VERSION=ruby-4.0.2
ENV RUBY_FILE=$RUBY_VERSION.tar.gz
RUN curl -fsSLO https://cache.ruby-lang.org/pub/ruby/4.0/$RUBY_FILE \
  && tar -C . -xzf $RUBY_FILE \
  && rm $RUBY_FILE

WORKDIR /home/circleci/$RUBY_VERSION
RUN ./configure \
  && make \
  && make install \
  && rm -rf /home/circleci/$RUBY_VERSION

# Install gems, https://rubygems.org/gems/rubygems-update.
RUN gem update --system 4.0.8 \
  && gem install bundler -v 4.0.8 \
  && gem install executable-hooks -v 1.7.1

# Install bazelisk.
WORKDIR /usr/local/bin
RUN set -eu; \
  machine="$(dpkg --print-architecture)"; \
  asset="bazelisk-linux-${machine}"; \
  curl -fsSLo "$asset" "https://github.com/bazelbuild/bazelisk/releases/download/v1.28.1/${asset}"; \
  mv "$asset" bazelisk; \
  chmod +x bazelisk

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
