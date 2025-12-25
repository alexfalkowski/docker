FROM alexfalkowski/root:1.58

USER root

# Add sources.
RUN echo 'deb [trusted=yes] https://cli.github.com/packages stable main' | tee /etc/apt/sources.list.d/github.list
RUN echo 'deb [trusted=yes] https://fury.upliftci.dev/apt/ /' | tee /etc/apt/sources.list.d/uplift.list
RUN echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | tee /etc/apt/sources.list.d/goreleaser.list

RUN apt-get update && apt-get install --no-install-recommends -y \
  gh=2.83.2 \
  goreleaser=2.13.2 \
  uplift=2.26.0 \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

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

# Setup git.
RUN git config --global user.email "ci@lean-thoughts.com" \
  && git config --global user.name "lean-thoughts-ci"

# Install go tools.
RUN go install github.com/alexfalkowski/infraops/v2/cmd/bump@v2.368.0 \
  && go clean --cache -testcache -fuzzcache -modcache && rm -rf "${GOPATH}/pkg"
