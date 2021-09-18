#!/bin/bash

changed() {
    current_branch=$(git symbolic-ref -q --short HEAD)

    if [ "$current_branch" == "master" ]; then
        diff=$(git describe --tags --abbrev=0 HEAD)
    else
        diff="master"
    fi

    git diff --quiet HEAD "$diff" -- "$1" || echo changed
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
