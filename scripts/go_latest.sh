#!/bin/bash
TAG=$1

git clone -b release-branch.$TAG https://github.com/golang/go.git $TAG

cd $TAG/src

./all.bash



