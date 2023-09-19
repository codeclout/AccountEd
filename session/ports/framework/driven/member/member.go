package member

import (
	"context"
	"time"

	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

type SessionDrivenMemberPort interface {
	GetTokenPayload(ctx context.Context, memberId, tokenId string, ttl time.Duration) (*sessiontypes.TokenPayload, error)
}
