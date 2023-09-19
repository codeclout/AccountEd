package member

import (
	"context"
	"time"

	"aidanwoods.dev/go-paseto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"

	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

type cc = context.Context

type Adapter struct {
	config  map[string]interface{}
	monitor *drivers.Adapter
}

func NewAdapter(config map[string]interface{}, monitor *drivers.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) createToken(ctx cc, in *sessiontypes.TokenPayload) (string, error) {
	serviceName, ok := a.config["ServiceName"].(string)
	if !ok {
		const msg = "unable to find service name in environment settings"
		a.monitor.LogGrpcError(ctx, msg)
		return "", status.Error(codes.FailedPrecondition, msg)
	}

	claims := make(map[string]any)
	token := paseto.NewToken()

	token.SetExpiration(in.ExpiresAt)
	token.SetIssuedAt(in.IssuedAt)
	token.SetJti(in.ID)
	token.SetNotBefore(in.IssuedAt)
	token.SetIssuer(serviceName)

	claims["member-id"] = in.MemberID
	return token.V4Sign(in.Private, nil), nil
}

func (a *Adapter) validateToken(ctx cc, publicKey, token, tokenId string) (bool, error) {
	serviceName, ok := a.config["ServiceName"].(string)
	if !ok {
		const msg = "unable to find service name in environment settings"
		a.monitor.LogGrpcError(ctx, msg)
		return false, status.Error(codes.FailedPrecondition, msg)
	}

	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.IssuedBy(serviceName))
	parser.AddRule(paseto.IdentifiedBy(tokenId))
	parser.AddRule(paseto.ValidAt(time.Now()))

	key, e := paseto.NewV4AsymmetricPublicKeyFromHex(publicKey)
	if e != nil {
		return false, e
	}

	_, e = parser.ParseV4Public(key, token, nil)
	if e != nil {
		return false, e
	}

	return true, nil
}

func (a *Adapter) ProcessTokenCreation(ctx cc, in *sessiontypes.TokenPayload) (*sessiontypes.TokenCreateOut, error) {
	token, e := a.createToken(ctx, in)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, e
	}

	out := sessiontypes.TokenCreateOut{
		Token:        token,
		TokenPayload: in,
		TTL:          in.ExpiresAt.Sub(in.IssuedAt),
	}

	return &out, nil
}

func (a *Adapter) ProcessTokenValidation(ctx cc, in *dynamov1.FetchTokenResponse) (bool, error) {
	response, e := a.validateToken(ctx, in.GetPublicKey(), in.GetToken(), in.GetTokenId())
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return false, status.Error(codes.Internal, e.Error())
	}

	return response, nil
}
