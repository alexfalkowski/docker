FROM cimg/base:2025.04

USER root

# Install packages.
RUN apt-get update && apt-get install --no-install-recommends -y \
  clang \
  libc6-dev \
  libpq-dev \
  libyaml-dev \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install just.
WORKDIR /usr/local
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to bin \
  && chmod +x bin/just

# Install go
ENV GO_VERSION=1.24.2
RUN curl -sSOL https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz \
  && tar -C . -xf go$GO_VERSION.linux-amd64.tar.gz \
  && rm go$GO_VERSION.linux-amd64.tar.gz

# Install ruby.
WORKDIR /home/circleci/
ENV RUBY_VERSION=ruby-3.4.2
ENV RUBY_FILE=$RUBY_VERSION.tar.gz
RUN curl -sSOL https://cache.ruby-lang.org/pub/ruby/3.4/$RUBY_FILE \
  && tar -C . -xzf $RUBY_FILE \
  && rm $RUBY_FILE

WORKDIR /home/circleci/$RUBY_VERSION
RUN ./configure \
  && make \
  && make install

# Install gems.
RUN gem update --system 3.6.5 \
  && gem install bundler -v 2.6.5 \
  && gem install executable-hooks -v 1.7.1

USER circleci
WORKDIR /home/circleci/

# Setup go.
RUN mkdir -p go/bin
ENV GOPATH=/home/circleci/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
