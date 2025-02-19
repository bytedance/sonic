#!/bin/bash

set -xe

source "$(dirname "$0")/../scripts/go_flags.sh"

compile_flag=$(get_go_linkname_flag || echo "")

cd ./issue_test
cp race_test_go race_test.go
go test "$compile_flag" -v -run=TestRaceEncode -race -count=100 . > test_race.log || true
if ! grep -q "WARNING: DATA RACE" ./test_race.log; then
    echo "TEST FAILED: should data race here"
    exit 1
fi
mv race_test.go race_test_go
rm -vrf test_race.log