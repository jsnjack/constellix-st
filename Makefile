BINARY:=constellix-st
PWD:=$(shell pwd)
VERSION=0.0.0
MONOVA:=$(shell which monova dot 2> /dev/null)

version:
ifdef MONOVA
override VERSION=$(shell monova)
else
	$(info "Install monova (https://github.com/jsnjack/monova) to calculate version")
endif

.ONESHELL:
bin/${BINARY}: version *.go
	go build -ldflags="-X main.Version=${VERSION}" -o bin/${BINARY}

build: bin/${BINARY}

release: build
	release_on_github -f bin/${BINARY} -r jsnjack/constellix-st -t "v`monova`"

.PHONY: version release build
