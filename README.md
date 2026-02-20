![Docker](https://www.docker.com/app/uploads/2024/01/icon-docker-square.svg)
[![CircleCI](https://circleci.com/gh/alexfalkowski/docker.svg?style=svg)](https://circleci.com/gh/alexfalkowski/docker)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Docker

A collection of Docker images (published under `docker.io/alexfalkowski/*`) plus a `compose.yml` for running common local dependencies.

This repository is designed to be driven via `make` targets and CI runs on CircleCI.

## Repository layout

Each top-level directory is an image (or runtime config used by the compose stack):

- Image directories:
  - `docker/`, `go/`, `k8s/`, `monitoror/`, `release/`, `root/`, `ruby/`
  - Each contains:
    - `Dockerfile`
    - `Makefile` declaring `IMAGE` and `VERSION`
    - `.hadolint.yaml` (image-specific hadolint rule suppressions)
- Local dependency stack:
  - `compose.yml`
  - `grafana/`, `prometheus/`, `status/` (config mounted into compose services)
- Shared build targets:
  - `make/docker.mk` (used by all image Makefiles)
- Helper scripts:
  - `scripts/compose` (prefers `podman compose`, falls back to `docker compose`)
  - `scripts/clean` (prunes unused images)
  - `scripts/lint` (runs hadolint + shellcheck)

## Prerequisites

- `make`
- `git`
- `docker` or `podman`
- `hadolint` and `shellcheck` (required for `make lint`)

### Submodule (required)

The top-level `Makefile` includes make fragments from the `bin/` submodule. If `bin/` is missing, most `make` targets will fail.

```sh
git submodule sync
git submodule update --init
```

## Make targets

List available targets:

```sh
make
# or
make help
```

### Lint

Lint everything (Dockerfiles + shell scripts):

```sh
make lint
```

What it does (see `scripts/lint`):

- Runs `hadolint` against every `Dockerfile` in the repo (excluding `./bin`).
- Runs `shellcheck` against:
  - `scripts/lint`, `scripts/clean`, `scripts/compose`
  - `release/deploy`, `release/package`, `release/version`

### Build images locally

Each image directory has a `Makefile` that sets `IMAGE` and `VERSION` and then includes `../make/docker.mk`.

Build an image locally:

```sh
make -C go build-docker
```

That produces a local image tagged like:

- `alexfalkowski/go:2.109` (based on `go/Makefile`)

Build another image:

```sh
make -C docker build-docker
```

### Build/push platform-tagged images

The shared make targets support per-platform tags of the form `<VERSION>.<platform>`.

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

### Create multi-arch manifests

After pushing platform images, publish multi-arch manifests:

```sh
make -C go manifest-platform-docker
```

This pushes two manifests:

- `alexfalkowski/go:2.109`
- `alexfalkowski/go` (equivalent to `:latest`)

## Compose (local dependencies)

The `compose.yml` file is intended to provide a shared set of local dependencies used by multiple projects.

The repository wraps compose via `scripts/compose`, which:

- uses `podman compose` when `podman` is available
- otherwise uses `docker compose`

### Pull

```sh
make docker-pull
# pull a single service
make docker-pull service=postgres
```

### Start/stop

```sh
make start
# start a single service
make start service=redis

make stop
```

### Logs

```sh
make logs service=postgres
```

### Included services

From `compose.yml`:

| Service | Image | Ports (host:container) | Notes |
|---|---|---:|---|
| `postgres` | `postgres:17-bullseye` | `5432:5432` | `POSTGRES_USER=test`, `POSTGRES_PASSWORD=test` |
| `redis` | `redis:8` | `6379:6379` |  |
| `localstack` | `localstack/localstack:4.0` | `4566:4566` | `SERVICES=s3,sqs,ssm` |
| `vault` | `hashicorp/vault:1.18` | `8200:8200` | dev token: `vault-plaintext-root-token` |
| `prometheus` | `prom/prometheus:v3.1.0` | `9090:9090` | depends on `mimir` |
| `mimir` | `grafana/mimir` | `9009:9009` |  |
| `loki` | `grafana/loki` | `3100:3100` |  |
| `tempo` | `grafana/tempo` | `3200:3200`, `4317:4317`, `4318:4318` | OTLP gRPC/HTTP exposed |
| `grafana` | `grafana/grafana-oss` | `10000:3000` | depends on metrics/logs/traces stack |
| `status` | `alexfalkowski/status` | `15000:8080`, `15001:6060` | mounts `./status/server.yml` |
| `flipt` | `flipt/flipt` | `8080:8080`, `9000:9000` |  |

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

## Clean unused images

Warning: this is destructive.

```sh
make clean
```

This runs `docker image prune -a -f` (or the podman equivalent) via `scripts/clean`.

## CI (CircleCI)

CircleCI is configured as a setup workflow using path-filtering:

- `.circleci/config.yml` maps changed paths to pipeline parameters.
- `.circleci/continue_config.yml` contains the actual per-image build/lint/push workflows.

On `master`, CI pushes platform images and manifests (requires `DOCKERHUB_USERNAME` and `DOCKERHUB_PASS`).
