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
	-docker buildx rm AccountEdSession
	docker buildx create --name AccountEdSession
	docker buildx use AccountEdSession
	docker buildx inspect AccountEdSession --bootstrap

	docker buildx build  \
		--build-arg ENVIRONMENT=$${ENVIRONMENT} \
		--build-arg AWS_ACCESS_KEY_ID=$${AWS_ACCESS_KEY_ID} \
		--build-arg AWS_PRE_REGISTRATION_PARAM=$${AWS_PRE_REGISTRATION_PARAM} \
		--build-arg AWS_REGION=$${AWS_REGION} \
		--build-arg AWS_ROLE_TO_ASSUME=$${AWS_ROLE_TO_ASSUME} \
		--build-arg AWS_SECRET_ACCESS_KEY=$${AWS_SECRET_ACCESS_KEY} \
		--build-arg SESSION_PORT=$${SESSION_PORT} \
		--build-arg SESSION_STORAGE_TABLE_NAME=$${SESSION_STORAGE_TABLE_NAME} \
		--platform linux/arm64 --load --target=prod -t sch00l.io/session-$${ENVIRONMENT} .

.PHONY: init-local-environment
init-local-environment: build-image
	docker compose -f ./docker-compose.yml -p accounted up -d $(shell if [ "$${SESSION_DEV}" = true ]; \
	then echo 'session-dev'; \
	else echo 'session'; \
	fi)

.PHONY: release-image
release-image: build-data
	docker tag $${ENVIRONMENT}-sch00l.io-session-$(shell go env GOARCH) \
	ghcr.io/$(GH_REPOSITORY):$(ENVIRONMENT)-$(shell echo $(IMAGE_TAG) | cut -c 1-12)

	docker push ghcr.io/$(GH_REPOSITORY):$(ENVIRONMENT)-$(shell echo $(IMAGE_TAG) | cut -c 1-12)
	docker logout ghcr.io/$${GH_ACTOR}

.PHONY: test-all
test-all:
	go test ./...

.PHONY: update-go-packages
update-go-packages:
	go get -u ./...
