FROM cimg/go:1.22.3

USER root
WORKDIR /usr/local/bin

# Install kubectl.
RUN curl -sSOL https://dl.k8s.io/release/v1.30.0/bin/linux/amd64/kubectl && \
  chmod +x kubectl

# Install helm.
ENV HELM_VERSION v3.14.4
RUN curl -sSOL https://get.helm.sh/helm-$HELM_VERSION-linux-amd64.tar.gz \
  && tar C . -xf helm-$HELM_VERSION-linux-amd64.tar.gz \
  && mv linux-amd64/helm .

# Install doctl.
ENV DOCTL_VERSION 1.104.0
RUN curl -sSOL https://github.com/digitalocean/doctl/releases/download/v$DOCTL_VERSION/doctl-$DOCTL_VERSION-linux-amd64.tar.gz \
  && tar C . -xf doctl-$DOCTL_VERSION-linux-amd64.tar.gz \
  && rm doctl-$DOCTL_VERSION-linux-amd64.tar.gz

USER circleci
WORKDIR /home/circleci/

# Install go tools.
RUN go install golang.stackrox.io/kube-linter/cmd/kube-linter@v0.6.8 && \
  go install github.com/zegl/kube-score/cmd/kube-score@v1.18.0 \
  && go clean --cache -testcache -fuzzcache -modcache
