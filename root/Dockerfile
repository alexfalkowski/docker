FROM cimg/base:2026.02-24.04

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
ENV GO_VERSION=1.26.0
RUN export machine="$(dpkg --print-architecture)"; curl -sSOL https://go.dev/dl/go$GO_VERSION.linux-$machine.tar.gz \
  && tar -C . -xf go$GO_VERSION.linux-$machine.tar.gz \
  && rm go$GO_VERSION.linux-$machine.tar.gz

# Install ruby.
WORKDIR /home/circleci/
ENV RUBY_VERSION=ruby-3.4.8
ENV RUBY_FILE=$RUBY_VERSION.tar.gz
RUN curl -sSOL https://cache.ruby-lang.org/pub/ruby/3.4/$RUBY_FILE \
  && tar -C . -xzf $RUBY_FILE \
  && rm $RUBY_FILE

WORKDIR /home/circleci/$RUBY_VERSION
RUN ./configure \
  && make \
  && make install

# Install gems, https://rubygems.org/gems/rubygems-update.
RUN gem update --system 4.0.3 \
  && gem install bundler -v 4.0.3 \
  && gem install executable-hooks -v 1.7.1

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
