FROM alexfalkowski/root:3.10

USER root
WORKDIR /usr/local/bin

ENV GO_LINT_VERSION=2.12.2

COPY scripts/install-image-tool /usr/local/bin/install-image-tool
COPY scripts/install-go-tool /usr/local/bin/install-go-tool
COPY scripts/clean-go /usr/local/bin/clean-go
COPY scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
COPY go/scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
RUN chmod +x /usr/local/bin/install-image-tool \
  && chmod +x /usr/local/bin/install-go-tool \
  && chmod +x /usr/local/bin/clean-go \
  && install-image-tool dockerize 0.9.3 \
  && install-image-tool shellcheck 0.11.0 \
  && install-image-tool hadolint 2.14.0 \
  && install-image-tool buf 1.71.0 \
  && install-image-tool trivy 0.71.2 \
  && install-image-tool mkcert 1.4.4 \
  && install-image-tool codecov 11.2.8 \
  && install-image-tool golangci-lint "$GO_LINT_VERSION"

USER circleci
WORKDIR /home/circleci/

# Setup go lint.
RUN echo "$GO_LINT_VERSION" > .go-lint-version

# Install go tools.
ENV GOEXPERIMENT=jsonv2
RUN install-go-tool github.com/alexfalkowski/gocovmerge/v2 2.29.2 \
  && install-go-tool golang.org/x/vuln/cmd/govulncheck 1.3.0 \
  && install-go-tool gotest.tools/gotestsum 1.13.0 \
  && install-go-tool golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment 0.44.0 \
  && install-go-tool github.com/boyter/scc/v3 3.7.0 \
  && install-go-tool github.com/Zxilly/go-size-analyzer/cmd/gsa 1.13.0 \
  && clean-go
