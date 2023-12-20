VERSION ?= $(shell git describe --tags --always --dirty="-dev")


.PHONY: lint
lint:
	go vet ./...
	golangci-lint run --timeout=5m
