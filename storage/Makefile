SHELL=/bin/bash
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS = -uec

BUILD_EPOCH := $(shell date +"%s")
BUILD_TIME := $(shell date -u -r $(BUILD_EPOCH) +'%Y%m%d-%H%M%S')
GO_VERSION := 1.21.1

IMAGE_TAG ?= $(shell git rev-parse --short HEAD)

.PHONY: build-data
build-data:
	@echo build-epoch $(BUILD_EPOCH)
	@echo image-tag $(IMAGE_TAG)
	@echo build-time $(BUILD_TIME)

.PHONY: build-image
build-image:
	docker buildx install
	-docker buildx rm AccountEdStorage
	docker buildx create --name AccountEdStorage
	docker buildx use AccountEdStorage
	docker buildx inspect AccountEdStorage --bootstrap

	docker buildx build  \
		--build-arg ENVIRONMENT=$${ENVIRONMENT} \
		--build-arg AWS_REGION=$${AWS_REGION} \
		--build-arg AWS_ROLE_TO_ASSUME=$${AWS_ROLE_TO_ASSUME} \
		--build-arg STORAGE_PORT=$${STORAGE_PORT} \
		--platform linux/arm64 --load --target=prod -t sch00l.io/storage-$${ENVIRONMENT} .

.PHONY: init-local-environment
init-local-environment: build-image
	docker compose -f ./docker-compose.yml -p accounted up -d $(shell if [ "$${STORAGE_DEV}" = true ]; \
	then echo 'storage-dev'; \
	else echo 'storage'; \
	fi)

.PHONY: release-image
release-image: build-data
	docker tag sch00l.io/storage-$${ENVIRONMENT} \
	ghcr.io/$(GH_REPOSITORY):$(ENVIRONMENT)-$(shell echo $(IMAGE_TAG) | cut -c 1-12)

	docker push ghcr.io/$(GH_REPOSITORY):$(ENVIRONMENT)-$(shell echo $(IMAGE_TAG) | cut -c 1-12)
	docker logout ghcr.io/$${GH_ACTOR}

.PHONY: test-all
test-all:
	go test ./...

.PHONY: update-go-packages
update-go-packages:
	go get -u ./...
