FROM circleci/openjdk:16-buster-node-browsers-legacy

USER root
WORKDIR /usr/local/bin

# Install dockerize
ENV DOCKERIZE_VERSION v0.6.1
RUN curl -OL https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C . -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Install hadolint
RUN curl -OL https://github.com/hadolint/hadolint/releases/download/v2.5.0/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint

USER circleci
WORKDIR /home/circleci/

# Install kotlin
RUN curl -s "https://get.sdkman.io" | bash && \
  bash -c "source /home/circleci/.sdkman/bin/sdkman-init.sh && \
  sdk install kotlin 1.5.0 && \
  sdk install gradle 7.0.2"
ENV PATH "/home/circleci/.sdkman/candidates/kotlin/1.5.0/bin:/home/circleci/.sdkman/candidates/gradle/6.8.3/bin:$PATH"

# Install ruby.
RUN gpg --keyserver hkp://pool.sks-keyservers.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 7D2BAF1CF37B13E2069D6956105BD0E739499BDB && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && \
  rvm requirements && \
  rvm install 3.0.0 && rvm use 3.0.0"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.0.0/bin:$PATH"

# Update bundler
RUN gem update bundler
