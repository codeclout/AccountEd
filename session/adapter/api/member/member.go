package member

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	"github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
	"github.com/codeclout/AccountEd/session/ports/core/member"
	"github.com/codeclout/AccountEd/session/ports/framework/driven/cloud"
	drivenMemberSession "github.com/codeclout/AccountEd/session/ports/framework/driven/member"
	sessionTypes "github.com/codeclout/AccountEd/session/session-types"
)

type aws = cloud.CredentialsAWSPort
type dmp = drivenMemberSession.SessionDrivenMemberPort
type mtr = monitoring.Adapter
type scp = member.SessionCoreMemberPort

type cc = context.Context
type gRPC = *protocol.AdapterServiceClients
type wait = *sync.WaitGroup

type GenerateTokenResponse = pb.GenerateTokenResponse
type NewTokenPayload = sessionTypes.NewTokenPayload
type ValidateTokenPayload = sessionTypes.ValidateTokenPayload
type ValidateTokenResponse = pb.ValidateTokenResponse

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

func (a *Adapter) CreateMemberToken(ctx cc, awscreds []byte, in *NewTokenPayload, tch chan *GenerateTokenResponse, ech chan error) {
	var sessionExpiry = time.Hour

	if in == nil {
		const msg = "request to encrypt session id received nil input"
		a.monitor.LogGrpcError(ctx, msg)

		ech <- status.Error(codes.InvalidArgument, msg)
		return
	}

	driven, e := a.drivenMember.GetTokenPayload(ctx, in.MemberId, in.TokenId, sessionExpiry)
	if e != nil {
		ech <- e
		return
	}

	core, e := a.core.ProcessTokenCreation(ctx, driven)
	if e != nil {
		ech <- e
		return
	}

	e = a.drivenCloud.StoreToken(ctx, a.grpcClient.SessionStorageclient, core, in.HasAutoCorrect, awscreds)
	if e != nil {
		ech <- e
		return
	}

	out := pb.GenerateTokenResponse{
		Token: core.Token,
	}

	tch <- &out
}
