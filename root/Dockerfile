FROM cimg/base:2024.07-22.04

USER root

# Install packages.
RUN apt-get update && apt-get install --no-install-recommends -y \
    clang \
    libc6-dev \
    libpq-dev \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install go
ENV GO_VERSION 1.23.2
RUN curl -sSOL https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz \
  && tar C /usr/local/ -xf go$GO_VERSION.linux-amd64.tar.gz \
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
  bash -c "source $HOME/.rvm/scripts/rvm && rvm requirements && rvm install 3.3.5 && rvm use 3.3.5"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.3.5/bin:$PATH"

# Install gems.
RUN gem update bundler --force
