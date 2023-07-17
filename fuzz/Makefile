testname  := FuzzMain
corpusdir := ./testdata/fuzz/${testname}

fuzz: 
	mkdir -p ${corpusdir}
	rm -rf ./go-fuzz-corpus
	git clone https://github.com/dvyukov/go-fuzz-corpus.git ./go-fuzz-corpus/
	file2fuzz -o ${corpusdir} ./go-fuzz-corpus/json/corpus/*  ./corpus/* 

run:
	GOARCH=amd64 go test -fuzz=${testname} -v

run-arm:
	GOARCH=arm64 go test -fuzz=FuzzAst -v

clean:
	rm -rf ./go-fuzz-corpus/
	rm -rf ./testdata/