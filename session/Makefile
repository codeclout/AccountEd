SHELL=/bin/bash
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS = -uec

BUILD_EPOCH := $(shell date +"%s")
BUILD_TIME := $(shell date -u -r $(BUILD_EPOCH) +'%Y%m%d-%H%M%S')
GO_VERSION := 1.20.5

ifndef IMAGE_TAG
	IMAGE_TAG := $(shell git rev-parse --short HEAD)
endif

ifeq ($(strip ${IMAGE_TAG}),)
	IMAGE_TAG := $(shell git rev-parse --short HEAD)
endif

ifndef AWS_PRE_REGISTRATION_HASH_PARAM
$(error AWS_PRE_REGISTRATION_HASH_PARAM is not set)
endif

ifndef AWS_PRE_REGISTRATION_ROLE
$(error AWS_PRE_REGISTRATION_ROLE is not set)
endif

ifndef AWS_REGION
$(error AWS_REGION is not set)
endif

ifndef NOTIFICATION_SERVER_HOST
$(error NOTIFICATION_SERVER_HOST is not set)
endif

ifndef NOTIFICATION_SERVER_PORT
$(error NOTIFICATION_SERVER_PORT is not set)
endif

ifndef SESSION_SERVER_HOST
$(error SESSION_SERVER_HOST is not set)
endif

ifndef SESSION_SERVER_PORT
$(error SESSION_SERVER_PORT is not set)
endif

ifndef DOMAIN
$(error DOMAIN is not set)
endif

.PHONY: ci-buildx-register-image
ci-buildx-register-image: build-data test-all
	docker buildx install
	docker buildx create --name memberServiceBuilder
	docker buildx use memberServiceBuilder
	docker buildx inspect memberServiceBuilder --bootstrap

	docker buildx build \
	--build-arg AWS_PRE_REGISTRATION_HASH_PARAM=$${AWS_PRE_REGISTRATION_HASH_PARAM} \
	--build-arg AWS_PRE_REGISTRATION_ROLE=$${AWS_PRE_REGISTRATION_ROLE} \
	--build-arg AWS_REGION=$${AWS_REGION} \
	--build-arg NOTIFICATION_SERVER_HOST=$${NOTIFICATION_SERVER_HOST} \
	--build-arg NOTIFICATION_SERVER_PORT=$${NOTIFICATION_SERVER_PORT} \
	--build-arg SESSION_SERVER_HOST=$${SESSION_SERVER_HOST} \
	--build-arg SESSION_SERVER_PORT=$${SESSION_SERVER_PORT} \
	--build-arg DOMAIN=$${DOMAIN} \

	--platform=linux/amd64,linux/arm64 \
	--progress=plain \
	--target=prod \
	-t ghcr.io/$${GH_REPOSITORY}:$(shell echo $${ENV})-$(shell echo $${IMAGE_TAG} | cut -c 1-12) --push .

	docker logout ghcr.io/$${GH_ACTOR}

.PHONY: test-all
test-all:
	go test ./...

.PHONY: build-data
build-data:
	@echo build-epoch $(BUILD_EPOCH)
	@echo image-tag $(IMAGE_TAG)
	@echo build-time $(BUILD_TIME)
