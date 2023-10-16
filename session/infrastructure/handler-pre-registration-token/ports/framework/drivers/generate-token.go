package drivers

import (
	"context"

	t "handler-pre-registration-token/token-generation-types"
)

type TokenGenerator interface {
	GenerateMemberToken(ctx context.Context, event *t.GenerateTokenRequest) (*t.GenerateTokenResponse, error)
}
