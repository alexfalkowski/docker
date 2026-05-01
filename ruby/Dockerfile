FROM alexfalkowski/root:2.8

USER root
WORKDIR /usr/local/bin

COPY scripts/install-image-tool /usr/local/bin/install-image-tool
COPY scripts/install-go-tool /usr/local/bin/install-go-tool
COPY scripts/clean-go /usr/local/bin/clean-go
COPY scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
COPY ruby/scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
RUN chmod +x /usr/local/bin/install-image-tool \
  && chmod +x /usr/local/bin/install-go-tool \
  && chmod +x /usr/local/bin/clean-go \
  && install-image-tool dockerize v0.9.3 \
  && install-image-tool shellcheck v0.11.0 \
  && install-image-tool hadolint v2.14.0 \
  && install-image-tool buf v1.69.0 \
  && install-image-tool trivy 0.70.0 \
  && install-image-tool mkcert v1.4.4 \
  && install-image-tool codecov v11.2.8

USER circleci
WORKDIR /home/circleci/

# Install go tools.
RUN install-go-tool github.com/boyter/scc/v3 v3.7.0 \
  && clean-go
