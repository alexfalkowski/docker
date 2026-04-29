FROM alexfalkowski/root:2.7

USER root
WORKDIR /usr/local/bin

ARG KUBECTL_VERSION=1.36.0
ARG VEGETA_VERSION=12.12.0
ARG HELM_VERSION=v4.1.1
ARG DOCTL_VERSION=1.155.0
ARG KUBESCAPE_VERSION=4.0.5
ARG PULUMI_VERSION=3.232.0

COPY scripts/install-image-tool /usr/local/bin/install-image-tool
COPY k8s/scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
RUN chmod +x /usr/local/bin/install-image-tool \
  && install-image-tool kubectl \
  && install-image-tool vegeta \
  && install-image-tool helm \
  && install-image-tool doctl \
  && install-image-tool kubescape \
  && install-image-tool pulumi

USER circleci
WORKDIR /home/circleci/

# Install go tools.
RUN go install github.com/zegl/kube-score/cmd/kube-score@v1.20.0 \
  && go clean --cache -testcache -fuzzcache -modcache \
  && rm -rf go/pkg \
  && rm -rf .cache
