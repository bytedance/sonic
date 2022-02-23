#!/usr/bin/env bash

pwd=$(pwd)
export SONIC_NO_ASYNC_GC=1

# cd $pwd/encoder
# go test -benchmem -run=^$ -benchtime=100000x -bench "^(BenchmarkEncoder_Generic_Sonic|BenchmarkEncoder_Generic_Sonic_Fast|BenchmarkEncoder_Generic_JsonIter|BenchmarkEncoder_Generic_GoJson|BenchmarkEncoder_Generic_StdLib|BenchmarkEncoder_Binding_Sonic|BenchmarkEncoder_Binding_Sonic_Fast|BenchmarkEncoder_Binding_JsonIter|BenchmarkEncoder_Binding_GoJson|BenchmarkEncoder_Binding_StdLib|BenchmarkEncoder_Parallel_Generic_Sonic|BenchmarkEncoder_Parallel_Generic_Sonic_Fast|BenchmarkEncoder_Parallel_Generic_JsonIter|BenchmarkEncoder_Parallel_Generic_GoJson|BenchmarkEncoder_Parallel_Generic_StdLib|BenchmarkEncoder_Parallel_Binding_Sonic|BenchmarkEncoder_Parallel_Binding_Sonic_Fast|BenchmarkEncoder_Parallel_Binding_JsonIter|BenchmarkEncoder_Parallel_Binding_GoJson|BenchmarkEncoder_Parallel_Binding_StdLib)$"

cd $pwd/decoder
go test -benchmem -run=^$ -benchtime=100000x -bench "^(BenchmarkDecoder_Generic_Sonic|BenchmarkDecoder_Generic_StdLib|BenchmarkDecoder_Generic_JsonIter|BenchmarkDecoder_Generic_GoJson|BenchmarkDecoder_Binding_Sonic|BenchmarkDecoder_Binding_StdLib|BenchmarkDecoder_Binding_JsonIter|BenchmarkDecoder_Binding_GoJson|BenchmarkDecoder_Parallel_Generic_Sonic|BenchmarkDecoder_Parallel_Generic_StdLib|BenchmarkDecoder_Parallel_Generic_JsonIter|BenchmarkDecoder_Parallel_Generic_GoJson|BenchmarkDecoder_Parallel_Binding_Sonic|BenchmarkDecoder_Parallel_Binding_StdLib|BenchmarkDecoder_Parallel_Binding_JsonIter|BenchmarkDecoder_Parallel_Binding_GoJson)$"

# cd $pwd/ast
# go test -benchmem -run=^$ -benchtime=100000x -bench "^(BenchmarkGetOne_Sonic|BenchmarkGetOne_Gjson|BenchmarkGetOne_Jsoniter|BenchmarkGetOne_Parallel_Sonic|BenchmarkGetOne_Parallel_Gjson|BenchmarkGetOne_Parallel_Jsoniter|BenchmarkSetOne_Sonic|BenchmarkSetOne_Sjson|BenchmarkSetOne_Jsoniter|BenchmarkSetOne_Parallel_Sonic|BenchmarkSetOne_Parallel_Sjson|BenchmarkSetOne_Parallel_Jsoniter)$"

# go test -benchmem -run=^$ -benchtime=10000x -bench "^(BenchmarkParser_Sonic|BenchmarkParser_Gjson|BenchmarkParser_JsonIter|BenchmarkParser_Parallel_Sonic|BenchmarkParser_Parallel_Gjson|BenchmarkParser_Parallel_StdLib|BenchmarkParser_Parallel_JsonIter|BenchmarkParseOne_Sonic|BenchmarkParseOne_Gjson|BenchmarkParseOne_Jsoniter|BenchmarkParseOne_Parallel_Sonic|BenchmarkParseOne_Parallel_Gjson|BenchmarkParseOne_Parallel_Jsoniter|BenchmarkParseSeven_Sonic|BenchmarkParseSeven_Gjson|BenchmarkParseSeven_Jsoniter|BenchmarkParseSeven_Parallel_Sonic|BenchmarkParseSeven_Parallel_Gjson|BenchmarkParseSeven_Parallel_Jsoniter)$"

# go test -benchmem -run=^$ -benchtime=100000x -bench '^(BenchmarkEncodeRaw|BenchmarkEncodeSkip|BenchmarkEncodeLoad)$'

# go test -benchmem -run=^$ -benchtime=10000000x -bench "^(BenchmarkNodeGetByPath|BenchmarkStructGetByPath|BenchmarkNodeIndex|BenchmarkStructIndex|BenchmarkSliceIndex|BenchmarkMapIndex|BenchmarkNodeGet|BenchmarkSliceGet|BenchmarkMapGet|BenchmarkNodeSet|BenchmarkMapSet|BenchmarkNodeSetByIndex|BenchmarkSliceSetByIndex|BenchmarkStructSetByIndex|BenchmarkNodeUnset|BenchmarkMapUnset|BenchmarkNodUnsetByIndex|BenchmarkSliceUnsetByIndex|BenchmarkNodeAdd|BenchmarkSliceAdd|BenchmarkMapAdd)$"

unset SONIC_NO_ASYNC_GC
cd $pwd