FROM golang:1.21.3-alpine3.18 as build

WORKDIR /usr/ci-svc-usr

COPY ./go.mod ./go.sum ./
RUN go mod donwload

COPY ./ .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -v -o ./main main.go

FROM alpine:3.18 as prod

ARG AWS_REGION
ENV AWS_REGION="${AWS_REGION}"

ARG ENVIRONMENT
ENV ENVIRONMENT="${ENVIRONMENT}"

ARG WORKFLOW_STEP_PRE_REGISTRATION_ROLE
ENV AWS_ROLE_TO_ASSUME="${WORKFLOW_STEP_PRE_REGISTRATION_ROLE}"

WORKDIR /usr/ci-svc-usr
COPY --from=build --chown=1700:users /usr/ci-svc-usr .

ENTRYPOINT [ "/usr/ci-svc-usr/main" ]