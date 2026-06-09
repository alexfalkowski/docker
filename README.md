# 🐳 Docker

![Docker](https://www.docker.com/app/uploads/2024/01/icon-docker-square.svg)
[![CircleCI](https://circleci.com/gh/alexfalkowski/docker.svg?style=svg)](https://circleci.com/gh/alexfalkowski/docker)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

A collection of Docker images (published under `docker.io/alexfalkowski/*`) plus
a `compose.yml` for running common local dependencies.

This repository is designed to be driven via `make` targets and CI runs on CircleCI.

## 🗺️ Repository layout

Each top-level directory is an image (or runtime config used by the compose
stack):

- Image directories:
  - `docker/`, `go/`, `k8s/`, `release/`, `root/`, `ruby/`
  - Each contains:
    - `Dockerfile`
    - `Makefile` declaring `IMAGE` and `VERSION`
    - `.hadolint.yaml` (image-specific hadolint rule suppressions)
    - `scripts/install-image-tool.d/` snippets for image-specific tool
      installs, when needed
- Local dependency stack:
  - `compose.yml`
  - `grafana/`, `otelcol/`, `prometheus/`, `status/` (config mounted into
    compose services)
- Shared build targets:
  - `make/docker.mk` (used by all image Makefiles)
- Helper scripts:
  - `scripts/compose` (prefers `podman compose`, falls back to `docker compose`)
  - `scripts/clean` (prunes unused images)
  - `scripts/clean-go` (shared Go cache cleanup runner: `clean-go`)
  - `scripts/install-go-tool` (shared Go build-time installer runner:
    `install-go-tool <module> <version>`)
  - `scripts/install-image-tool` (shared image build-time installer runner:
    `install-image-tool <tool> <version>`)
  - `scripts/install-image-tool.d/` (shared install snippets used by multiple
    images)
  - `scripts/lint` (runs hadolint + shellcheck)

## ✅ Prerequisites

- GNU Make 4 or newer
  - On macOS, install GNU Make and use `gmake`; the system `/usr/bin/make`
    3.81 cannot parse the shared `bin/` make fragments.
- `git`
- `docker` (required for image build, push, and manifest targets)
- `docker` or `podman` (required for compose and clean targets)
- `hadolint` and `shellcheck` (required for `make lint`)

> [!NOTE]
> Examples below use `make`; substitute `gmake` if GNU Make is installed under
> that name.

### 📦 Submodule (required)

The top-level `Makefile` includes make fragments from the `bin/` submodule. If
`bin/` is missing, most `make` targets will fail.

> [!IMPORTANT]
> Initialize the `bin/` submodule before running repository targets in a fresh
> checkout.

```sh
git submodule sync
git submodule update --init
```

## 🛠️ Make targets

List available targets:

```sh
make
# or
make help
```

> [!TIP]
> Use `make help` to discover the current target list before guessing target
> names.

On macOS with Homebrew GNU Make:

```sh
gmake help
```

### 🔎 Lint

Lint everything (Dockerfiles + shell scripts):

```sh
make lint
```

What it does (see `scripts/lint`):

- Runs `hadolint` against every `Dockerfile` in the repo (excluding `./bin`).
- Runs `shellcheck` against:
  - `scripts/lint`, `scripts/clean`, `scripts/clean-go`, `scripts/compose`,
    `scripts/install-go-tool`, `scripts/install-image-tool`
  - `release/deploy`, `release/package`, `release/version`
  - shared `scripts/install-image-tool.d/*` snippets
  - per-image `scripts/install-image-tool.d/*` snippets

### 🏗️ Build images locally

Each image directory has a `Makefile` that sets `IMAGE` and `VERSION` and then
includes `../make/docker.mk`.
The shared build targets run from the image directory but use the repository
root as Docker build context so Dockerfiles can copy the shared installer
script.
Dockerfiles pass tool versions directly to
`install-image-tool <tool> <version>`; installer snippets do not provide
fallback versions.
Go tools are installed with `install-go-tool <module> <version>`, which runs
`go install <module>@v<version>`.
Go caches are cleaned with `clean-go`.

> [!IMPORTANT]
> Build targets use the Docker CLI directly. If you add large or sensitive
> top-level paths, update `.dockerignore` before building or publishing images
> because the repository root is the build context.

Build an image locally:

```sh
make -C go build-docker
```

That produces a local image tagged like:

- `alexfalkowski/go:<VERSION>` (based on `go/Makefile`)

Build another image:

```sh
make -C docker build-docker
```

### 🚀 Build/push platform-tagged images

The shared make targets support per-platform tags of the form
`<VERSION>.<platform>`.

> [!WARNING]
> Push targets require DockerHub authentication and publish to
> `docker.io/alexfalkowski/*`.

Build:

```sh
make -C go platform=amd64 build-platform-docker
make -C go platform=arm64 build-platform-docker
```

Push (requires DockerHub login):

```sh
make -C go platform=amd64 push-platform-docker
make -C go platform=arm64 push-platform-docker
```

### 🧩 Create multi-arch manifests

After pushing platform images, publish multi-arch manifests:

```sh
make -C go manifest-platform-docker
```

This pushes two manifests:

- `alexfalkowski/go:<VERSION>`
- `alexfalkowski/go` (equivalent to `:latest`)

## 🧭 Maintainer image updates

Image `VERSION` values are managed by the maintainer. Do not bump them unless a
release change requires it.

When updating the `root/` image, publish it in two stages:

1. Update only `root/` and bump `root/Makefile`'s `VERSION`. Use a minor bump
   unless the change alters the root image contract in a major-version-worthy
   way, including major upgrades to dependencies shipped by the root image.
2. After the new root image is published, update Dockerfiles that depend on
   `alexfalkowski/root` and bump each dependent image `VERSION` in a separate
   change. If `root` had a major bump, dependents get a major bump; otherwise
   dependents get a minor bump.

## 🧱 Compose (local dependencies)

The `compose.yml` file is intended to provide a shared set of local
dependencies used by multiple projects.

The repository wraps compose via `scripts/compose`, which:

- uses `podman compose` when `podman` is available
- otherwise uses `docker compose`

### ⬇️ Pull

```sh
make docker-pull
# pull a single service
make docker-pull service=postgres
```

### ▶️ Start/stop

```sh
make start
# start a single service
make start service=redis

make stop
```

### 📜 Logs

```sh
make logs service=postgres
```

### 📋 Included services

From `compose.yml`:

> [!NOTE]
> Published service ports are bound to `127.0.0.1` in `compose.yml`, even
> though the table below shows only `host:container` port numbers.

<!-- markdownlint-disable MD013 -->

| Service | Image | Ports (host:container) | Notes |
| --- | --- | ---: | --- |
| `postgres` | `postgres:18-trixie` | `5432:5432` | `POSTGRES_USER=test`, `POSTGRES_PASSWORD=test` |
| `redis` | `valkey/valkey:9` | `6379:6379` | |
| `aws` | `hectorvent/floci` | `4566:4566` | `SERVICES=s3,sqs,ssm` |
| `vault` | `hashicorp/vault:1.21` | `8200:8200` | dev token: `vault-plaintext-root-token` |
| `prometheus` | `prom/prometheus:v3` | `9090:9090` | depends on `mimir` |
| `mimir` | `grafana/mimir` | `9009:9009` | mounts `./grafana` |
| `loki` | `grafana/loki` | `3100:3100` | mounts `./grafana` |
| `memcached` | `memcached:1.6` | `11211:11211` | cache backend for `tempo` |
| `tempo` | `grafana/tempo:2.10.5` | `3200:3200` | depends on `memcached` |
| `otel-collector` | `otel/opentelemetry-collector-contrib` | `4317:4317`, `4318:4318` | OTLP gRPC/HTTP ingress |
| `grafana` | `grafana/grafana-oss` | `10000:3000` | depends on metrics/logs/traces stack |
| `status` | `alexfalkowski/status` | `15000:8080`, `15001:6060` | mounts `./status` |
| `flipt` | `flipt/flipt` | `8080:8080`, `9000:9000` | |

<!-- markdownlint-enable MD013 -->

Examples:

```sh
# Connect to postgres
psql "postgresql://test:test@localhost:5432/postgres"

# Vault dev login
export VAULT_ADDR="http://127.0.0.1:8200"
vault login "vault-plaintext-root-token"

# Grafana UI
open "http://127.0.0.1:10000"
```

## 🧹 Clean unused images

> [!CAUTION]
> This is destructive: it prunes all unused images from the selected engine.
> `scripts/clean` uses Podman when `podman` is available, otherwise Docker.

```sh
make clean
```

This runs `docker image prune -a -f` (or the podman equivalent) via `scripts/clean`.

## 🔁 CI (CircleCI)

CircleCI is configured as a setup workflow using path-filtering:

- `.circleci/config.yml` maps changed paths to pipeline parameters.
- `.circleci/continue_config.yml` contains the actual per-image
  build/lint/push workflows.

On `master`, CI pushes platform images and manifests (requires
`DOCKERHUB_USERNAME` and `DOCKERHUB_PASS`).
