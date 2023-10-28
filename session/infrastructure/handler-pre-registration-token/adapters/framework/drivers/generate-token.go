package drivers

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/pkg/errors"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	"handler-pre-registration-token/ports/api"
	t "handler-pre-registration-token/token-generation-types"
)

type LambdaName string

type Adapter struct {
	api     api.TokenGenerator
	config  map[string]interface{}
	monitor monitoring.Adapter
}

func NewAdapter(api api.TokenGenerator, config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		api:     api,
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) setContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	deadline, _ := ctx.Deadline()
	name, ok := a.config["ServiceName"].(string)
	if !ok {
		a.monitor.LogGenericError("ServiceName not available in environment")
		os.Exit(1)
	}

	n := LambdaName(name)

	ctx = context.WithValue(ctx, n, lambdacontext.FunctionName)
	ctx, cancel := context.WithDeadline(ctx, deadline)

	return ctx, cancel
}

func (a *Adapter) processMemberTokenGeneration(in *t.GenerateTokenRequest) (*t.GenerateTokenRequest, error) {
	var s string

	if in == nil || in.MemberId == (s) || in.TokenId == (s) {
		return nil, errors.New("invalid new token request")
	}

	return in, nil
}

func (a *Adapter) GenerateMemberToken(ctx context.Context, event *t.GenerateTokenRequest) (*t.GenerateTokenResponse, error) {
	ech := make(chan error, 1)
	tch := make(chan *t.GenerateTokenResponse, 1)

	apiData, e := a.processMemberTokenGeneration(event)
	if e != nil {
		return nil, e
	}

	ctx, cancel := a.setContextTimeout(ctx)

	defer cancel()
	a.api.CreateMemberToken(ctx, apiData, tch, ech)

	select {
	case <-ctx.Done():
		const msg = "token generation request timeout"
		a.monitor.Logger.ErrorContext(ctx, msg)
		return nil, errors.New(msg)

	case out := <-tch:
		a.monitor.LogGenericInfo("success")
		return out, nil

	case e := <-ech:
		a.monitor.Logger.ErrorContext(ctx, e.Error())
		return nil, e
	}
}
