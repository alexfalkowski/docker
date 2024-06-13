FROM alexfalkowski/root:1.0

USER root
WORKDIR /usr/local/bin

# Install yq.
RUN curl -sSOL https://github.com/mikefarah/yq/releases/download/v4.44.1/yq_linux_amd64 && \
  mv yq_linux_amd64 yq && \
  chmod +x yq

# Install pkl.
RUN curl -sSOL https://github.com/apple/pkl/releases/download/0.25.3/pkl-linux-amd64 && \
  mv pkl-linux-amd64 pkl && \
  chmod +x pkl

USER circleci
WORKDIR /home/circleci/

# Install go tools.
RUN go install github.com/alexfalkowski/servicectl@v0.45.2 \
  && go clean --cache -testcache -fuzzcache -modcache
