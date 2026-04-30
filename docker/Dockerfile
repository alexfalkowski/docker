FROM alexfalkowski/root:2.7

USER root
WORKDIR /usr/local/bin

COPY scripts/install-image-tool /usr/local/bin/install-image-tool
COPY scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
COPY docker/scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
RUN chmod +x /usr/local/bin/install-image-tool \
  && install-image-tool hadolint v2.14.0 \
  && install-image-tool shellcheck v0.11.0
