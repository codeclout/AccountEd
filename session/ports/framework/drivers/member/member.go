package member

import (
	"context"

	"github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
)

type SessionDriverMemberPort interface {
	GetEncryptedSessionId(ctx context.Context, request *sessions.EncryptedStringRequest) (*sessions.EncryptedStringResponse, error)
}
