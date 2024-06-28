testname  := FuzzMain
corpusdir := ./testdata/fuzz/${testname}

fuzz: 
	mkdir -p ${corpusdir}
	rm -rf ./go-fuzz-corpus
	git clone https://github.com/dvyukov/go-fuzz-corpus.git ./go-fuzz-corpus/
	go install golang.org/x/tools/cmd/file2fuzz@latest
	file2fuzz -o ${corpusdir} ./go-fuzz-corpus/json/corpus/*  ./corpus/* 

run:
	SONIC_FUZZ_MEM_LIMIT=2 GOMAXPROCS=2 go test -fuzz=${testname} -v -fuzztime 5m

runopt:
	SONIC_FUZZ_MEM_LIMIT=2 SONIC_USE_OPTDEC=1 SONIC_USE_FASTMAP=1 SONIC_ENCODER_USE_VM=1  GOMAXPROCS=2 go test -fuzz=${testname} -v -fuzztime 5m

clean:
	rm -rf ./go-fuzz-corpus/
	rm -rf ./testdata/