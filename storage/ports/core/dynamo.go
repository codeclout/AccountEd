package core

import (
	"context"

	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
)

type DynamoDbCorePort interface {
	PreRegistrationConfirmationCore(ctx context.Context) (*dynamov1.PreRegistrationConfirmationResponse, error)
}
