# AGENTS.md

## Shared skills

This repository uses the shared skills from `bin/skills/`. Read
`bin/AGENTS.md` for the canonical shared skill list and use the smallest
matching skill for the task.

Docker images plus a `compose.yml` local dependency stack.

## Basics

- Use GNU Make 4+. On macOS, use `gmake`; `/usr/bin/make` 3.81 cannot parse the shared `bin/` make fragments.
- Initialize the required submodule with `git submodule sync && git submodule update --init`.
- Show targets with `make help` or `gmake help`.

## Map

- Images: `docker/`, `go/`, `k8s/`, `release/`, `root/`, `ruby/`.
- Shared image targets: `make/docker.mk`; shared scripts: `scripts/`.
- Shared install snippets: `scripts/install-image-tool.d/`; image-specific snippets: `<image>/scripts/install-image-tool.d/`.
- Compose config: `compose.yml`, `grafana/`, `otelcol/`, `prometheus/`, `status/`.
- CI: `.circleci/`.

## Commands

- `make lint`: lint Dockerfiles and shell scripts.
- `make -C <dir> build-docker` / `lint-docker`: build or lint one image.
- `make -C <dir> platform=amd64 build-platform-docker`: build a platform image.
- `make -C <dir> manifest-platform-docker`: publish the multi-arch manifests.
- `make docker-pull`, `make start`, `make stop`, `make logs service=<name>`: manage compose services.

## Rules

- Image `Makefile`s set `IMAGE` and `VERSION`, then include `../make/docker.mk`.
- Image `VERSION` values are managed by the maintainer; do not bump them unless explicitly asked.
- `make/docker.mk` builds from the repo root context with `docker build -f Dockerfile ... ..`.
- Updating `root/` is a two-stage process driven by the maintainer:
  1. First update only `root/` and bump `root/Makefile`'s `VERSION`; use a minor bump unless the change alters the root image contract in a major-version-worthy way, including major upgrades to dependencies shipped by the root image.
  2. After the new root image is published, update the Dockerfiles that depend on `alexfalkowski/root` and bump each dependent image `VERSION` in a separate change. If root had a major bump, dependents get a major bump; otherwise dependents get a minor bump.
- Keep `.dockerignore` current when adding large or sensitive top-level paths.
- Dockerfiles call `install-image-tool <tool> <version>` and `install-go-tool <module> <version>`; run `clean-go` after Go tool installs.
- Hadolint suppressions live in each image directory's `.hadolint.yaml`; there is no top-level hadolint config.
- `release/` installs `gh`, `goreleaser`, and `uplift`, and copies `release/deploy`, `release/package`, `release/version`, and `release/.uplift.yml`.
- `release/.uplift.yml` uses release commits shaped like `chore(release): $VERSION [ci skip]`.

## Gotchas

- `.gitmodules` uses the SSH URL `git@github.com:alexfalkowski/bin.git`.
- Push and manifest targets require DockerHub credentials.
- `scripts/compose` prefers `podman compose` over `docker compose`.
- `make clean` is destructive: it prunes all unused Docker or Podman images.
- Mutable base image tags are intentional in this repository; do not flag tag-based `FROM` lines as review findings unless asked to make builds digest-pinned.
- The images intentionally allow root-level operations where needed, including for tools such as `mkcert`; do not flag the root/sudo model as a review finding unless the task is specifically about hardening runtime privileges.
