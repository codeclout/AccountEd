SHELL=/bin/bash
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS = -uec

.PHONY: update-go-packages
update-go-packages:
	go get -u ./...
	go mod tidy

.PHONY: build-binary
build-binary:
	CGO_ENABLED=0 GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) go build -o accountEd -ldflags="-s" -v ./