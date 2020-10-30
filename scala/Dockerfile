FROM openjdk:15-jdk-buster

USER root

# Install packages
RUN apt-get update && apt-get install -y \
  curl \
  unzip \
  zip \
  build-essential \
  bison \
  zlib1g-dev \
  libyaml-dev \
  libssl-dev \
  libgdbm-dev \
  libreadline-dev \
  libncurses5-dev \
  libffi-dev \
  && rm -rf /var/lib/apt/lists/*

# Install scala
RUN curl -s "https://get.sdkman.io" | bash && \
  bash -c "source /root/.sdkman/bin/sdkman-init.sh && \
  sdk install scala 2.13.3 && \
  sdk install sbt 1.4.1"
ENV PATH "/root/.sdkman/candidates/sbt/1.4.1/bin:/root/.sdkman/candidates/scala/2.13.3/bin:$PATH"
RUN sbt version -Dsbt.rootdir=true

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
