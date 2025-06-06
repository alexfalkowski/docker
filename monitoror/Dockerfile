FROM alexfalkowski/root:1.37 AS build

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
