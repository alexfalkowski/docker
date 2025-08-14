FROM alexfalkowski/root:1.45

USER root
WORKDIR /usr/local/bin

# Install kubectl.
RUN export machine="$(dpkg --print-architecture)"; curl -sSOL https://dl.k8s.io/release/v1.33.4/bin/linux/$machine/kubectl \
  && chmod +x kubectl

# Install vegeta.
ARG VEGETA_VERSION=12.12.0
RUN export machine="$(dpkg --print-architecture)"; curl -sSOL https://github.com/tsenart/vegeta/releases/download/v${VEGETA_VERSION}/vegeta_${VEGETA_VERSION}_linux_$machine.tar.gz \
  && tar C . -xf vegeta_${VEGETA_VERSION}_linux_$machine.tar.gz \
  && chmod +x vegeta

# Install helm.
ARG HELM_VERSION=v3.18.5
RUN export machine="$(dpkg --print-architecture)"; curl -sSOL https://get.helm.sh/helm-${HELM_VERSION}-linux-$machine.tar.gz \
  && tar C . -xf helm-${HELM_VERSION}-linux-$machine.tar.gz \
  && mv linux-$machine/helm . \
  && rm -rf linux-$machine \
  && rm helm-${HELM_VERSION}-linux-$machine.tar.gz

# Install doctl.
ARG DOCTL_VERSION=1.138.0
RUN export machine="$(dpkg --print-architecture)"; curl -sSOL https://github.com/digitalocean/doctl/releases/download/v${DOCTL_VERSION}/doctl-${DOCTL_VERSION}-linux-$machine.tar.gz \
  && tar C . -xf doctl-${DOCTL_VERSION}-linux-$machine.tar.gz \
  && rm doctl-${DOCTL_VERSION}-linux-$machine.tar.gz

# Install kubescape.
RUN curl -s https://raw.githubusercontent.com/kubescape/kubescape/master/install.sh | /bin/bash -s -- -v v3.0.37

USER circleci
WORKDIR /home/circleci/

# Install pulumi.
RUN curl -fsSL https://get.pulumi.com | sh -s -- --version 3.189.0
ENV PATH=/home/circleci/.pulumi/bin:$PATH

# Install go tools.
RUN go install github.com/zegl/kube-score/cmd/kube-score@v1.20.0 \
  && go clean --cache -testcache -fuzzcache -modcache && rm -rf "${GOPATH}/pkg"
