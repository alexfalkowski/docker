#!/bin/bash
set -euo pipefail

standard-version --releaseCommitMessageFormat "chore(release): {{currentTag}} [ci skip]"
git push --follow-tags origin master
