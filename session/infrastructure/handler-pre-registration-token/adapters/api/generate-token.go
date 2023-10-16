package api

import (
	"context"
	"errors"
	"time"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	"handler-pre-registration-token/ports/core"
	"handler-pre-registration-token/ports/framework/driven"
	t "handler-pre-registration-token/token-generation-types"
)

type cc = context.Context

type Adapter struct {
	core    core.TokenGenerator
	driven  driven.TokenGenerator
	monitor monitoring.Adapter
}

func NewAdapter(core core.TokenGenerator, driven driven.TokenGenerator, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		core:    core,
		driven:  driven,
		monitor: monitor,
	}
}

func (a *Adapter) CreateMemberToken(ctx cc, in *t.GenerateTokenRequest, tch chan *t.GenerateTokenResponse, ech chan error) {
	var sessionExpiry = time.Hour * 24

	if in == nil {
		const msg = "request to create token received nil input"
		a.monitor.LogGrpcError(ctx, msg)

		ech <- errors.New(msg)
		return
	}

	drivenResult, e := a.driven.GetTokenPayload(ctx, in.MemberId, in.TokenId, sessionExpiry)
	if e != nil {
		ech <- e
		return
	}

	coreResult, e := a.core.ProcessTokenCreation(ctx, drivenResult)
	if e != nil {
		ech <- e
		return
	}

	out := t.GenerateTokenResponse{
		Token: coreResult.Token,
	}

	tch <- &out
}
