# AGENTS.md

## Shared guidance

Use `bin/AGENTS.md` for shared skills and cross-repository defaults.

Docker images plus a `compose.yml` local dependency stack.

## Basics

- Show targets with `make help` or `gmake help`.
- Image subdirectories also include the shared help target, so bare
  `make -C <image-dir>` prints help instead of building an image.

## Map

- Images: `docker/`, `go/`, `k8s/`, `release/`, `root/`, `ruby/`.
- Shared image targets: `make/docker.mk`; shared scripts: `scripts/`.
- Shared install snippets: `scripts/install-image-tool.d/`; image-specific snippets: `<image>/scripts/install-image-tool.d/`.
- Compose config: `compose.yml`, `grafana/`, `otelcol/`, `prometheus/`, `status/`.
- CI: `.circleci/`.

## Commands

- `make lint`: lint Dockerfiles and shell scripts.
- `make -C <dir> build-docker` / `test-docker` / `release-docker`: build, build+scan, or build+scan+push the versioned and unqualified `latest` tags for one image.
- `make -C <dir> platform=amd64 build-platform-docker` / `test-platform-docker` / `release-platform-docker`: build, build+scan, or build+scan+push one versioned platform image tag.
- `make -C <dir> manifest-platform-docker`: publish the versioned and unqualified `latest` multi-arch manifests.
- `make pull-latest`, `make start`, `make stop`, `make logs service=<name>`: manage compose services.

## Rules

- Image `Makefile`s set `IMAGE` and `VERSION`, then include `../make/docker.mk`.
- Image `VERSION` values are managed by the maintainer; do not bump them unless explicitly asked.
- `make/docker.mk` owns Docker build invocation from the repo root context.
- The root `trivy-repo` target intentionally delegates to the shared
  `bin/build/sec/trivy-repo` helper without including a language-specific
  shared Make fragment. Do not replace that by including `go.mak`, `ruby.mak`,
  or `_service.mak` in this repository root.
- Updating `root/` is a two-stage process driven by the maintainer:
  1. First update only `root/` and bump `root/Makefile`'s `VERSION`; use a minor bump unless the change alters the root image contract in a major-version-worthy way, including major upgrades to dependencies shipped by the root image.
  2. After the new root image is published, update the Dockerfiles that depend on `alexfalkowski/root` and bump each dependent image `VERSION` in a separate change. If root had a major bump, dependents get a major bump; otherwise dependents get a minor bump.
- Keep `.dockerignore` current when adding large or sensitive top-level paths.
- Dockerfiles call `install-image-tool <tool> <version>` and `install-go-tool <module> <version>`; run `clean-go` after Go tool installs. `install-image-tool` sources `/usr/local/lib/install-image-tool/<tool>` from a temporary directory with the bare version as `$1`; snippets should use the helper functions in `scripts/install-image-tool` for architecture selection, downloads, checksum verification, release tags, and binary installs.
- Hadolint suppressions live in each image directory's `.hadolint.yaml`; there is no top-level hadolint config.
- `release/` installs `gh`, `goreleaser`, and `uplift`, and copies `release/deploy`, `release/package`, `release/version`, and `release/.uplift.yml`.
- `release/.uplift.yml` uses release commits shaped like `chore(release): $VERSION [ci skip]`.

## Gotchas

- Push, release, and manifest targets require DockerHub credentials.
- `scripts/compose` prefers `podman compose` over `docker compose`.
- `make clean` is destructive: `scripts/clean` runs `image prune -a -f`, preferring Podman when `podman` is in `PATH` and falling back to Docker.
- Mutable base image tags are intentional in this repository; do not flag tag-based `FROM` lines as review findings unless asked to make builds digest-pinned.
- The images intentionally allow root-level operations where needed, including for tools such as `mkcert`; do not flag the root/sudo model as a review finding unless the task is specifically about hardening runtime privileges.
