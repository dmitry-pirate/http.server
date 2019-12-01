.PHONY: build
build:
	go build -v ./cmd/server

.PHONY: test
test:
	go test -v -race -timeout 30s ./tests

.DEFAULT_GOAL := build