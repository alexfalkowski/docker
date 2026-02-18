FROM alexfalkowski/root:1.64

USER root
WORKDIR /usr/local/bin

# Install bazelisk.
RUN export machine="$(dpkg --print-architecture)"; curl -sSOL https://github.com/bazelbuild/bazelisk/releases/download/v1.28.1/bazelisk-linux-$machine \
  && mv bazelisk-linux-$machine bazelisk \
  && chmod +x bazelisk

USER circleci
WORKDIR /home/circleci/

# Setup bazelisk.
ENV USE_BAZEL_VERSION=9.0.0
RUN bazelisk version
