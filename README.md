# 🐳 Docker

![Docker](https://www.docker.com/app/uploads/2024/01/icon-docker-square.svg)
[![CircleCI](https://circleci.com/gh/alexfalkowski/docker.svg?style=svg)](https://circleci.com/gh/alexfalkowski/docker)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

Docker images published under `docker.io/alexfalkowski/*`, plus a
`compose.yml` stack for local dependencies used across projects.

The repository is intentionally Make-driven. Use `make help` as the source of
truth for available local targets.

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

Lint:

```sh
make lint
```

Build or scan one image:

```sh
make -C go build-docker
make -C go test-docker
```

Push one image after a successful local build and scan:

```sh
make -C go release-docker
```

Build or scan one platform image:

```sh
make -C go platform=amd64 build-platform-docker
make -C go platform=amd64 test-platform-docker
```

Publish one platform image and then publish the multi-arch manifests:

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

The repository root is the Docker build context. Keep `.dockerignore` current
when adding large or sensitive top-level paths.

The `release/` image contains the `version`, `package`, and `deploy` helper
commands. Read those scripts before using them locally; they can create tags,
publish releases, clone over SSH, and raise remote changes.

## 🧱 Local Dependencies

The compose stack is managed through `scripts/compose`, which prefers
`podman compose` and falls back to `docker compose`.

```sh
make docker-pull
make start
make start service=redis
make logs service=postgres
make stop
```

Useful local entry points:

- Postgres: `postgresql://test:test@localhost:5432/postgres`
- Vault: `http://127.0.0.1:8200` with token `vault-plaintext-root-token`
- Grafana: `http://127.0.0.1:10000`
- OTLP: gRPC `127.0.0.1:4317`, HTTP `127.0.0.1:4318`

Most compose services do not declare named volumes, so treat their data as
disposable. Prometheus is the exception and uses `prometheus_data`.

For exact services, ports, images, and mounted config, read `compose.yml`.
For dashboard imports and observability config, start with `grafana/`,
`otelcol/config.yml`, and `prometheus/config.yml`.

## 🧹 Cleanup

> [!CAUTION]
> `make clean` is destructive: it prunes unused images from the selected Docker
> or Podman engine.

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
