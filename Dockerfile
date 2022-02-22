FROM golang:1.17.7-bullseye as build

ARG ENV

ENV ENVIRONMENT="${ENV}"

RUN env GOOS=$(go env GOOS)
RUN env GOARCH=$(go env GOARCH)

WORKDIR /usr/ci-svc-accountEd

RUN set -ex \
    && apt update \
    && apt upgrade -y \
    && apt full-upgrade -y \
    && apt --purge autoremove \
    && adduser --disabled-password --gecos "" ci-svc-accountEd --force-badname \
    && usermod -aG users ci-svc-accountEd \
    && echo "ci-svc-accountEd ALL=(ALL:ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chmod 0440 /etc/sudoers

RUN su - ci-svc-accountEd

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
COPY ./config.hcl ./

RUN go mod tidy 
RUN go mod verify
RUN go get github.com/hashicorp/hcl/v2/hclsimple

COPY ./src .

RUN env CGO_ENABLED=0
RUN env GOOS="${GO_OS}" 
RUN env GOARCH="${GO_ARCH}"

RUN cd cmd && go build -v -o accountEd ./...



FROM alpine:latest as prod

RUN set -ex \
    && apk update \
    && apk upgrade \
    && adduser -D -u 1700 ci-svc-accountEd -G users users \
    && echo "ci-svc-accountEd ALL=(ALL:ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chmod 0440 /etc/sudoers

RUN su - ci-svc-accountEd

WORKDIR /usr/ci-svc-accountEd
COPY --from=build /usr/ci-svc-accountEd .

ENTRYPOINT [ "cmd/accountEd" ]