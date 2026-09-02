FROM alexfalkowski/root:3.18

USER root
WORKDIR /usr/local/bin

COPY scripts/install-image-tool /usr/local/bin/install-image-tool
COPY scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
COPY docker/scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
RUN chmod +x /usr/local/bin/install-image-tool \
  && install-image-tool hadolint 2.15.1 \
  && install-image-tool shellcheck 0.11.0 \
  && install-image-tool trivy 0.74.0
