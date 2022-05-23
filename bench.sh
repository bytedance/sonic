#!/usr/bin/env bash

pwd=$(pwd)
export SONIC_NO_ASYNC_GC=1

cd $pwd/encoder
go test -benchmem -run=^$ -benchtime=100000x -bench "^(BenchmarkEncoder_.*)$"

cd $pwd/decoder
go test -benchmem -run=^$ -benchtime=100000x -bench "^(BenchmarkDecoder_.*)$"

cd $pwd/ast
go test -benchmem -run=^$ -benchtime=100000x -bench "^(BenchmarkGet.*|BenchmarkEncode.*)$"

go test -benchmem -run=^$ -benchtime=10000x -bench "^(BenchmarkParser_.*)$"

go test -benchmem -run=^$ -benchtime=10000000x -bench "^(BenchmarkNodeGetByPath|BenchmarkStructGetByPath|BenchmarkNodeIndex|BenchmarkStructIndex|BenchmarkSliceIndex|BenchmarkMapIndex|BenchmarkNodeGet|BenchmarkSliceGet|BenchmarkMapGet|BenchmarkNodeSet|BenchmarkMapSet|BenchmarkNodeSetByIndex|BenchmarkSliceSetByIndex|BenchmarkStructSetByIndex|BenchmarkNodeUnset|BenchmarkMapUnset|BenchmarkNodUnsetByIndex|BenchmarkSliceUnsetByIndex|BenchmarkNodeAdd|BenchmarkSliceAdd|BenchmarkMapAdd)$"

unset SONIC_NO_ASYNC_GC
cd $pwd