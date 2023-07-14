FROM golang:1.20.5-bullseye as build

EXPOSE 8088

HEALTHCHECK --interval=30s --timeout=30s --start-period=7s --retries=3 \
    CMD curl -f http://0.0.0.0/8088/hc || exit 1

WORKDIR /usr/ci-svc-usr

COPY ./go.mod ./go.sum ./

RUN go mod tidy
RUN go mod verify

COPY ./ .

RUN CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -ldflags="-s -w" -v -o ./accountEd-member-service ./

FROM alpine:3.17.0 as prod

ARG AWS_PRE_REGISTRATION_HASH_PARAM
ENV AWS_PRE_REGISTRATION_HASH_PARAM="${AWS_PRE_REGISTRATION_HASH_PARAM}"

ARG AWS_PRE_REGISTRATION_ROLE
ENV AWS_PRE_REGISTRATION_ROLE="${AWS_PRE_REGISTRATION_ROLE}"

ARG AWS_REGION
ENV AWS_REGION="${AWS_REGION}"

ARG NOTIFICATION_SERVER_HOST
ENV NOTIFICATION_SERVER_HOST="${NOTIFICATION_SERVER_HOST}"

ARG NOTIFICATION_SERVER_PORT
ENV NOTIFICATION_SERVER_PORT="${NOTIFICATION_SERVER_PORT}"

ARG SESSION_SERVER_HOST
ENV SESSION_SERVER_HOST="${SESSION_SERVER_HOST}"

ARG SESSION_SERVER_PORT
ENV SESSION_SERVER_PORT="${SESSION_SERVER_PORT}"

ARG DOMAIN
ENV DOMAIN="${DOMAIN}"

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