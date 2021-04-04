#!/bin/bash
set -euo pipefail

# Tag the new release
standard-version --releaseCommitMessageFormat "chore(release): {{currentTag}} [ci skip]"
git push --follow-tags origin master

# Create a GH release
echo "${GITHUB_TOKEN}" | gh auth login --with-token
latest_tag=$(git tag | sort -V | tail -1)
gh release create "$latest_tag" -t "$latest_tag"
