FROM ruby:4.0.3-slim-trixie

USER root

# Install packages.
RUN apt-get update && apt-get install --no-install-recommends -y \
  bash \
  build-essential \
  ca-certificates \
  clang \
  curl \
  docker-buildx \
  docker-cli \
  docker-compose \
  docker.io \
  git \
  libc6-dev \
  libffi-dev \
  libgdbm-dev \
  libncurses-dev \
  libpq-dev \
  libreadline-dev \
  libssl-dev \
  libyaml-dev \
  make \
  netcat-openbsd \
  openssh-client \
  pkg-config \
  sudo \
  xz-utils \
  zlib1g-dev \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN groupadd --gid 3434 circleci \
  && useradd --uid 3434 --gid circleci --create-home --shell /bin/bash circleci \
  && echo "circleci ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/circleci \
  && chmod 0440 /etc/sudoers.d/circleci

ENV GO_VERSION=1.26.2

WORKDIR /usr/local/bin
COPY scripts/install-image-tool /usr/local/bin/install-image-tool
COPY root/scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
RUN chmod +x /usr/local/bin/install-image-tool \
  && install-image-tool go "$GO_VERSION"

# Install gems, https://rubygems.org/gems/rubygems-update.
RUN gem update --system 4.0.11 \
  && gem install bundler -v 4.0.11 \
  && gem install executable-hooks -v 1.7.1

USER circleci
WORKDIR /home/circleci/

# Make flags.
ENV MAKEFLAGS=--no-print-directory

# Setup ruby
RUN echo "ruby-$RUBY_VERSION" > .ruby-version

# Setup go.
RUN mkdir -p go/bin
RUN echo "$GO_VERSION" > .go-version
ENV GOPATH=/home/circleci/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
