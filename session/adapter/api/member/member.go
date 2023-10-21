package member

import (
	"context"
	"sync"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	"github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
	"github.com/codeclout/AccountEd/session/ports/core/member"
	"github.com/codeclout/AccountEd/session/ports/framework/driven/cloud"
	drivenMemberSession "github.com/codeclout/AccountEd/session/ports/framework/driven/member"
	sessionTypes "github.com/codeclout/AccountEd/session/session-types"
)

type (
	aws                   = cloud.CredentialsAWSPort
	dmp                   = drivenMemberSession.SessionDrivenMemberPort
	mtr                   = monitoring.Adapter
	scp                   = member.SessionCoreMemberPort
	cc                    = context.Context
	gRPC                  = *protocol.AdapterServiceClients
	wait                  = *sync.WaitGroup
	ValidateTokenPayload  = sessionTypes.ValidateTokenPayload
	ValidateTokenResponse = pb.ValidateTokenResponse
)

type Adapter struct {
	config       map[string]interface{}
	core         scp
	drivenCloud  aws
	drivenMember dmp
	grpcClient   gRPC
	monitor      mtr
	wg           wait
}

func NewAdapter(config map[string]interface{}, core scp, cloud aws, dms dmp, grpc gRPC, monitor mtr, wg wait) *Adapter {
	return &Adapter{
		drivenCloud:  cloud,
		config:       config,
		core:         core,
		grpcClient:   grpc,
		monitor:      monitor,
		drivenMember: dms,
		wg:           wg,
	}
}

func (a *Adapter) ValidateMemberToken(ctx cc, awscreds []byte, in *ValidateTokenPayload, tch chan *ValidateTokenResponse, ech chan error) {
	driven, e := a.drivenCloud.GetToken(ctx, awscreds, in.Token, a.grpcClient.SessionStorageclient)
	if e != nil {
		ech <- e
		return
	}

	ok, e := a.core.ProcessTokenValidation(ctx, driven)
	if e != nil {
		ech <- e
		return
	}

	tch <- &pb.ValidateTokenResponse{IsValidToken: ok}
}
