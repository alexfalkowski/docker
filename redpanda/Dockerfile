FROM cimg/base:2021.04

USER root

## Run the setup script to download and install the repo
RUN curl -1sLf 'https://dl.redpanda.com/nzc4ZYQK3WRGd9sy/redpanda/cfg/setup/bash.deb.sh' | bash && \
    ## Use apt to install redpanda
    apt-get update && apt-get install --no-install-recommends -y \
    redpanda \
    && rm -rf /var/lib/apt/lists/*

## Start redpanda as a service
RUN systemctl start redpanda

USER circleci
WORKDIR /home/circleci/
CMD ["redpanda", "start"]
