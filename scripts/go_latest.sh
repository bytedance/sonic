#!/bin/bash

BRANCH=$1

git clone https://github.com/golang/go.git go_$BRANCH

cd go_$BRANCH/src
git checkout -b build/$BRANCH origin/$BRANCH 

./all.bash

cd ..
export GOROOT=$PWD
export PATH=$GOROOT/bin:$PATH
go version

cd ..
GOMAXPROCS=4 go test -v -race ./...