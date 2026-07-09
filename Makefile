BINARY=bin/mco

build:
	mkdir -p bin
	go build -o $(BINARY) ./cmd/mco

test:
	go test ./...
