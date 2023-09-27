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
	-docker buildx rm AccountEdMembers
	docker buildx create --name AccountEdMembers
	docker buildx use AccountEdMembers
	docker buildx inspect AccountEdMembers --bootstrap

	docker buildx build  \
		--build-arg ENVIRONMENT=$${ENVIRONMENT} \
		--build-arg AWS_PRE_REGISTRATION_HASH_PARAM=$${AWS_PRE_REGISTRATION_HASH_PARAM} \
		--build-arg AWS_PRE_REGISTRATION_ROLE=$${AWS_PRE_REGISTRATION_ROLE} \
		--build-arg AWS_REGION=$${AWS_REGION} \
		--build-arg DOMAIN=$${DOMAIN} \
		--build-arg MEMBER_PORT=$${MEMBER_PORT} \
		--platform linux/arm64 --load --target=prod -t sch00l.io/members-$${ENVIRONMENT} -f ./Dockerfile .

.PHONY: init-local-environment
init-local-environment: build-image
	docker compose -f ./docker-compose.yml -p accounted up -d $(shell if [ "$${MEMBER_DEV}" = true ]; \
	then echo 'member-dev'; \
	else echo 'member'; \
	fi)

.PHONY: release-image
release-image: build-data
	docker tag sch00l.io/members-$${ENVIRONMENT} \
	ghcr.io/$(GH_REPOSITORY):$(ENV)-$(shell echo $(IMAGE_TAG) | cut -c 1-12)

	docker push ghcr.io/$(GH_REPOSITORY):$(ENV)-$(shell echo $(IMAGE_TAG) | cut -c 1-12)
	docker logout ghcr.io/$${GH_ACTOR}

.PHONY: test-all
test-all:
	go test ./...

.PHONY: update-go-packages
update-go-packages:
	go get -u ./...
