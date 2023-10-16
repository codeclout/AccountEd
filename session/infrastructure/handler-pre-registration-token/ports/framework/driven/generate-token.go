package driven

import (
	"context"
	"time"

	t "handler-pre-registration-token/token-generation-types"
)

type TokenGenerator interface {
	GetTokenPayload(ctx context.Context, memberId, tokenId string, ttl time.Duration) (*t.TokenPayload, error)
}
