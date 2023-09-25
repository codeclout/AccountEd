FROM golang:1.21.1-alpine3.18 as build

WORKDIR /usr/ci-svc-usr

COPY ./go.mod ./go.sum ./

RUN go mod tidy
RUN go mod verify

COPY ./ .

RUN CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -ldflags="-s -w" -a -v -o ./accountEd-notification-service ./

FROM alpine:3.18 as prod

ARG AWS_REGION
ENV AWS_REGION="${AWS_REGION}"

ARG EMAIL_PROCESSOR_API_KEY
ENV EMAIL_PROCESSOR_API_KEY="${EMAIL_PROCESSOR_API_KEY}"

ARG EMAIL_PROCESSOR_DOMAIN
ENV EMAIL_PROCESSOR_DOMAIN="${EMAIL_PROCESSOR_DOMAIN}"

ARG EMAIL_VERIFIER_API_PATH
ENV EMAIL_VERIFIER_API_PATH="${EMAIL_VERIFIER_API_PATH}"

ARG NOTIFICATIONS_PORT
ENV NOTIFICATIONS_PORT="${NOTIFICATIONS_PORT}"

ARG PERFORMANCE_SLA
ENV PERFORMANCE_SLA="${PERFORMANCE_SLA}"

EXPOSE $NOTIFICATIONS_PORT

RUN set -ex \
    && apk update \
    && apk upgrade \
    && adduser -D -u 1700 ci-svc-usr -G users users \
    && echo "ci-svc-usr ALL=(ALL:ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chmod 0440 /etc/sudoers

RUN su - ci-svc-usr && apk add -U --no-cache ca-certificates

WORKDIR /usr/ci-svc-usr
COPY --from=build --chown=1700:users /usr/ci-svc-usr .

CMD [ "/usr/ci-svc-usr/accountEd-notification-service" ]