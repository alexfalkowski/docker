# 🐳 Docker

![Docker](https://www.docker.com/app/uploads/2024/01/icon-docker-square.svg)
[![CircleCI](https://circleci.com/gh/alexfalkowski/docker.svg?style=svg)](https://circleci.com/gh/alexfalkowski/docker)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

Docker images published under `docker.io/alexfalkowski/*`, plus a
`compose.yml` stack for local dependencies used across projects.

The repository is intentionally Make-driven. Use `make help` as the source of
truth for root, local-stack, and shared git targets. Image build targets are
shared by the image directories and documented below.

## 🗺️ Map

Image directories:

| Directory | Image | Purpose |
| --- | --- | --- |
| `root/` | `alexfalkowski/root` | Base image used by the other CI images. |
| `docker/` | `alexfalkowski/docker` | Dockerfile, shell, and image-security tooling. |
| `go/` | `alexfalkowski/go` | Go project CI tooling. |
| `k8s/` | `alexfalkowski/k8s` | Kubernetes and infrastructure tooling. |
| `release/` | `alexfalkowski/release` | Release automation tooling and helper commands. |
| `ruby/` | `alexfalkowski/ruby` | Ruby project CI tooling. |

Image command surface:

| Image | Notable tools and contracts |
| --- | --- |
| `root` | Ruby, Bundler, Go, Docker CLI/buildx/compose, Git, jq, Make, SSH, and the default `circleci` user used by dependent images. |
| `docker` | `hadolint`, `shellcheck`, and `trivy` for Dockerfile, shell, and image-security checks. |
| `go` | `buf`, `codecov`, `dockerize`, `fieldalignment`, `gocovmerge`, `golangci-lint`, `govulncheck`, `gotestsum`, `gsa`, `hadolint`, `mkcert`, `scc`, `shellcheck`, and `trivy`; sets `GOEXPERIMENT=jsonv2`. |
| `k8s` | `doctl`, `kubectl`, `kubescape`, `pulumi`, `vegeta`, and `kube-score`. |
| `release` | `gh`, `goreleaser`, `uplift`, `bump`, and the `version`, `package`, and `deploy` helpers. |
| `ruby` | `buf`, `codecov`, `dockerize`, `hadolint`, `mkcert`, `scc`, `shellcheck`, and `trivy`. |

Dockerfiles are the source of truth for exact pinned tool versions.

Other useful paths:

- `make/docker.mk`: shared image build, scan, push, and manifest targets.
- `scripts/`: compose wrapper, lint wrapper, cleanup, and install helpers.
- `scripts/install-image-tool.d/`: shared image install snippets.
- `<image>/scripts/install-image-tool.d/`: image-specific install snippets.
- `compose.yml`: local dependency stack.
- `grafana/`, `otelcol/`, `prometheus/`, `status/`: config mounted by the
  compose stack.
- `.circleci/`: path-filtered image build and publish workflows.

## ✅ Setup

> [!IMPORTANT]
> Initialize the shared `bin/` submodule before running repository targets.

Use GNU Make 4 or newer. On macOS, use `gmake` if Homebrew installed GNU Make
under that name; `/usr/bin/make` 3.81 cannot parse the shared `bin/` make
fragments.

Initialize the shared `bin/` submodule before running repository targets:

```sh
git submodule sync
git submodule update --init
```

The submodule URL is `git@github.com:alexfalkowski/bin.git`, so fresh
checkouts need GitHub SSH access.

Common local tools:

- `docker`: image builds, scans, pushes, and manifests.
- `podman` or `docker`: local compose stack and image cleanup.
- `ruby`: shared `bin/` make fragments.
- `hadolint` and `shellcheck`: `make lint`.
- `trivy`: image scan targets.

## 🛠️ Commands

Discover targets:

```sh
make help
```

Image directories do not expose their own `help` target; use the examples below
for the shared image target names.

Lint:

```sh
make lint
```

Build or scan one image:

```sh
make -C go build-docker
make -C go test-docker
```

Push the versioned and unqualified `latest` image tags after a successful local
build and scan:

```sh
make -C go release-docker
```

Build or scan one platform image:

```sh
make -C go platform=amd64 build-platform-docker
make -C go platform=amd64 test-platform-docker
```

Publish one platform image and then publish the versioned and unqualified
`latest` multi-arch manifests:

```sh
make -C go platform=amd64 release-platform-docker
make -C go platform=arm64 release-platform-docker
make -C go manifest-platform-docker
```

> [!WARNING]
> Push, release, and manifest targets publish to DockerHub and require
> DockerHub credentials.

Image scans are implemented in `make/docker.mk` and currently run Trivy
OS-package vulnerability checks.

## 🧭 Maintainer Notes

Image `Makefile`s set `IMAGE` and `VERSION`, then include `../make/docker.mk`.
`VERSION` values are maintainer-managed; do not bump them unless the change is
part of a release.

> [!IMPORTANT]
> Root image updates are staged so dependent images can move after the new root
> image is published.

Root image updates are staged:

1. Update only `root/` and bump `root/Makefile`'s `VERSION`.
2. After that root image is published, update dependent Dockerfiles that use
   `alexfalkowski/root` and bump their versions in a separate change.

Use a minor bump for ordinary root image changes. Use a major bump when the
root image contract changes in a major-version-worthy way, including major
upgrades to dependencies shipped by the root image.

The root image contract is the runtime surface inherited by dependent images:
the Ruby slim Debian base, Go toolchain and `GOPATH`/`PATH` setup, Docker CLI
tooling, RubyGems and Bundler, the passwordless-sudo `circleci` user, and the
default `/home/circleci/` working directory. Treat changes to those surfaces as
compatibility changes for root and dependent images.

The repository root is the Docker build context. Keep `.dockerignore` current
when adding large or sensitive top-level paths.

Dockerfiles install pinned archive tools with
`install-image-tool <tool> <version>`. Put shared installer snippets in
`scripts/install-image-tool.d/` and image-specific snippets in
`<image>/scripts/install-image-tool.d/`; the snippet name is the tool name and
receives the bare version as `$1`. Snippets run from a temporary directory and
can use the helper functions from `scripts/install-image-tool` for architecture
selection, downloads, checksum verification, release tags, and binary installs.
Install pinned Go module tools with `install-go-tool <module> <version>`, which
adds the `v` prefix for `go install`. Run `clean-go` after Go tool installs to
remove Go build caches, module cache, `go/pkg`, and `.cache`.

The `go/` image sets `GOEXPERIMENT=jsonv2` globally. Downstream projects that
need default Go experiment behavior should override it for that command, for
example `GOEXPERIMENT= go test ./...`.

The `release/` image contains the `version`, `package`, and `deploy` helper
commands. They share `APP_VERSION_FILE`, defaulting to
`/tmp/workspace/release-version.txt`, and expect to run from the repository
being released.

| Command | Behavior |
| --- | --- |
| `version` | Runs `uplift release --skip-changelog --config-dir /etc/uplift`, removes any stale version file first, and writes the new tag to `APP_VERSION_FILE` only when a new tag points at `HEAD`. |
| `package` | Runs `goreleaser release` only when the working tree contains a path whose name includes `goreleaser` outside `vendor/` and `APP_VERSION_FILE` is non-empty. Otherwise it exits without publishing. |
| `deploy` | Runs only when the repository contains `.cd` and `APP_VERSION_FILE` exists. It expects the version file to contain the tag written by `version`, derives the app name from the current directory, clones `alexfalkowski/infraops` over SSH into `./infraops`, bumps the released app version, and raises the infraops change through `make ready`. Failed runs can leave `./infraops` behind; remove it before rerunning. |

Use the release helpers only in an authenticated release environment. They can
create tags, publish releases, clone over SSH, push branches, and raise remote
changes.

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

Useful local entry points:

| Service | Local endpoint | Use |
| --- | --- | --- |
| Postgres | `postgresql://test:test@localhost:5432/postgres` | Application database. |
| Valkey/Redis | `redis://127.0.0.1:6379` | Redis-compatible cache or queue dependency. |
| AWS emulator | `http://127.0.0.1:4566` | Local S3, SQS, and SSM. |
| Vault | `http://127.0.0.1:8200` | Development Vault with token `vault-plaintext-root-token`. |
| Prometheus | `http://127.0.0.1:9090` | Local metrics scraping and query UI. |
| Mimir | `http://127.0.0.1:9009` | Prometheus remote-write backend. |
| Loki | `http://127.0.0.1:3100` | OTLP log backend. |
| Memcached | `127.0.0.1:11211` | Tempo query-frontend cache. |
| Tempo | `http://127.0.0.1:3200` | Trace backend and query API. |
| OTLP collector | gRPC `127.0.0.1:4317`, HTTP `127.0.0.1:4318` | Metrics, logs, and traces ingest. |
| Grafana | `http://127.0.0.1:10000` | Dashboard UI. |
| Status | `http://127.0.0.1:15000`, debug `http://127.0.0.1:15001` | Example service scraped by Prometheus. |
| Flipt | `http://127.0.0.1:8080`, `127.0.0.1:9000` | Feature flag service APIs. |

Most compose services do not declare named volumes, so treat their data as
disposable. Prometheus is the exception and uses `prometheus_data` for its
local TSDB. Mimir, Loki, and Tempo use container-local filesystem storage in
this stack, so remote-written metrics, logs, and traces are disposable when
their containers are recreated unless you add storage for them.

Observability flow:

- Prometheus scrapes `prometheus` and `status`, then remote-writes samples to
  Mimir.
- The OpenTelemetry collector receives OTLP metrics, logs, and traces, then
  sends metrics to Mimir, logs to Loki, and traces to Tempo.
- The bundled `status` service exposes Prometheus metrics and sends traces
  directly to Tempo. Use the collector endpoints when testing collector ingest
  from external clients.
- Tempo generates service graph and span metrics, then remote-writes them back
  to Prometheus.

Grafana is not provisioned with datasources or dashboards by `compose.yml`.
It uses the upstream Grafana login defaults unless you override them locally.
Create datasources manually before importing dashboards from `grafana/`:

| Datasource | Type | URL from Grafana |
| --- | --- | --- |
| `prometheus` | Prometheus | `http://prometheus:9090` |
| `loki` | Loki | `http://loki:3100` |
| `tempo` | Tempo | `http://tempo:3200` |

The bundled service, Go metrics, and Prometheus dashboards expect the
`prometheus` datasource. `grafana/alertmanager.json` requires an Alertmanager
service and Prometheus scrape target, which this compose stack does not start by
default.

For exact services, ports, images, and mounted config, read `compose.yml`.
For dashboard imports and observability config, start with `grafana/`,
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

CircleCI uses path filtering:

- `.circleci/config.yml` maps changed paths to image pipeline parameters.
- `.circleci/continue_config.yml` defines the image build, publish, manifest,
  sync, and release jobs.

On `master`, CI publishes platform images and manifests with the CircleCI
`docker` context. The `version` job uses the `gh` context. On non-`master`
branches, CI builds the changed images and then runs `make sync push`.
