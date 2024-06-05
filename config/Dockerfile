FROM cimg/go:1.22.4

USER root
WORKDIR /usr/local/bin

# Cleanup.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install yq.
RUN curl -sSOL https://github.com/mikefarah/yq/releases/download/v4.44.1/yq_linux_amd64 && \
  mv yq_linux_amd64 yq && \
  chmod +x yq

# Install pkl.
RUN curl -sSOL https://github.com/apple/pkl/releases/download/0.25.3/pkl-linux-amd64 && \
  mv pkl-linux-amd64 pkl && \
  chmod +x pkl

USER circleci
WORKDIR /home/circleci/

# Install go tools.
RUN go install github.com/alexfalkowski/servicectl@v0.40.0 \
  && go clean --cache -testcache -fuzzcache -modcache

# Install ruby.
RUN curl -sSL https://rvm.io/mpapis.asc | gpg --import - && \
  curl -sSL https://rvm.io/pkuczynski.asc | gpg --import - && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && rvm requirements && rvm install 3.3.2 && rvm use 3.3.2"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.3.2/bin:$PATH"

# Install gems.
RUN gem update bundler --force
