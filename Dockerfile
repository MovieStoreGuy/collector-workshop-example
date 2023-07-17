# Builder Base will contain all the required tooling in order to build
# A custom distrobution of the collector and allow for caching.
FROM golang:1.20.6 AS builder-base

RUN go install go.opentelemetry.io/collector/cmd/builder@v0.80.0

# Collector Builder will load all configuration and any local files
# that are to be part of the build so that they can also be included.
# this step is expected to be the largest since modules are dynamically fetched.
FROM builder-base AS collector-builder

WORKDIR /src
COPY . . 

ENV CGO_ENABLED=0 GOOS=linux

RUN builder --config=build-config.yml

# Final is the release of the built open telemetry binary
# shipped as fairly contained docker image with no system depedancies.
# For components that require additional packages being installed such
# as javametricsreceiver or collectd, then additional efforts are required
# in order to get those packages installed.
FROM gcr.io/distroless/static:nonroot AS final

COPY --from=collector-builder /src/dist/otelcol /otelcol

ENTRYPOINT [ "/otelcol" ]