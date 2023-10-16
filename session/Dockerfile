FROM golang:1.21.3-alpine3.18 as build

WORKDIR /usr/ci-svc-usr

COPY ./go.mod ./go.sum ./

RUN go mod tidy
RUN go mod verify

COPY ./ .

RUN CGO_ENABLED=0 GOOS=$(go env GOOS) GOARCH=$(go env GOARCH) go build -ldflags="-s -w" -v -o ./accountEd-session-service ./

FROM alpine:3.18 as prod

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

ARG ENVIRONMENT
ENV ENVIRONMENT="${ENVIRONMENT}"

ARG SESSION_PORT
ENV SESSION_PORT="${SESSION_PORT}"

ARG SESSION_STORAGE_TABLE_NAME
ENV SESSION_STORAGE_TABLE_NAME="${SESSION_STORAGE_TABLE_NAME}"

EXPOSE $SESSION_PORT

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