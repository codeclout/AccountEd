FROM golang:1.20.5-bullseye as build

EXPOSE 8088

HEALTHCHECK --interval=30s --timeout=30s --start-period=7s --retries=3 \
    CMD curl -f http://0.0.0.0/8088/hc || exit 1

WORKDIR /usr/ci-svc-usr

COPY ./go.mod ./go.sum ./

RUN go mod tidy
RUN go mod verify

COPY ./ .

RUN CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -ldflags="-s -w" -v -o ./accountEd-notification-service ./

FROM alpine:3.17.0 as prod

ARG AWS_REGION
ENV AWS_REGION="${AWS_REGION}"

ARG EMAIL_PROCESSOR_API_KEY
ENV EMAIL_PROCESSOR_API_KEY="${EMAIL_PROCESSOR_API_KEY}"

ARG EMAIL_PROCESSOR_DOMAIN
ENV EMAIL_PROCESSOR_DOMAIN="${EMAIL_PROCESSOR_DOMAIN}"

ARG EMAIL_VERIFIER_API_PATH
ENV EMAIL_VERIFIER_API_PATH="${EMAIL_VERIFIER_API_PATH}"

ARG PERFORMANCE_SLA
ENV PERFORMANCE_SLA="${PERFORMANCE_SLA}"

ARG PORT
ENV PORT="${PORT}"

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