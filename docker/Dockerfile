FROM cimg/go:1.22.4

USER root
WORKDIR /usr/local/bin

# Cleanup.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install hadolint.
RUN curl -sSOL https://github.com/hadolint/hadolint/releases/download/v2.12.0/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint

# Install slim.
RUN curl -sSOL https://github.com/slimtoolkit/slim/releases/download/1.40.11/dist_linux.tar.gz \
  && tar C . -xf dist_linux.tar.gz \
  && rm dist_linux.tar.gz

USER circleci
WORKDIR /home/circleci/

# Install ruby.
RUN curl -sSL https://rvm.io/mpapis.asc | gpg --import - && \
  curl -sSL https://rvm.io/pkuczynski.asc | gpg --import - && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && rvm requirements && rvm install 3.3.2 && rvm use 3.3.2"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.3.2/bin:$PATH"

# Install gems.
RUN gem update bundler --force
