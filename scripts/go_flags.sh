#!/bin/bash

set -xe

if [ -n "$_GO_FLAGS_LOADED" ]; then
    return
fi
_GO_FLAGS_LOADED=1

get_go_linkname_flag() {
    if ! command -v go &> /dev/null; then
        return 1
    fi

    local go_version
    go_version=$(go version | awk '{print $3}' | sed -E 's/go([0-9]+\.[0-9]+(\.[0-9]+)?).*/\1/')
    IFS='.' read -r major minor _ <<< "$go_version"
    
    if ! [[ "$major" =~ ^[0-9]+$ ]] || ! [[ "$minor" =~ ^[0-9]+$ ]]; then
        return 1
    fi

    if (( major > 1 || (major == 1 && minor >= 23) )); then
        echo "-ldflags=-checklinkname=0"
    fi
}