package api

import (
	"context"

	t "handler-pre-registration-token/token-generation-types"
)

type TokenGenerator interface {
	CreateMemberToken(ctx context.Context, in *t.GenerateTokenRequest, tch chan *t.GenerateTokenResponse, ech chan error)
}
