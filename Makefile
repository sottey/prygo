SHELL := /bin/bash

GO ?= go
GOLANGCI_LINT ?= golangci-lint

.NAME_PROTECTION: all ci build test lint fmt clean

all: build

ci: clean fmt lint test build

build:
	$(GO) build ./...

test:
	$(GO) test -count=1 ./...

lint:
	$(GOLANGCI_LINT) run ./...

fmt:
	$(GO) fmt ./...

clean:
	rm -rf prygo tmp *.out
