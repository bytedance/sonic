#!/bin/bash

set -xe
cd ./issue_test
mv race_test_go race_test.go
go test -v -run=TestRaceEncode -race -count=100 . > test_race.log || true
if ! grep -q "WARNING: DATA RACE" ./test_race.log; then
    echo "TEST FAILED: should data race here"
    exit 1
fi
mv race_test.go race_test_go
rm -vrf test_race.log