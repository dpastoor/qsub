SHELL := /bin/bash

export GOVER ?= "1.16"
export BUILD ?= "0"
export COMMIT ?= `git rev-parse --short HEAD`
export DATE ?= `date +%s`
export VERSION ?= "dev"

all: clean release 

clean:
	rm -f qsub  

release:
	rm -rf dist
	goreleaser --rm-dist

test:
	go test -v ./...
	
tidy:
	go mod tidy
	