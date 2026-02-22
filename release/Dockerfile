FROM alexfalkowski/root:1.65

USER root

# Add sources.
RUN echo 'deb [trusted=yes] https://fury.upliftci.dev/apt/ /' | tee /etc/apt/sources.list.d/uplift.list
RUN echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | tee /etc/apt/sources.list.d/goreleaser.list

RUN apt-get update && apt-get install --no-install-recommends -y \
  goreleaser=2.14.0 \
  uplift=2.26.0 \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install gh from a downloaded .deb.
RUN set -eux; \
  version="2.87.2"; \
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
