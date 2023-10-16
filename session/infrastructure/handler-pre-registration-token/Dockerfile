FROM golang:1.21.3-alpine3.18 as build

WORKDIR /usr/ci-svc-usr

COPY ./go.mod ./go.sum ./

COPY ./ .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -v -o ./main main.go

FROM alpine:3.18 as prod

WORKDIR /usr/ci-svc-usr
COPY --from=build --chown=1700:users /usr/ci-svc-usr .

ENTRYPOINT [ "/usr/ci-svc-usr/main" ]