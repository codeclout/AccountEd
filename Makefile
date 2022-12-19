SHELL=/bin/bash
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS = -uec

BUILD_EPOCH := $(shell date +"%s")
BUILD_TIME := $(shell date -u -r $(BUILD_EPOCH) +'%Y%m%d-%H%M%S')
GO_VERSION := 1.19.4

ifndef IMAGE_TAG
	IMAGE_TAG := $(shell git rev-parse --short HEAD)
endif

ifeq ($(strip ${IMAGE_TAG}),)
	IMAGE_TAG := $(shell git rev-parse --short HEAD)
endif

.PHONY: init-local-environment
init-local-environment: build-image
	cd ./migrations/mongo && npm i
	docker compose config
	docker compose up

.PHONY: build-image
build-image:
	$(shell docker buildx build --load --target=prod -t accounted-$(shell go env GOARCH)-$(shell echo $${ENV}) .)

.PHONY: update-go-packages
update-go-packages:
	${MAKE} -C backend update-go-packages

.PHONY: ci-buildx-register-image
ci-buildx-register-image: build-data
	docker buildx install
	docker buildx create --name accountEdBuilder
	docker buildx use accountEdBuilder
	docker buildx inspect accountEdBuilder --bootstrap

	docker buildx build --build-arg ENV=$${ENV} \
	--platform=linux/amd64,linux/arm64 \
	--progress=plain \
	--target=prod \
	-t ghcr.io/$${GH_REPOSITORY}:$(shell echo $${{ENV}})-$(shell echo $${IMAGE_TAG} | cut -c 1-12) --push .

	docker logout ghcr.io/$${GH_ACTOR}

.PHONY: build-data
build-data:
	@echo build-epoch $(BUILD_EPOCH)
	@echo image-tag $(IMAGE_TAG)
	@echo build-time $(BUILD_TIME)
