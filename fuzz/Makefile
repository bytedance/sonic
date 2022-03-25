
fuzz: 
	mkdir -p ./testdata/fuzz/FuzzMain
	git clone https://github.com/dvyukov/go-fuzz-corpus.git ./go-fuzz-corpus/
	file2fuzz -o ./testdata/fuzz/FuzzMain ./go-fuzz-corpus/json/corpus/*  ./corpus/* 
	rm -rf ./go-fuzz-corpus/
	go test -fuzz=Fuzz -v

clean:
	rm -rf ./go-fuzz-corpus/
	rm -rf ./testdata/