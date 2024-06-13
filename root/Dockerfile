FROM cimg/go:1.22.4

USER root
WORKDIR /usr/local/bin

# Cleanup.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

USER circleci
WORKDIR /home/circleci/

# Install ruby.
RUN curl -sSL https://rvm.io/mpapis.asc | gpg --import - && \
  curl -sSL https://rvm.io/pkuczynski.asc | gpg --import - && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && rvm requirements && rvm install 3.3.3 && rvm use 3.3.3"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.3.3/bin:$PATH"

# Install gems.
RUN gem update bundler --force
