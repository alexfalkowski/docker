FROM circleci/ruby:2.7-buster-node-browsers-legacy

USER root
WORKDIR /usr/local/bin

# Install hadolint
RUN curl -OL https://github.com/hadolint/hadolint/releases/download/v1.18.2/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint

USER circleci
WORKDIR /home/circleci/
