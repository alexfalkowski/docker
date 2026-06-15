# 🐳 Docker

![Docker](https://www.docker.com/app/uploads/2024/01/icon-docker-square.svg)
[![CircleCI](https://circleci.com/gh/alexfalkowski/docker.svg?style=svg)](https://circleci.com/gh/alexfalkowski/docker)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

Docker images published under `docker.io/alexfalkowski/*`, plus a
`compose.yml` stack for local dependencies used across projects.

The repository is Make-driven. Use `make help` for the current command list;
Dockerfiles and image `Makefile`s are the source of truth for exact tool and
image versions.

## 🗺️ Map

| Directory | Image | Purpose |
| --- | --- | --- |
| `root/` | `alexfalkowski/root` | Base image used by the other CI images. |
| `docker/` | `alexfalkowski/docker` | Dockerfile, shell, and image-security tooling. |
| `go/` | `alexfalkowski/go` | Go project CI tooling. |
| `k8s/` | `alexfalkowski/k8s` | Kubernetes and infrastructure tooling. |
| `release/` | `alexfalkowski/release` | Release automation tooling and helper commands. |
| `ruby/` | `alexfalkowski/ruby` | Ruby project CI tooling. |

Useful paths:

- `make/docker.mk`: shared build, scan, push, and manifest targets.
- `scripts/`: compose, lint, cleanup, and install helpers.
- `scripts/install-image-tool.d/`: shared installer snippets.
- `<image>/scripts/install-image-tool.d/`: image-specific installer snippets.
- `compose.yml`, `grafana/`, `otelcol/`, `prometheus/`, `status/`: local
  dependency and observability stack.

## ✅ Setup

> [!IMPORTANT]
> Initialize the shared `bin/` submodule before running repository targets.

Use GNU Make 4 or newer. On macOS, use `gmake`; `/usr/bin/make` 3.81 cannot
parse the shared `bin/` make fragments.

```sh
git submodule sync
git submodule update --init
```

The submodule URL is `git@github.com:alexfalkowski/bin.git`, so fresh
checkouts need GitHub SSH access.

Common local tools are Docker or Podman, Ruby, Hadolint, ShellCheck, and Trivy.

## 🛠️ Commands

Use a published image by pinning the version from the image directory's
`Makefile`. The unqualified image tag is a mutable `latest` tag.

```sh
docker run --rm docker.io/alexfalkowski/go:4.17 go version
```

For CI, use the same versioned image tag:

```yaml
docker:
  - image: alexfalkowski/go:4.17
```

Discover targets:

```sh
make help
```

Lint:

```sh
make lint
```

Build or scan one image:

```sh
make -C go build-docker
make -C go test-docker
```

Build, scan, push, and publish manifests:

```sh
make -C go release-docker
make -C go platform=amd64 release-platform-docker
make -C go platform=arm64 release-platform-docker
make -C go manifest-platform-docker
```

Override `DOCKER_IMAGE` when building a fork or local registry image:

```sh
make -C go DOCKER_IMAGE=example/go build-docker
```

> [!WARNING]
> Push, release, and manifest targets publish to DockerHub and require
> DockerHub credentials.

## 🧭 Maintainer Notes

Image `Makefile`s set `IMAGE` and `VERSION`, then include `../make/docker.mk`.
`VERSION` values are maintainer-managed; do not bump them unless the change is
part of a release.

Routine Docker dependency maintenance is handled by the shared
`update-docker-dep` workflow in
[`alexfalkowski/scripts`](https://github.com/alexfalkowski/scripts). Keep local
changes focused on repository-specific image contracts, installer snippets, and
validation.

Root image updates are staged:

1. Update only `root/` and bump `root/Makefile`'s `VERSION`.
2. After the root image is published, update dependent Dockerfiles that use
   `alexfalkowski/root` and bump their versions in a separate change.

Use a minor bump for ordinary root image changes. Use a major bump when the
root image contract changes in a major-version-worthy way, including major
upgrades to dependencies shipped by the root image. If root gets a major bump,
dependent images get major bumps when they move to the new root; otherwise they
get minor bumps.

The repository root is the Docker build context. Keep `.dockerignore` current
when adding large or sensitive top-level paths.

Installer snippets run through `install-image-tool <tool> <version>` and
receive the bare version as `$1`. Use helper functions from
`scripts/install-image-tool` for architecture selection, downloads, checksum
verification, release tags, and binary installs. Go module tools are installed
with `install-go-tool <module> <version>`; run `clean-go` after Go tool
installs.

## 🧱 Local Dependencies

The compose stack is managed through `scripts/compose`, which prefers
`podman compose` and falls back to `docker compose`.

```sh
make pull-latest
make start
make start service=redis
make logs service=postgres
make stop
```

`make start` runs the compose stack in detached mode and does not wait for
services to become ready. Check `make logs service=<name>` or the local
endpoints before running dependent applications.

Useful local endpoints:

| Service | Local endpoint |
| --- | --- |
| Postgres | `postgresql://test:test@localhost:5432/postgres` |
| Valkey/Redis | `redis://127.0.0.1:6379` |
| AWS emulator | `http://127.0.0.1:4566` |
| Vault | `http://127.0.0.1:8200` |
| Prometheus | `http://127.0.0.1:9090` |
| Mimir | `http://127.0.0.1:9009` |
| Loki | `http://127.0.0.1:3100` |
| Memcached | `127.0.0.1:11211` |
| Tempo | `http://127.0.0.1:3200` |
| OTLP collector | gRPC `127.0.0.1:4317`, HTTP `127.0.0.1:4318` |
| Grafana | `http://127.0.0.1:10000` |
| Status | `http://127.0.0.1:15000`, debug `http://127.0.0.1:15001` |
| Flipt | `http://127.0.0.1:8080`, `127.0.0.1:9000` |

Most compose services use disposable container-local storage. Prometheus is the
exception and uses the `prometheus_data` volume.

External applications should send telemetry to the OpenTelemetry collector:

```sh
OTEL_EXPORTER_OTLP_PROTOCOL=http/protobuf \
OTEL_EXPORTER_OTLP_ENDPOINT=http://127.0.0.1:4318 \
<app command>
```

Use `127.0.0.1:4317` for OTLP/gRPC clients. For exact services, ports, images,
mounted config, and dashboards, read `compose.yml`, `grafana/`,
`otelcol/config.yml`, and `prometheus/config.yml`.

## 🧹 Cleanup

> [!CAUTION]
> `make clean` is destructive: it runs `image prune -a -f` through
> `scripts/clean`, using Podman when `podman` is in `PATH` and Docker
> otherwise.

```sh
make clean
```

## 🔁 CI

CircleCI uses path filtering from `.circleci/config.yml`; the continued config
defines image build, publish, manifest, sync, and release jobs.

On `master`, CI publishes platform images and manifests with the CircleCI
`docker` context. On non-`master` branches, CI builds changed images and runs
`make sync push`.

Stack-only changes to `compose.yml`, `grafana/`, `otelcol/`, `prometheus/`, or
`status/` are outside the image path filters, so validate them locally with
`make start`, `make logs service=<name>`, and the relevant endpoints.
