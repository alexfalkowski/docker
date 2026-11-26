# Accepted design for images

These entries are part of the [accepted-design index](../accepted-design.md).

## Base image tags

The Dockerfiles intentionally use mutable tags for base images. Raise
tag-based `FROM` lines as a review finding only when the task calls for
digest-pinned builds.

[Root Dockerfile](../../../root/Dockerfile) ·
[Go Dockerfile](../../../go/Dockerfile)

## Runtime privileges

The images intentionally allow root-level operations where needed, including
for tools such as `mkcert`. The Docker tooling image runs as root, and the
`circleci` user in the other images has passwordless `sudo` through the root
base image. Raise this privilege model as a review finding only when the task
specifically calls for hardening runtime privileges.

[Docker tooling Dockerfile](../../../docker/Dockerfile) ·
[Root Dockerfile](../../../root/Dockerfile) ·
[Go Dockerfile](../../../go/Dockerfile) ·
[Ruby Dockerfile](../../../ruby/Dockerfile)
