FROM circleci/openjdk:15-buster-node-browsers-legacy

USER root

# Install gradle
RUN curl -s "https://get.sdkman.io" | bash && \
  bash -c "source $HOME/.sdkman/bin/sdkman-init.sh && \
  sdk install gradle 6.7"
ENV PATH "$PATH:/opt/gradle/gradle-6.7/bin"

# Install ruby.
RUN cd /tmp && \
  curl -L https://github.com/postmodern/ruby-install/archive/v0.7.1.tar.gz > ruby-install-0.7.1.tar.gz && \
  tar -xzvf ruby-install-0.7.1.tar.gz && \
  cd ruby-install-0.7.1/ && \
  make install
RUN ruby-install ruby 2.7.2
RUN cd /tmp && \
    rm -rf ruby-install-*
ENV PATH "$PATH:/opt/rubies/ruby-2.7.2/bin"
