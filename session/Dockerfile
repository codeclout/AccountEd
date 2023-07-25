FROM golang:1.20.5-bullseye as build

WORKDIR /usr/ci-svc-usr

COPY ./go.mod ./go.sum ./

RUN go mod tidy
RUN go mod verify

COPY ./ .

RUN CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -ldflags="-s -w" -v -o ./accountEd-session-service ./

FROM alpine:3.17.0 as prod

ARG AWS_ACCESS_KEY_ID
ENV AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}"

ARG AWS_PRE_REGISTRATION_PARAM
ENV AWS_PRE_REGISTRATION_PARAM="${AWS_PRE_REGISTRATION_PARAM}"

ARG AWS_REGION
ENV AWS_REGION="${AWS_REGION}"

ARG AWS_ROLE_TO_ASSUME
ENV AWS_ROLE_TO_ASSUME="${AWS_ROLE_TO_ASSUME}"

ARG AWS_SECRET_ACCESS_KEY
ENV AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}"

ARG PORT
ENV PORT="${PORT}"

EXPOSE $PORT

HEALTHCHECK --interval=30s --timeout=30s --start-period=7s --retries=3 \
    CMD curl -f "http://0.0.0.0/$PORT/hc" || exit 1

RUN set -ex \
    && apk update \
    && apk upgrade \
    && adduser -D -u 1700 ci-svc-usr -G users users \
    && echo "ci-svc-usr ALL=(ALL:ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chmod 0440 /etc/sudoers

RUN su - ci-svc-usr && apk add -U --no-cache ca-certificates

WORKDIR /usr/ci-svc-usr
COPY --from=build --chown=1700:users /usr/ci-svc-usr .

CMD [ "/usr/ci-svc-usr/accountEd-session-service" ]