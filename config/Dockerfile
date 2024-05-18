FROM cimg/go:1.22.3

USER root
WORKDIR /usr/local/bin

# Install yq.
RUN curl -sSOL https://github.com/mikefarah/yq/releases/download/v4.44.1/yq_linux_amd64 && \
  mv yq_linux_amd64 yq && \
  chmod +x yq

USER circleci
WORKDIR /home/circleci/
