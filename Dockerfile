FROM golang:1.19.1-bullseye as build

ARG ENV
ENV ENVIRONMENT="${ENV}"

EXPOSE 8088

HEALTHCHECK --interval=30s --timeout=30s --start-period=7s --retries=3 \
    CMD curl -f http://0.0.0.0/8088/hc || exit 1

WORKDIR /usr/ci-svc-usr

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./backend/go.mod ./backend/go.sum ./

RUN go mod tidy
RUN go mod verify

COPY ./backend .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s" -v -o ./accountEd ./

FROM alpine:3.14.6 as prod

RUN set -ex \
    && apk update \
    && apk upgrade \
    && adduser -D -u 1700 ci-svc-usr -G users users \
    && echo "ci-svc-usr ALL=(ALL:ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chmod 0440 /etc/sudoers

RUN su - ci-svc-usr && apk add -U --no-cache ca-certificates

WORKDIR /usr/ci-svc-usr
COPY --from=build --chown=1700:users /usr/ci-svc-usr .

CMD [ "/usr/ci-svc-usr/accountEd" ]