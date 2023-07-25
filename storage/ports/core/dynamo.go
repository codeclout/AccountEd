package core

import (
	"context"
)

type DynamoDbCorePort interface {
	PreRegistrationConfirmationCore(ctx context.Context)
}
