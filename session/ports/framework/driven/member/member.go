package member

import "context"

type SessionDrivenMemberPort interface {
	GetSessionIdKey(ctx context.Context, awsconfig []byte) (*string, error)
}
