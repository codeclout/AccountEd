FROM golang:1.21.1-alpine3.18 as build

WORKDIR /usr/ci-svc-usr

COPY ./go.mod ./go.sum ./

RUN go mod tidy
RUN go mod verify

COPY ./ .

RUN CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -ldflags="-s -w" -v -o ./accountEd-member-service ./

FROM alpine:3.18 as prod

ARG AWS_PRE_REGISTRATION_HASH_PARAM
ENV AWS_PRE_REGISTRATION_HASH_PARAM="${AWS_PRE_REGISTRATION_HASH_PARAM}"

ARG AWS_PRE_REGISTRATION_ROLE
ENV AWS_PRE_REGISTRATION_ROLE="${AWS_PRE_REGISTRATION_ROLE}"

ARG AWS_REGION
ENV AWS_REGION="${AWS_REGION}"

ARG DOMAIN
ENV DOMAIN="${DOMAIN}"

ARG ENVIRONMENT
ENV ENVIRONMENT="${ENVIRONMENT}"

ARG MEMBER_PORT
ENV MEMBER_PORT="${MEMBER_PORT}"

RUN set -ex \
    && apk update \
    && apk upgrade \
    && adduser -D -u 1700 ci-svc-usr -G users users \
    && echo "ci-svc-usr ALL=(ALL:ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chmod 0440 /etc/sudoers

RUN su - ci-svc-usr && apk add -U --no-cache ca-certificates

WORKDIR /usr/ci-svc-usr
COPY --from=build --chown=1700:users /usr/ci-svc-usr .

CMD [ "/usr/ci-svc-usr/accountEd-member-service" ]