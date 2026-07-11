BINARY=bin/mco

.PHONY: fmt
fmt:
	gofmt -w .

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	mkdir -p bin
	go build -o $(BINARY) ./cmd/mco

.PHONY: generate
generate: build
	./$(BINARY) generate graph
	./$(BINARY) generate graph-html

.PHONY: validate
validate: build
	./$(BINARY) validate

.PHONY: graph
graph: build
	./$(BINARY) graph

.PHONY: doctor
doctor: build
	./$(BINARY) doctor

.PHONY: dev
dev: fmt test build generate
