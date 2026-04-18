FROM alexfalkowski/root:2.3

USER root

# Install release tools from pinned .deb packages.
RUN set -eux; \
  goreleaser_version="2.15.3"; \
  uplift_version="2.26.0"; \
  machine="$(dpkg --print-architecture)"; \
  goreleaser_deb="$(mktemp --suffix=.deb)"; \
  uplift_deb="$(mktemp --suffix=.deb)"; \
  curl -fsSL "https://github.com/goreleaser/goreleaser/releases/download/v${goreleaser_version}/goreleaser_${goreleaser_version}_${machine}.deb" -o "$goreleaser_deb"; \
  curl -fsSL "https://github.com/gembaadvantage/uplift/releases/download/v${uplift_version}/uplift_${uplift_version}_${machine}.deb" -o "$uplift_deb"; \
  apt-get update; \
  apt-get install --no-install-recommends -y "$goreleaser_deb" "$uplift_deb"; \
  apt-get clean; \
  rm -f "$goreleaser_deb" "$uplift_deb"; \
  rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install gh from a downloaded .deb.
RUN set -eux; \
  version="2.90.0"; \
  machine="$(dpkg --print-architecture)"; \
  url="https://github.com/cli/cli/releases/download/v${version}/gh_${version}_linux_${machine}.deb"; \
  out="$(mktemp)"; \
  curl -fsSL "$url" -o "$out"; \
  dpkg -i "$out"; \
  rm -f "$out"

# Add uplift config.
RUN mkdir /etc/uplift
COPY .uplift.yml /etc/uplift

# Copy scripts.
COPY deploy /usr/local/bin/
RUN chmod +x /usr/local/bin/deploy

COPY package /usr/local/bin/
RUN chmod +x /usr/local/bin/package

COPY version /usr/local/bin/
RUN chmod +x /usr/local/bin/version

USER circleci
WORKDIR /home/circleci/

# Setup git.
RUN git config --global user.email "ci@lean-thoughts.com" \
  && git config --global user.name "lean-thoughts-ci"

# Install go tools.
RUN go install github.com/alexfalkowski/infraops/v2/cmd/bump@v2.415.0 \
  && go clean --cache -testcache -fuzzcache -modcache \
  && rm -rf go/pkg \
  && rm -rf .cache
