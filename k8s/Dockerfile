FROM alexfalkowski/root:2.8

USER root
WORKDIR /usr/local/bin

COPY scripts/install-image-tool /usr/local/bin/install-image-tool
COPY k8s/scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
RUN chmod +x /usr/local/bin/install-image-tool \
  && install-image-tool kubectl 1.36.0 \
  && install-image-tool vegeta 12.12.0 \
  && install-image-tool helm v4.1.1 \
  && install-image-tool doctl 1.155.0 \
  && install-image-tool kubescape 4.0.5 \
  && install-image-tool pulumi 3.233.0

USER circleci
WORKDIR /home/circleci/

# Install go tools.
RUN go install github.com/zegl/kube-score/cmd/kube-score@v1.20.0 \
  && go clean --cache -testcache -fuzzcache -modcache \
  && rm -rf go/pkg \
  && rm -rf .cache
