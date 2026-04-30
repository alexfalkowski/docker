# AGENTS.md

This repo builds Docker images plus a `compose.yml` local dependency stack.

Use the shared `coding-standards` skill from `bin/skills/coding-standards` for code changes, bug fixes, refactors, reviews, tests, linting, documentation, PR summaries, commits, Makefile changes, CI validation, and verification.

## Setup

- The `bin/` submodule is required for most `make` targets:

  ```sh
  git submodule sync
  git submodule update --init
  ```

- Show targets with `make` or `make help`.

## Layout

- Image directories: `docker/`, `go/`, `k8s/`, `release/`, `root/`, `ruby/`.
- Each image directory has `Dockerfile`, `Makefile`, `.hadolint.yaml`, and may have `scripts/install-image-tool.d/`.
- Shared image build targets live in `make/docker.mk`.
- Shared scripts live in `scripts/`; `scripts/install-image-tool` is the common runner used by Dockerfiles.
- Shared install snippets used by multiple images live in `scripts/install-image-tool.d/`.
- CircleCI path-filtering and workflows live under `.circleci/`.

## Commands

- Lint all Dockerfiles and shell scripts:

  ```sh
  make lint
  ```

- Build or lint one image:

  ```sh
  make -C <dir> build-docker
  make -C <dir> lint-docker
  ```

- Platform images and manifests:

  ```sh
  make -C <dir> platform=amd64 build-platform-docker
  make -C <dir> platform=arm64 build-platform-docker
  make -C <dir> manifest-platform-docker
  ```

## Image Build Pattern

- Image `Makefile`s set `IMAGE` and `VERSION`, then include `../make/docker.mk`.
- `make/docker.mk` intentionally builds from the repo root context with `docker build -f Dockerfile ... ..` so Dockerfiles can copy shared files such as `scripts/install-image-tool`.
- Keep `.dockerignore` current when adding large or sensitive top-level paths.
- Dockerfiles should call `install-image-tool`; put reusable download/checksum/extract logic in `scripts/install-image-tool.d/` and image-specific logic in that image's `scripts/install-image-tool.d/` directory.
- If changing hadolint suppressions, update the image directory's `.hadolint.yaml`; there is no top-level hadolint config.

## Release Image

- `release/` installs `gh`, `goreleaser`, and `uplift` through `release/scripts/install-image-tool.d/`.
- It also copies `release/deploy`, `release/package`, `release/version`, and `release/.uplift.yml`.
- `release/.uplift.yml` uses release commits shaped like `chore(release): $VERSION [ci skip]`.

## Gotchas

- `.gitmodules` uses the SSH URL `git@github.com:alexfalkowski/bin.git`.
- Push and manifest targets require DockerHub credentials.
- `scripts/compose` prefers `podman compose` over `docker compose`.
- `make clean` is destructive: it prunes all unused Docker or Podman images.
