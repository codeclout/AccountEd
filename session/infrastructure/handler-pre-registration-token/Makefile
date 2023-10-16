SHELL=/bin/bash
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS = -uec

.PHONY: build-image
build-image:
	docker buildx install
	-docker buildx rm SessionPreRegistrationTokenLCI
	docker buildx create --name SessionPreRegistrationTokenLCI
	docker buildx use SessionPreRegistrationTokenLCI
	docker buildx inspect SessionPreRegistrationTokenLCI --bootstrap

	docker buildx build \
	--platform linux/arm64 --load --target=prod -t sch00l.io/session-pre-registration-token-$${ENVIRONMENT} .

.PHONY: test-image-build
test-image-build:
	docker run -d -v ~/.aws-lambda-rie:/aws-lambda -p 8800:8080 \
  --entrypoint /aws-lambda/aws-lambda-rie \
  sch00l.io/session-pre-registration-token-$${ENVIRONMENT}:latest \
      /usr/ci-svc-usr/main

.PHONY: test-image-response
test-image-response:
	curl -X POST -H "Content-Type: application/json" \
  	-d '{ "has_auto_correct": true, "member_id": "12345abc", "token_id": "67890xyz" }' \
  	http://localhost:8800/2015-03-31/functions/function/invocations
