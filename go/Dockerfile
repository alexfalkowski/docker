FROM circleci/golang:1.15-buster-node-browsers-legacy

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.31.0

# Install go-mod-outdated
RUN go get -u github.com/psampaz/go-mod-outdated

# Install ruby.
RUN gpg --keyserver hkp://pool.sks-keyservers.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 7D2BAF1CF37B13E2069D6956105BD0E739499BDB && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && \
  rvm requirements && \
  rvm install 2.7.2 && rvm use 2.7.2"
ENV PATH "/home/circleci/.rvm/rubies/ruby-2.7.2/bin:$PATH"