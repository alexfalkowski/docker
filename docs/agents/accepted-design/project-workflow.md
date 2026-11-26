# Accepted design for the project workflow

These entries are part of the [accepted-design index](../accepted-design.md).

## Image versions and root releases

The maintainer owns image `VERSION` values. Bump them only when explicitly
asked. The maintainer releases root image changes in two stages.

1. Update only `root/` and bump `VERSION` in `root/Makefile`. Use a minor bump
   for ordinary changes and a major bump for major changes to the image
   contract, including major upgrades to dependencies shipped in the image.
2. After the new root image is published, update the Dockerfiles that use
   `alexfalkowski/root` and bump their image versions in a separate change.
   Use a major bump for dependent images if root received a major bump and
   a minor bump otherwise.

[Root Makefile](../../../root/Makefile) ·
[Maintainer notes](../../../README.md#-maintainer-notes)

## Repository scanning

The root `trivy-repo` target calls the shared `bin/build/sec/trivy-repo` helper
directly. Preserve this wiring when changing the root Makefile. Keep the
`go.mak`, `ruby.mak`, and `_service.mak` fragments out of the root Makefile.

[Root Makefile](../../../Makefile) ·
[CircleCI workflow](../../../.circleci/continue_config.yml)
