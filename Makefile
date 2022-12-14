SHELL=/bin/bash
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS = -uec

GO_VERSION := 1.19.4

.PHONY: init-local-environment
init-local-environment: build-image
	cd ./migrations/mongo && npm i
	docker compose config
	docker compose up

.PHONY: build-image
build-image:
	$(shell docker buildx build --load --target=prod -t accounted-$(shell go env GOARCH)-$(shell echo $${ENV}) .)

.PHONY: build-binary
build-binary:
	CGO_ENABLED=0 GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) go build -o accountEd -ldflags="-s" -v backend/

.PHONY: update-go-packages
update-go-packages:
	${MAKE} -C backend update-go-packages

.PHONY: ci-buildx-register-image
ci-buildx-register-image:
	docker buildx install
	docker buildx create --name accountEdBuilder
	docker buildx use accountEdBuilder
	docker buildx inspect accountEdBuilder --bootstrap
	
	docker buildx build --build-arg ENV=$${ENV} --target=prod --platform=linux/amd64,linux/arm64 -t ghcr.io/$${GH_REPOSITORY}:$(shell echo $${IMAGE_TAG} | cut -c 1-12) --push .
	#docker buildx build --build-arg ENV=$${ENV} --target=prod --platform linux/amd64,linux/arm64 -t $${ECR_REGISTRY}/$${ECR_REPOSITORY}:$(shell echo $${IMAGE_TAG} | cut -c 1-12) --push .

	docker logout ghcr.io/$${GH_ACTOR}
	#docker logout $${ECR_REGISTRY}


