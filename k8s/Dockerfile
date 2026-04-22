FROM alexfalkowski/root:2.5

USER root
WORKDIR /usr/local/bin

# Install kubectl.
RUN set -eu; \
  machine="$(dpkg --print-architecture)"; \
  base_url="https://dl.k8s.io/release/v1.35.1/bin/linux/${machine}"; \
  curl -fsSLo kubectl "$base_url/kubectl"; \
  curl -fsSLo kubectl.sha256 "$base_url/kubectl.sha256"; \
  echo "$(cat kubectl.sha256)  kubectl" | sha256sum -c -; \
  chmod +x kubectl; \
  rm -f kubectl.sha256

# Install vegeta.
ARG VEGETA_VERSION=12.12.0
RUN set -eu; \
  machine="$(dpkg --print-architecture)"; \
  asset="vegeta_${VEGETA_VERSION}_linux_${machine}.tar.gz"; \
  base_url="https://github.com/tsenart/vegeta/releases/download/v${VEGETA_VERSION}"; \
  curl -fsSLO "$base_url/$asset"; \
  curl -fsSLO "$base_url/vegeta_${VEGETA_VERSION}_checksums.txt"; \
  grep "  ${asset}$" "vegeta_${VEGETA_VERSION}_checksums.txt" | sha256sum -c -; \
  tar -C . -xf "$asset"; \
  rm -f "$asset" "vegeta_${VEGETA_VERSION}_checksums.txt"; \
  chmod +x vegeta

# Install helm.
ARG HELM_VERSION=v4.1.1
RUN set -eu; \
  machine="$(dpkg --print-architecture)"; \
  asset="helm-${HELM_VERSION}-linux-${machine}.tar.gz"; \
  base_url="https://get.helm.sh"; \
  curl -fsSLO "$base_url/$asset"; \
  curl -fsSLO "$base_url/$asset.sha256sum"; \
  sha256sum -c "$asset.sha256sum"; \
  tar -C . -xf "$asset"; \
  mv "linux-${machine}/helm" .; \
  rm -rf "linux-${machine}" "$asset" "$asset.sha256sum"

# Install doctl.
ARG DOCTL_VERSION=1.155.0
RUN set -eu; \
  machine="$(dpkg --print-architecture)"; \
  asset="doctl-${DOCTL_VERSION}-linux-${machine}.tar.gz"; \
  base_url="https://github.com/digitalocean/doctl/releases/download/v${DOCTL_VERSION}"; \
  curl -fsSLO "$base_url/$asset"; \
  curl -fsSLO "$base_url/doctl-${DOCTL_VERSION}-checksums.sha256"; \
  grep "  ${asset}$" "doctl-${DOCTL_VERSION}-checksums.sha256" | sha256sum -c -; \
  tar -C . -xf "$asset"; \
  rm -f "$asset" "doctl-${DOCTL_VERSION}-checksums.sha256"

# Install kubescape.
RUN set -eu; \
  version="4.0.5"; \
  machine="$(dpkg --print-architecture)"; \
  asset="kubescape_${version}_linux_${machine}.tar.gz"; \
  base_url="https://github.com/kubescape/kubescape/releases/download/v${version}"; \
  curl -fsSLO "$base_url/$asset"; \
  curl -fsSLO "$base_url/checksums.sha256"; \
  grep "  ${asset}$" checksums.sha256 | sha256sum -c -; \
  tar -C . -xzf "$asset" kubescape; \
  rm -f "$asset" checksums.sha256

# Install pulumi.
RUN set -eux; \
  version="3.231.0"; \
  machine="$(dpkg --print-architecture)"; \
  case "$machine" in \
  amd64) pulumi_arch="x64" ;; \
  arm64) pulumi_arch="arm64" ;; \
  *) echo "unsupported arch: $machine" >&2; exit 1 ;; \
  esac; \
  asset="pulumi-v${version}-linux-${pulumi_arch}.tar.gz"; \
  base_url="https://github.com/pulumi/pulumi/releases/download/v${version}"; \
  curl -fsSLO "$base_url/$asset"; \
  curl -fsSLO "$base_url/pulumi-${version}-checksums.txt"; \
  grep "  ${asset}$" "pulumi-${version}-checksums.txt" | sha256sum -c -; \
  tar -xzf "$asset" --strip-components=1; \
  rm -f "$asset" "pulumi-${version}-checksums.txt"

USER circleci
WORKDIR /home/circleci/

# Install go tools.
RUN go install github.com/zegl/kube-score/cmd/kube-score@v1.20.0 \
  && go clean --cache -testcache -fuzzcache -modcache \
  && rm -rf go/pkg \
  && rm -rf .cache
