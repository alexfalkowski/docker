FROM circleci/openjdk:16-buster-node-browsers-legacy

USER root
WORKDIR /usr/local/bin

# Install hadolint
RUN curl -OL https://github.com/hadolint/hadolint/releases/download/v2.1.0/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint

USER circleci
WORKDIR /home/circleci/

# Install scala
RUN curl -s "https://get.sdkman.io" | bash && \
  bash -c "source /home/circleci/.sdkman/bin/sdkman-init.sh && \
  sdk install scala 2.13.5 && \
  sdk install sbt 1.4.9"
ENV PATH "/home/circleci/.sdkman/candidates/sbt/1.4.9/bin:/home/circleci/.sdkman/candidates/scala/2.13.5/bin:$PATH"

# Install ruby.
RUN gpg --keyserver hkp://pool.sks-keyservers.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 7D2BAF1CF37B13E2069D6956105BD0E739499BDB && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && \
  rvm requirements && \
  rvm install 2.7.2 && rvm use 2.7.2"
ENV PATH "/home/circleci/.rvm/rubies/ruby-2.7.2/bin:$PATH"
