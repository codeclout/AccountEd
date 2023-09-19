package member

import (
	"context"
	"time"

	"aidanwoods.dev/go-paseto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

type Adapter struct {
	config  map[string]interface{}
	monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) newTokenPayload(memberId, tokenId string, ttl time.Duration) (*sessiontypes.TokenPayload, error) {
	var s string

	if memberId == (s) || tokenId == (s) {
		return nil, status.Error(codes.InvalidArgument, "member and session id are required")
	}

	privateKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := privateKey.Public()

	tp := sessiontypes.TokenPayload{
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

func (a *Adapter) GetTokenPayload(ctx context.Context, memberId, tokenId string, ttl time.Duration) (*sessiontypes.TokenPayload, error) {
	out, e := a.newTokenPayload(memberId, tokenId, ttl)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, e
	}

	return out, nil
}
