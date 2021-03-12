
VERSION := $(shell git describe --tags HEAD 2>/dev/null || echo VAR_NOT_SET)

.PHONY: all help docker-build

all: help

help: Makefile
	@echo "Available commands:"
	@echo "\thelp\t\tShow this help"
	@echo "\tdocker-build\tBuild docker image"

docker-build:
	DOCKER_BUILDKIT=1 docker build --build-arg flovar=alpine -t apiserver:$(VERSION) .
