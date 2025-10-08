FROM debian:trixie AS build

# Install packages.
RUN apt-get update && apt-get install --no-install-recommends -y \
  curl \
  ca-certificates \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

WORKDIR /tmp

# Install monitoror.
ENV MONITOROR_VERSION=4.0.1
RUN curl -sSOL https://github.com/monitoror/monitoror/releases/download/$MONITOROR_VERSION/monitoror-linux-amd64-$MONITOROR_VERSION \
  && mv monitoror-linux-amd64-$MONITOROR_VERSION monitoror \
  && chmod +x monitoror

FROM gcr.io/distroless/static

WORKDIR /

COPY config.json /
ENV MO_CONFIG="/config.json"

COPY --from=build /tmp/monitoror /monitoror

ENTRYPOINT ["/monitoror"]
