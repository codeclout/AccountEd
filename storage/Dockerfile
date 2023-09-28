FROM golang:1.21.1-alpine3.18 as build

WORKDIR /usr/ci-svc-usr

COPY ./go.mod ./go.sum ./

RUN go mod tidy
RUN go mod verify

COPY ./ .

RUN CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -ldflags="-s -w" -v -o ./accountEd-storage-service ./

FROM alpine:3.18 as prod

ARG AWS_REGION
ENV AWS_REGION="${AWS_REGION}"

ARG AWS_ROLE_TO_ASSUME
ENV AWS_ROLE_TO_ASSUME="${AWS_ROLE_TO_ASSUME}"

ARG ENVIRONMENT
ENV ENVIRONMENT="${ENVIRONMENT}"

ARG STORAGE_PORT
ENV STORAGE_PORT="${STORAGE_PORT}"

RUN set -ex \
    && apk update \
    && apk upgrade \
    && adduser -D -u 1700 ci-svc-usr -G users users \
    && echo "ci-svc-usr ALL=(ALL:ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chmod 0440 /etc/sudoers

RUN su - ci-svc-usr && apk add -U --no-cache ca-certificates

WORKDIR /usr/ci-svc-usr
COPY --from=build --chown=1700:users /usr/ci-svc-usr .

CMD [ "/usr/ci-svc-usr/accountEd-storage-service" ]