package driven

import "context"

type DynamoDbDrivenPort interface {
	StoreSession(ctx context.Context) error
}
