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
	-docker buildx rm AccountEdNotifications
	docker buildx create --name AccountEdNotifications
	docker buildx use AccountEdNotifications
	docker buildx inspect AccountEdNotifications --bootstrap

	docker buildx build  \
		--build-arg ENV=$${ENV} \
		--build-arg AWS_REGION=$${AWS_REGION} \
		--build-arg EMAIL_PROCESSOR_API_KEY=$${EMAIL_PROCESSOR_API_KEY} \
		--build-arg EMAIL_PROCESSOR_DOMAIN=$${EMAIL_PROCESSOR_DOMAIN} \
		--build-arg EMAIL_VERIFIER_API_PATH=$${EMAIL_VERIFIER_API_PATH} \
		--build-arg NOTIFICATIONS_PORT=$${NOTIFICATIONS_PORT} \
		--build-arg PERFORMANCE_SLA=$${PERFORMANCE_SLA} \
		--platform linux/arm64 --load --target=prod -t sch00l.io/notifications-$${ENV} .

.PHONY: init-local-environment
init-local-environment: build-image
	docker compose -f ./docker-compose.yml -p accounted up -d $(shell if [ "$${NOTIFICATIONS_DEV}" = true ]; \
	then echo 'notifications-dev'; \
	else echo 'notifications'; \
	fi)

.PHONY: release-image
release-image: build-data
	docker tag $${ENV}-sch00l.io-notifications-$(shell go env GOARCH) \
	ghcr.io/$(GH_REPOSITORY):$(ENV)-$(shell echo $(IMAGE_TAG) | cut -c 1-12)

	docker push ghcr.io/$(GH_REPOSITORY):$(ENV)-$(shell echo $(IMAGE_TAG) | cut -c 1-12)
	docker logout ghcr.io/$${GH_ACTOR}

.PHONY: test-all
test-all:
	go test ./...

.PHONY: update-go-packages
update-go-packages:
	go get -u ./...
