testname  := FuzzMain
corpusdir := ./testdata/fuzz/${testname}
go := go

fuzz:
	${go} version
	mkdir -p ${corpusdir}
	rm -rf ./go-fuzz-corpus
	git clone https://github.com/dvyukov/go-fuzz-corpus.git ./go-fuzz-corpus/
	${go} install golang.org/x/tools/cmd/file2fuzz@latest
	file2fuzz -o ${corpusdir} ./go-fuzz-corpus/json/corpus/*  ./corpus/* 

run:
	${go} version
	SONIC_FUZZ_MEM_LIMIT=2 GOMAXPROCS=4 ${go} test -fuzz=${testname} -v -fuzztime 10m
	
clean:
	rm -rf ./go-fuzz-corpus/
	rm -rf ./testdata/