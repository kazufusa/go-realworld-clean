.PHONY: all

export GOOS=linux
export GOARCH=amd64

BINARY=go-realworld-clean
VERSION=$(shell git describe --abbrev=0 --tags 2>/dev/null || echo v0.1.0)
BUILD=$(shell git rev-parse HEAD 2>/dev/null || echo undefined)
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

all:
	go build -o ${BINARY} ${LDFLAGS}
	./${BINARY}

prepare:
	git config --local core.hooksPath .githooks
