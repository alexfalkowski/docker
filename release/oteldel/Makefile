.PHONY: dep build lint test

dep:
	go mod download

build:
	mkdir -p bin
	go build -o bin/oteldel ./cmd/oteldel

lint:
	test -z "$$(gofmt -l $$(find . -type f -name '*.go'))"
	go vet ./...

test:
	go test ./...
