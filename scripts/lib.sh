#!/usr/bin/env bash

perform() {
    # shellcheck disable=SC2038
    dirs=$(find . -mindepth 1 -maxdepth 1 -type d -not -name scripts -not -name bin -not -name grafana -not -name prometheus -not -name vault -not -name status -not -path '*/\.*' | xargs -n 1 basename)

    for dir in $dirs
    do
      make -C "$dir" "$1"
    done
}
