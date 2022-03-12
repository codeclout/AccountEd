GO_VERSION := 1.17.7

export GOOS = $(shell go env GOOS)
export GOARCH = $(shell go env GOARCH)

export TF_INPUT = $(shell go env TF_INPUT)
export TF_LOG = $(shell go env TF_LOG)

export TF_VAR_LMS_ACCOUNT_ROLE = $(shell go env LMS_ACCOUNT_ROLE)
export TF_VAR_LMS_USER_ACCOUNT_EMAIL = $(shell go env LMS_USER_ACCOUNT_EMAIL)
export TF_VAR_PROXY_ACCOUNT_USERS_EMAIL = $(shell go env PROXY_ACCOUNT_USERS_EMAIL)
export TF_VAR_PROXY_ACCOUNT_ROLE_NAME = $(shell go env PROXY_ACCOUNT_ROLE_NAME)

.PHONY: build-docker
build-docker:
	docker build --target=prod -t accountEd-$${GO_ARCH}-$${ENV} .

.PHONY: ci-buildx-register-container
ci-buildx-register-container:
	docker buildx install
	docker buildx create --name accountEdBuilder
	docker buildx use accountEdBuilder
	docker buildx inspect accountEdBuilder --bootstrap

	echo $${GITHUB_TOKEN} | docker login ghcr.io -u $${GITHUB_USERNAME} --password-stdin
	docker buildx build --platform $${GO_OS}/$${GO_ARCH} -t ghcr.io/$${GITHUB_ACCOUNT_OWNER}/accountEd-$${GO_ARCH}-$${ENV} . --push

.PHONY: setup
setup:
	install-go
	init-go

.PHONY: install-local-deps
install-local-deps:
	go mod tidy
	go mod verify

.PHONY: install-go
install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

.PHONY: init-go
init-go:
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> $$(HOME)/.bashrc
	echo 'export PATH=$$PATH:$$(HOME)/go/bin' >> $$(HOME)/.bashrc

# compile the main application into a binary named accountEd
.PHONY: init-local-build
init-local-build:
	GOOS=$${GO_OS} GOARCH=$${GO_ARCH} go build -v -o accountEd .

register-image:
	docker push $${CONTAINER_REGISTRY}-$${ENV}-$${VERSION}

.PHONY: upgrade-go
upgrade-go:
	sudo rm -rf /usr/bin/go
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz
