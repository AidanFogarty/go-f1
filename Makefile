PROJECT = $(shell basename $(PWD))

GOPATH ?= $(shell go env GOPATH)
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

ORG = github.com/AidanFogarty/newbie

.DEFAULT_GOAL := help

## build: Build the binary.
build:
	go build -ldflags="-w -s"

## generate: Generate mock files for tests in the project.
generate:
	go generate ./...

## test: Run all the tests in the project
test:
	go test ./... -v

## lint: Run linter on all Go files in the project
lint:
	golangci-lint run --exclude-use-default=false

## fmt: Formats all Go code in the project
fmt:
	go fmt ./...

vet:
	go vet ./...

dupl:
	dupl -t 100 $(shell find . ! -name '*_test.go' -name '*.go')

## ci: Run linter, formatter and tests
ci: build fmt vet lint dupl test

# Thanks to: https://github.com/azer/go-makefile-example
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo