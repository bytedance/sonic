testname  := FuzzMain
corpusdir := ./testdata/fuzz/${testname}

fuzz: 
	mkdir -p ${corpusdir}
	rm -rf ./go-fuzz-corpus
	git clone https://github.com/dvyukov/go-fuzz-corpus.git ./go-fuzz-corpus/
	file2fuzz -o ${corpusdir} ./go-fuzz-corpus/json/corpus/*  ./corpus/* 
	rm -rf ./go-fuzz-corpus/

run:
	GOMAXPROCS=4 GOARCH=amd64 go test -fuzz=${testname} -v -fuzztime 5m

clean:
	rm -rf ./go-fuzz-corpus/
	rm -rf ./testdata/