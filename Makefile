SHELL=/bin/bash
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS = -uec

GO_VERSION := 1.18.3

GOOS = $(shell go env GOOS)

export TF_INPUT = $(shell go env TF_INPUT)
export TF_LOG = $(shell go env TF_LOG)

export TF_VAR_LMS_ACCOUNT_ROLE_NAME = $(shell go env LMS_ACCOUNT_ROLE)
export TF_VAR_LMS_USER_ACCOUNT_EMAIL = $(shell go env LMS_USER_ACCOUNT_EMAIL)
export TF_VAR_PROXY_ACCOUNT_USERS_EMAIL = $(shell go env PROXY_ACCOUNT_USERS_EMAIL)
export TF_VAR_PROXY_ACCOUNT_ROLE_NAME = $(shell go env PROXY_ACCOUNT_ROLE_NAME)

.PHONY: init-local-environment
init-local-environment: build-docker
	cd ./migrations/mongo/local/migration && npm i
	docker compose config
	docker compose up

.PHONY: build-docker
build-docker:
	$(shell docker build --progress plain --target=prod -t accounted-$(shell go env GOARCH)-$${ENV} .)

.PHONY: update-go-packages
update-go-packages:
	${MAKE} -C backend update-go-packages

.PHONY: ci-buildx-register-container
ci-buildx-register-container:
	docker buildx install
	docker buildx create --name accountEdBuilder
	docker buildx use accountEdBuilder
	docker buildx inspect accountEdBuilder --bootstrap

	TAG := $(shell echo $${IMAGE_TAG} | cut -c 1-12)
	
	docker buildx build --build-arg ENV=$${ENV} --target=prod --platform linux/amd64,linux/arm64 -t ghcr.io/$${GH_ACTOR}/$${ECR_REPOSITORY}:${TAG} . --push
	docker buildx build --build-arg ENV=$${ENV} --target=prod --platform linux/amd64,linux/arm64 -t $${ECR_REGISTRY}/$${ECR_REPOSITORY}:${TAG} --push .

# .PHONY: setup-local-go
# setup-local-go:
# 	install-go
# 	init-go

# .PHONY: install-local-deps
# install-local-deps:
# 	go mod tidy
# 	go mod verify

# # compile the main application into a binary named accountEd
# .PHONY: init-local-build
# init-local-build:
# 	GOOS=$${GO_OS} GOARCH=$${GO_ARCH} go build -v -o accountEd ./backend

# register-image:
# 	docker push $${CONTAINER_REGISTRY}-$${ENV}-$${VERSION}
