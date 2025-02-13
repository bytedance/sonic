#!/bin/bash

set -xe

get_go_linkname_flag() {
    if ! command -v go &> /dev/null; then
        return
    fi

    local go_version=$(go version | awk '{print $3}' | sed -E 's/go([0-9]+\.[0-9]+(\.[0-9]+)?).*/\1/')
    IFS='.' read -r major minor _ <<< "$go_version"
    
    if ! [[ "$major" =~ ^[0-9]+$ ]] || ! [[ "$minor" =~ ^[0-9]+$ ]]; then
        return
    fi

    if (( major > 1 || (major == 1 && minor >= 23) )); then
        echo "-ldflags=-checklinkname=0"
    fi
}

cd ./issue_test
cp race_test_go race_test.go
go test "$(get_go_linkname_flag)" -v -run=TestRaceEncode -race -count=100 . > test_race.log || true
if ! grep -q "WARNING: DATA RACE" ./test_race.log; then
    echo "TEST FAILED: should data race here"
    exit 1
fi
rm -vrf race_test.go test_race.log