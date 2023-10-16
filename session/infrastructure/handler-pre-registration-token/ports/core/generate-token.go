package core

import (
	"context"

	t "handler-pre-registration-token/token-generation-types"
)

type TokenGenerator interface {
	ProcessTokenCreation(ctx context.Context, in *t.TokenPayload) (*t.TokenCreateOut, error)
}
