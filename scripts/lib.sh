#!/bin/bash

changed() {
    git diff --quiet HEAD "$(git describe --tags --abbrev=0 HEAD)" -- "$1" || echo changed
}

perform() {
    dirs=$(find . -mindepth 1 -maxdepth 1 -type d -not -name scripts -not -path '*/\.*' | xargs -n 1 basename)

    for dir in $dirs
    do
        res=$(changed "$dir")

        if [ "$res" == "changed" ]; then
            echo "$1 $dir"
            make -C "$dir" "$1"
        fi
    done
}
