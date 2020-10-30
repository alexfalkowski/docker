FROM circleci/openjdk:15-buster-node-browsers-legacy

# Install gradle
RUN curl -s "https://get.sdkman.io" | bash && \
  bash -c "source /home/circleci/.sdkman/bin/sdkman-init.sh && \
  sdk install gradle 6.7"
ENV PATH "/home/circleci/.sdkman/candidates/gradle/6.7/bin:$PATH"

# Install ruby.
RUN gpg --keyserver hkp://pool.sks-keyservers.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 7D2BAF1CF37B13E2069D6956105BD0E739499BDB && \
  curl -sSL https://get.rvm.io | bash -s stable && \
  bash -c "source $HOME/.rvm/scripts/rvm && \
  rvm requirements && \
  rvm install 2.7.2 && rvm use 2.7.2"
ENV PATH "/home/circleci/.rvm/rubies/ruby-2.7.2/bin:$PATH"
