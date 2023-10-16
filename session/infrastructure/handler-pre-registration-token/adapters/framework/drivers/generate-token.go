package drivers

import (
	"context"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/pkg/errors"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	"handler-pre-registration-token/ports/api"
	t "handler-pre-registration-token/token-generation-types"
)

type Adapter struct {
	api     api.TokenGenerator
	monitor monitoring.Adapter
}

func NewAdapter(api api.TokenGenerator, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		api:     api,
		monitor: monitor,
	}
}

func (a *Adapter) setContextTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	deadline, _ := ctx.Deadline()

	ctx = context.WithValue(ctx, "function_name", lambdacontext.FunctionName)
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
