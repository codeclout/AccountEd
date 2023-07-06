package cloud

import (
	"context"

	"github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
)

type MemberSessionDriverPort interface {
	GetEncryptedSessionId(ctx context.Context, request *sessions.EncryptedStringRequest) (*sessions.EncryptedStringResponse, error)
}
