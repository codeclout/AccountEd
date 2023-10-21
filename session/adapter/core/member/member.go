package member

import (
	"context"
	"time"

	"aidanwoods.dev/go-paseto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
)

type cc = context.Context

type Adapter struct {
	config  map[string]interface{}
	monitor *monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor *monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
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

func (a *Adapter) ProcessTokenValidation(ctx cc, in *dynamov1.FetchTokenResponse) (bool, error) {
	response, e := a.validateToken(ctx, in.GetPublicKey(), in.GetToken(), in.GetTokenId())
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return false, status.Error(codes.Internal, e.Error())
	}

	return response, nil
}
