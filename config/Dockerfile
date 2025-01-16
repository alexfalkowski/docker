FROM alexfalkowski/root:1.21

USER root
WORKDIR /usr/local/bin

# Install yq.
RUN curl -sSOL https://github.com/mikefarah/yq/releases/download/v4.45.1/yq_linux_amd64 && \
  mv yq_linux_amd64 yq && \
  chmod +x yq

USER circleci
WORKDIR /home/circleci/

# Install go tools.
RUN go install github.com/alexfalkowski/servicectl@v0.173.0 \
  && go clean --cache -testcache -fuzzcache -modcache && rm -rf "${GOPATH}/pkg"
