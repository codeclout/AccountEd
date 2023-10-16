package driven

import (
	"context"
	"errors"
	"time"

	"aidanwoods.dev/go-paseto"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	t "handler-pre-registration-token/token-generation-types"
)

type Adapter struct {
	monitor monitoring.Adapter
}

func NewAdapter(monitor monitoring.Adapter) *Adapter {
	return &Adapter{monitor: monitor}
}

func (a *Adapter) newTokenPayload(memberId, tokenId string, ttl time.Duration) (*t.TokenPayload, error) {
	var s string

	if memberId == (s) || tokenId == (s) {
		return nil, errors.New("member and session id are required")
	}

	privateKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := privateKey.Public()

	tp := t.TokenPayload{
		ExpiresAt:     time.Now().Add(ttl),
		ID:            tokenId,
		IssuedAt:      time.Now(),
		MemberID:      memberId,
		Private:       privateKey,
		Public:        publicKey,
		PublicEncoded: publicKey.ExportHex(),
	}

	return &tp, nil
}

func (a *Adapter) GetTokenPayload(ctx context.Context, memberId, tokenId string, ttl time.Duration) (*t.TokenPayload, error) {
	out, e := a.newTokenPayload(memberId, tokenId, ttl)
	if e != nil {
		a.monitor.Logger.ErrorContext(ctx, e.Error())
		return nil, e
	}

	return out, nil
}
