
build: 
	mkdir -p ./testdata/fuzz/FuzzMain
	git clone https://github.com/dvyukov/go-fuzz-corpus.git ./go-fuzz-corpus/
	go install golang.org/x/tools/cmd/file2fuzz@latest
	file2fuzz -o ./testdata/fuzz/FuzzMain ./go-fuzz-corpus/json/corpus/*  ./corpus/* 

run:
	go test -fuzz=Fuzz -v

clean:
	rm -rf ./go-fuzz-corpus/
	rm -rf ./testdata/