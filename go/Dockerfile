FROM alexfalkowski/root:2.7

USER root
WORKDIR /usr/local/bin

ENV GO_LINT_VERSION=2.11.4

COPY scripts/install-image-tool /usr/local/bin/install-image-tool
COPY scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
COPY go/scripts/install-image-tool.d/ /usr/local/lib/install-image-tool/
RUN chmod +x /usr/local/bin/install-image-tool \
  && install-image-tool dockerize v0.9.3 \
  && install-image-tool shellcheck v0.11.0 \
  && install-image-tool hadolint v2.14.0 \
  && install-image-tool buf v1.69.0 \
  && install-image-tool trivy 0.70.0 \
  && install-image-tool mkcert v1.4.4 \
  && install-image-tool codecov v11.2.8 \
  && install-image-tool golangci-lint "$GO_LINT_VERSION"

USER circleci
WORKDIR /home/circleci/

# Setup go lint.
RUN echo "$GO_LINT_VERSION" > .go-lint-version

# Install go tools.
ENV GOEXPERIMENT=jsonv2
RUN go install github.com/alexfalkowski/gocovmerge/v2@v2.26.0 \
  && go install golang.org/x/vuln/cmd/govulncheck@v1.1.4 \
  && go install gotest.tools/gotestsum@v1.13.0 \
  && go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@v0.44.0 \
  && go install github.com/boyter/scc/v3@v3.7.0 \
  && go install github.com/Zxilly/go-size-analyzer/cmd/gsa@v1.12.5 \
  && go clean --cache -testcache -fuzzcache -modcache \
  && rm -rf go/pkg \
  && rm -rf .cache
