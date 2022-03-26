FROM cimg/openjdk:17.0.2

USER root
WORKDIR /usr/local/bin

# Install dockerize.
ENV DOCKERIZE_VERSION v0.6.1
RUN curl -sSOL https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C . -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Install hadolint.
RUN curl -sSOL https://github.com/hadolint/hadolint/releases/download/v2.9.3/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint

# Install buf.
RUN curl -sSOL https://github.com/bufbuild/buf/releases/download/v1.3.0/buf-Linux-x86_64 && \
  mv buf-Linux-x86_64 buf && \
  chmod +x buf

# Install jq.
RUN curl -sSOL https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && \
  mv jq-linux64 jq && \
  chmod +x jq

# Install yq.
RUN curl -sSOL https://github.com/mikefarah/yq/releases/download/v4.23.1/yq_linux_amd64 && \
  mv yq_linux_amd64 yq && \
  chmod +x yq

USER circleci
WORKDIR /home/circleci/

# Install scala.
RUN curl -sSL "https://get.sdkman.io" | bash && \
  bash -c "source /home/circleci/.sdkman/bin/sdkman-init.sh && \
  sdk install scala 3.1.0 && \
  sdk install sbt 1.6.2"
ENV PATH "/home/circleci/.sdkman/candidates/sbt/1.6.2/bin:/home/circleci/.sdkman/candidates/scala/3.1.0/bin:$PATH"

# Install ruby.
RUN curl -sSL https://rvm.io/mpapis.asc | gpg --import - && \
  curl -sSL https://rvm.io/pkuczynski.asc | gpg --import - && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && \
  rvm requirements && \
  rvm install 3.0.3 && rvm use 3.0.3"
ENV PATH "/home/circleci/.rvm/rubies/ruby-3.0.3/bin:$PATH"

# Update bundler.
RUN gem update bundler
