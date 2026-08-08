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

| Directory | Image | Purpose | Default user |
| --- | --- | --- | --- |
| `root/` | `alexfalkowski/root` | Shared Ruby, Go, Docker, build, and Git base for the other CI images. | `circleci` |
| `docker/` | `alexfalkowski/docker` | Hadolint, ShellCheck, and Trivy tooling for Dockerfile and image work. | `root` |
| `go/` | `alexfalkowski/go` | Go analysis, test, security, lint, and coverage tooling. | `circleci` |
| `k8s/` | `alexfalkowski/k8s` | `kubectl`, Pulumi, `doctl`, Kubescape, kube-score, and Vegeta. | `circleci` |
| `release/` | `alexfalkowski/release` | `gh`, GoReleaser, Uplift, and release helper commands. | `circleci` |
| `ruby/` | `alexfalkowski/ruby` | Ruby CI tooling plus common Docker, Buf, and code-quality tools. | `circleci` |

The `circleci` user has passwordless `sudo`. Dockerfiles and image `Makefile`s
remain the source of truth for exact tool and image versions.

Useful paths:

- `make/docker.mk`: shared build, scan, push, and manifest targets.
- `scripts/`: compose, lint, cleanup, and install helpers.
- `scripts/install-image-tool.d/`: shared installer snippets.
- `<image>/scripts/install-image-tool.d/`: image-specific installer snippets.
- `compose.yml`, `grafana/`, `prometheus/`, `status/`: local
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
docker run --rm docker.io/alexfalkowski/go:<version> go version
```

For CI, use the same versioned image tag:

```yaml
docker:
  - image: alexfalkowski/go:<version>
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

### Release Helpers

The release image provides ordered helpers that share `APP_VERSION_FILE`
(default: `/tmp/workspace/release-version.txt`):

- `version` runs Uplift and, when it creates a new tag, writes it to the file.
- `package` runs GoReleaser only when `./.goreleaser.yml` and a non-empty
  version file exist.
- `deploy` runs only when `.cd` and the version file exist; it clones
  `alexfalkowski/infraops`, bumps the matching image version, and invokes
  `make ready` there.

> [!WARNING]
> These helpers can create tags, publish releases, and open a version-bump PR.
> Run them only in an authenticated release environment with GitHub SSH and
> release credentials.

## 🧱 Local Dependencies

The compose stack is managed through `scripts/compose`, which prefers
`podman compose` and falls back to `docker compose`.

```sh
make stack-config
make pull-latest
make start
make start service=redis
make restart service=lgtm
make logs service=postgres
make stop
```

`make start` runs the compose stack in detached mode and does not wait for
services to become ready. Check `make logs service=<name>` or the local
endpoints before running dependent applications.

Restart a running service after changing one of its bind-mounted configuration
files. For example, run `make restart service=lgtm` after updating
`prometheus/config.yml` so it reloads its scrape configuration.

Useful local endpoints:

| Service | Local endpoint |
| --- | --- |
| Postgres | `postgresql://test:test@localhost:5432/postgres` |
| Valkey/Redis | `redis://127.0.0.1:6379` |
| AWS emulator | `http://127.0.0.1:4566` |
| Vault | `http://127.0.0.1:8200` |
| Prometheus | `http://127.0.0.1:9090` |
| Loki | `http://127.0.0.1:3100` |
| Tempo | `http://127.0.0.1:3200` |
| LGTM OpenTelemetry Collector | gRPC `127.0.0.1:4317`, HTTP `127.0.0.1:4318` |
| Grafana | `http://127.0.0.1:10000` |
| Status | `http://127.0.0.1:15000`, debug `http://127.0.0.1:15001` |
| Flipt | `http://127.0.0.1:8080`, `127.0.0.1:9000` |

Most compose services use disposable container-local storage. The LGTM stack
persists Grafana, Prometheus, Loki, and Tempo data in the `lgtm_data` volume.
It is intended for local development, demos, and testing; it does not provide a
Mimir endpoint. This intentional internal-only migration does not import the
previous `prometheus_data` volume. Keep that volume until its historical metrics
are no longer needed, then remove it with your container runtime.

External applications should send telemetry to the OpenTelemetry Collector
bundled with LGTM:

```sh
OTEL_EXPORTER_OTLP_PROTOCOL=http/protobuf \
OTEL_EXPORTER_OTLP_ENDPOINT=http://127.0.0.1:4318 \
<app command>
```

Use `127.0.0.1:4317` for OTLP/gRPC clients. For exact services, ports, images,
and dashboards, read `compose.yml`, `grafana/`, and `prometheus/config.yml`.
Status sends its metrics and traces to LGTM over OTLP; Prometheus does not
scrape Status directly.

LGTM provisions the Grafana Prometheus, Loki, Tempo, and Pyroscope data sources
automatically. To use an included dashboard, open the Grafana endpoint and
import a dashboard JSON file from `grafana/`.

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

On non-`master` branches, CI builds changed images, validates image jobs after
CircleCI config changes, and runs `make sync push`. On `master`, CI publishes
platform images and manifests only for images whose version-bearing `Makefile`
changed, using the CircleCI `docker` context.

Stack-only changes to `compose.yml`, `grafana/`, `prometheus/`, or `status/`
are outside the image path filters, but still trigger the `stack`
path filter, which runs `make stack-config` in CI. That only validates the
compose config syntax, so also validate stack changes locally with
`make start`, `make logs service=<name>`, and the relevant endpoints.
