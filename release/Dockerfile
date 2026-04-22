FROM alexfalkowski/root:2.5

USER root

# Install release tools from pinned .deb packages.
RUN set -eux; \
  goreleaser_version="2.15.4"; \
  uplift_version="2.26.0"; \
  machine="$(dpkg --print-architecture)"; \
  goreleaser_deb="goreleaser_${goreleaser_version}_${machine}.deb"; \
  uplift_deb="uplift_${uplift_version}_${machine}.deb"; \
  curl -fsSLO "https://github.com/goreleaser/goreleaser/releases/download/v${goreleaser_version}/${goreleaser_deb}"; \
  curl -fsSLO "https://github.com/goreleaser/goreleaser/releases/download/v${goreleaser_version}/checksums.txt"; \
  grep "  ${goreleaser_deb}$" checksums.txt | sha256sum -c -; \
  curl -fsSLO "https://github.com/gembaadvantage/uplift/releases/download/v${uplift_version}/${uplift_deb}"; \
  curl -fsSLo uplift-checksums.txt "https://github.com/gembaadvantage/uplift/releases/download/v${uplift_version}/checksums.txt"; \
  grep "  ${uplift_deb}$" uplift-checksums.txt | sha256sum -c -; \
  apt-get update; \
  apt-get install --no-install-recommends -y "./${goreleaser_deb}" "./${uplift_deb}"; \
  apt-get clean; \
  rm -f "$goreleaser_deb" "$uplift_deb" checksums.txt uplift-checksums.txt; \
  rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install gh from a downloaded .deb.
RUN set -eux; \
  version="2.90.0"; \
  machine="$(dpkg --print-architecture)"; \
  asset="gh_${version}_linux_${machine}.deb"; \
  base_url="https://github.com/cli/cli/releases/download/v${version}"; \
  curl -fsSLO "$base_url/$asset"; \
  curl -fsSLO "$base_url/gh_${version}_checksums.txt"; \
  grep "  ${asset}$" "gh_${version}_checksums.txt" | sha256sum -c -; \
  dpkg -i "$asset"; \
  rm -f "$asset" "gh_${version}_checksums.txt"

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
