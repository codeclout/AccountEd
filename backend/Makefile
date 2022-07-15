SHELL=/bin/bash
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS = -uec

.PHONY: update-go-packages
update-go-packages:
	go get -u ./...
	go mod tidy