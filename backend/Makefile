
VERSION := $(shell git describe --tags HEAD 2>/dev/null || echo VAR_NOT_SET)
BUILD := $(shell git rev-parse --short HEAD 2>/dev/null || echo VAR_NOT_SET)
GOFILES := $(shell find . -type f -name '*.go')
DATETIME := $(shell date --utc -Iseconds)
OUTPUT ?= "./app"

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X 'main.BuildVersion=$(VERSION)' -X 'main.BuildTime=$(DATETIME)' -X 'main.BuildCommit=$(BUILD)'"

.PHONY: all build fmt test

all: help

help: Makefile
	@echo "Available commands:"
	@echo "\tbuild\t\tCombined call of fmt-check, test, go-build"
	@echo "\tgo-build\tCompile go files and bulid application"
	@echo "\tgo-deps\t\tDownload dependencies without building the project"
	@echo "\tfmt-check\tRun style checking"
	@echo "\tfmt-fix\t\tRun autofix style errors"
	@echo "\ttest\t\tRun unit tests"
	@echo "\tdocker-build\tBuild docker image"

build: fmt-check test go-build

go-deps:
	go get -d

go-build:
	go build $(LDFLAGS) -o $(OUTPUT)

fmt-check:
	[ -z "$$(gofmt -l $(GOFILES))" ] || (echo "Run 'make fmt-fix' to fix style errors first!" && exit 1)

fmt-fix:
	go fmt */**

test:
	go test -v */**

docker-build:
	DOCKER_BUILDKIT=1 docker build --build-arg alpine_flovar=alpine3.13 -t backend:$(VERSION) .
