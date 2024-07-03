FROM cimg/base:2024.07-22.04

USER root
WORKDIR /usr/local/bin

# Install packages.
RUN apt-get update && apt-get install --no-install-recommends -y \
    g++ \
    libc6-dev \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install go
ENV GO_VERSION 1.22.5
RUN curl -sSOL https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz \
  && tar C . -xf go$GO_VERSION.linux-amd64.tar.gz \
  && rm go$GO_VERSION.linux-amd64.tar.gz

USER circleci
WORKDIR /home/circleci/

# Setup go.
RUN mkdir -p go/bin
ENV GOPATH /home/circleci/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Install ruby.
RUN curl -sSL https://rvm.io/mpapis.asc | gpg --import - && \
  curl -sSL https://rvm.io/pkuczynski.asc | gpg --import - && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && rvm requirements && rvm install 3.3.3 && rvm use 3.3.3"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.3.3/bin:$PATH"

# Install gems.
RUN gem update bundler --force
