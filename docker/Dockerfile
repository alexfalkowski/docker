FROM cimg/go:1.22.4

USER root
WORKDIR /usr/local/bin

# Cleanup.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install slim.
RUN curl -sSOL https://github.com/slimtoolkit/slim/releases/download/1.40.11/dist_linux.tar.gz \
  && tar C . -xf dist_linux.tar.gz \
  && rm dist_linux.tar.gz

USER circleci
WORKDIR /home/circleci/
