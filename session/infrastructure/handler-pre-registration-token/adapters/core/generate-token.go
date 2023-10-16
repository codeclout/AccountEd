package core

import (
	"context"

	"aidanwoods.dev/go-paseto"
	"github.com/pkg/errors"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	t "handler-pre-registration-token/token-generation-types"
)

type Adapter struct {
	monitor monitoring.Adapter
}

func NewAdapter(monitor monitoring.Adapter) *Adapter {
	return &Adapter{monitor: monitor}
}

func (a *Adapter) createToken(ctx context.Context, in *t.TokenPayload) (string, error) {
	fnName, ok := ctx.Value("function_name").(string)
	if !ok {
		const msg = "unable to find service name in environment settings"
		a.monitor.LogGenericError(msg)
		return "", errors.New(msg)
	}

	claims := make(map[string]any)
	token := paseto.NewToken()

	token.SetExpiration(in.ExpiresAt)
	token.SetIssuedAt(in.IssuedAt)
	token.SetJti(in.ID)
	token.SetNotBefore(in.IssuedAt)
	token.SetIssuer(fnName)

	claims["member-id"] = in.MemberID
	return token.V4Sign(in.Private, nil), nil
}

func (a *Adapter) ProcessTokenCreation(ctx context.Context, in *t.TokenPayload) (*t.TokenCreateOut, error) {
	token, e := a.createToken(ctx, in)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, e
	}

	out := t.TokenCreateOut{
		Token:        token,
		TokenPayload: in,
		TTL:          in.ExpiresAt.Sub(in.IssuedAt),
	}

	return &out, nil
}
