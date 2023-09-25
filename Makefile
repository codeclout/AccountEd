SHELL=/bin/bash
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS = -uec

BUILD_EPOCH := $(shell date +"%s")
BUILD_TIME := $(shell date -u -r $(BUILD_EPOCH) +'%Y%m%d-%H%M%S')
GO_VERSION := 1.21.1
SUBDIRS := $(wildcard */.)

.PHONY: init-local-environment
init-local-environment: build-image
	cd ./migrations/mongo && npm i
	docker compose config
	docker compose up

.PHONY: build-image
build-image:
	$(shell docker buildx build --load --target=prod -t accounted-$(shell go env GOARCH)-$(shell echo $${ENV}) .)

.PHONY: subdirs $(SUBDIRS))
subdirs: $(SUBDIRS)

$(SUBDIRS):
	${MAKE} -C $@ update-go-packages

.PHONY: update-go-packages
update-go-packages: subdirs