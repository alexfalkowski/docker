FROM alexfalkowski/root:1.30

USER root
WORKDIR /usr/local/bin

# Install kubectl.
RUN curl -sSOL https://dl.k8s.io/release/v1.32.3/bin/linux/amd64/kubectl \
  && chmod +x kubectl

# Install helm.
ENV HELM_VERSION=v3.17.3
RUN curl -sSOL https://get.helm.sh/helm-$HELM_VERSION-linux-amd64.tar.gz \
  && tar C . -xf helm-$HELM_VERSION-linux-amd64.tar.gz \
  && mv linux-amd64/helm . \
  && rm -rf linux-amd64 \
  && rm helm-$HELM_VERSION-linux-amd64.tar.gz

# Install doctl.
ENV DOCTL_VERSION=1.124.0
RUN curl -sSOL https://github.com/digitalocean/doctl/releases/download/v$DOCTL_VERSION/doctl-$DOCTL_VERSION-linux-amd64.tar.gz \
  && tar C . -xf doctl-$DOCTL_VERSION-linux-amd64.tar.gz \
  && rm doctl-$DOCTL_VERSION-linux-amd64.tar.gz

# Install pulumi.
ENV PULUMI_VERSION=v3.163.0
RUN curl -sSOL https://get.pulumi.com/releases/sdk/pulumi-$PULUMI_VERSION-linux-x64.tar.gz \
  && tar C . -xf pulumi-$PULUMI_VERSION-linux-x64.tar.gz \
  && mv pulumi temp \
  && mv temp/* . \
  && rm -rf temp \
  && rm pulumi-$PULUMI_VERSION-linux-x64.tar.gz

# Install kubescape.
RUN curl -sSOL https://github.com/kubescape/kubescape/releases/download/v3.0.33/kubescape-ubuntu-latest \
  && mv kubescape-ubuntu-latest kubescape \
  && chmod +x kubescape

USER circleci
WORKDIR /home/circleci/

# Install go tools.
RUN go install github.com/zegl/kube-score/cmd/kube-score@v1.19.0 \
  && go clean --cache -testcache -fuzzcache -modcache && rm -rf "${GOPATH}/pkg"
