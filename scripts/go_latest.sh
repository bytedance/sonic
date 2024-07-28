#!/bin/bash
TAG=$1

git clone https://github.com/golang/go.git $TAG

cd $TAG/src
git checkout -b build/$TAG $TAG 

./all.bash > /dev/null


