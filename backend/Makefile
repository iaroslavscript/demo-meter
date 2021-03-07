
VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
GOFILES := $(shell find . -type f -name '*.go')

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X 'main.BuildVersion=$(VERSION)' -X 'main.BuildTime=$(shell date --utc --iso-8601=s)' -X 'main.BuildCommit=$(BUILD)'"

.PHONY: all build fmt test

all: help

help: Makefile
	@echo "Available commands:"
	@echo "\tbuild\t\tCombined call of fmt-check, test, go-build"
	@echo "\tgo-build\tCompile go files and bulid application"
	@echo "\tfmt-check\tRun style checking"
	@echo "\tfmt-fix\t\tRun autofix style errors"
	@echo "\ttest\t\tRun unit tests"

build: fmt-check test go-build

go-build:
	go build $(LDFLAGS)

fmt-check:
	[ -z "$$(gofmt -l $(GOFILES))" ] || (echo "Run 'make fmt-fix' to fix style errors first!" && exit 1)

fmt-fix:
	go fmt */**

test:
	go test -v */**
