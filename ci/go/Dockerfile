FROM circleci/golang:1.15-buster-node-browsers-legacy

USER root

RUN cd /tmp && \
  curl -L https://github.com/postmodern/ruby-install/archive/v0.7.1.tar.gz > ruby-install-0.7.1.tar.gz && \
  tar -xzvf ruby-install-0.7.1.tar.gz && \
  cd ruby-install-0.7.1/ && \
  make install
RUN ruby-install ruby 2.7.2
RUN cd /tmp && \
    rm -rf ruby-install-*
ENV PATH "$PATH:/opt/rubies/ruby-2.7.2/bin"

ENTRYPOINT ["docker-compose", "-f", "/etc/docker-compose.yml", "up", "-d"]
