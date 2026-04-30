FROM alexfalkowski/root:2.8

USER root

COPY scripts/install-image-tool /usr/local/bin/install-image-tool
COPY release/scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
RUN chmod +x /usr/local/bin/install-image-tool \
  && install-image-tool goreleaser 2.15.4 \
  && install-image-tool uplift 2.26.0 \
  && install-image-tool gh 2.92.0

# Add uplift config.
RUN mkdir /etc/uplift
COPY release/.uplift.yml /etc/uplift

# Copy scripts.
COPY release/deploy /usr/local/bin/
RUN chmod +x /usr/local/bin/deploy

COPY release/package /usr/local/bin/
RUN chmod +x /usr/local/bin/package

COPY release/version /usr/local/bin/
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
