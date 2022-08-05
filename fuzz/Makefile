
fuzz: 
	mkdir -p ./testdata/fuzz/FuzzMain
	rm -rf ./go-fuzz-corpus
	git clone https://github.com/dvyukov/go-fuzz-corpus.git ./go-fuzz-corpus/
	file2fuzz -o ./testdata/fuzz/FuzzMain ./go-fuzz-corpus/json/corpus/*  ./corpus/* 

run:
	go test -fuzz=Fuzz -v

clean:
	rm -rf ./go-fuzz-corpus/
	rm -rf ./testdata/