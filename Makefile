all: bin/main

bin/main: src/github.com/maciej-gol/word-suggest/cmd/main/main.go src/github.com/maciej-gol/word-suggest/internal/query/query.go
	GOPATH=`pwd` go get github.com/maciej-gol/word-suggest/cmd/main
	GOPATH=`pwd` go build -o bin/main src/github.com/maciej-gol/word-suggest/cmd/main/main.go

clean:
	rm -rf bin
