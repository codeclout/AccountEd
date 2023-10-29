SHELL=/bin/bash
.EXPORT_ALL_VARIABLES:
.SHELLFLAGS = -uec

ifndef IMAGE_TAG
	IMAGE_TAG := $(shell git rev-parse --short HEAD | awk '{$$1=$$1};1')
endif

.PHONY: build-local
build-local:
	docker buildx build \
	--platform linux/arm64 --load --target=prod -t sch00l.io/$${ENVIRONMENT}-workflow-step-session-generate-token .

.PHONY: build-image
build-image:
	docker buildx build \
	--platform linux/arm64 --load \
	--build-arg=AWS_REGION=$${AWS_REGION} \
	--build-arg=ENVIRONMENT=$${ENVIRONMENT} \
	--build-arg=WORKFLOW_STEP_PRE_REGISTRATION_ROLE=$${WORKFLOW_STEP_PRE_REGISTRATION_ROLE} \
	--target=prod \
	-t $${REGISTRY}/$${REPOSITORY}:$${ENVIRONMENT}-workflow-step-session-$${IMAGE_TAG} .

.PHONY: push-image
push-image:
	docker push \
  $${REGISTRY}/$${REPOSITORY}:$${ENVIRONMENT}-workflow-step-session-$${IMAGE_TAG}

.PHONY: test-image-build
test-image-build:
	docker run -d -v ~/.aws-lambda-rie:/aws-lambda -p 8800:8080 \
  --entrypoint /aws-lambda/aws-lambda-rie \
  sch00l.io/$${ENVIRONMENT}-workflow-step-session-generate-token:latest \
      /usr/ci-svc-usr/main

.PHONY: test-image-response
test-image-response:
	curl -X POST -H "Content-Type: application/json" \
  	-d '{ "has_auto_correct": true, "member_id": "12345abc", "token_id": "67890xyz" }' \
  	http://localhost:8800/2015-03-31/functions/function/invocations
