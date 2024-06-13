FROM alexfalkowski/root:1.0

USER root
WORKDIR /usr/local/bin

# Install hadolint.
RUN curl -sSOL https://github.com/hadolint/hadolint/releases/download/v2.12.0/hadolint-Linux-x86_64 && \
  mv hadolint-Linux-x86_64 hadolint && \
  chmod +x hadolint
