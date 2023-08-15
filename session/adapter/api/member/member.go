package member

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/grpc/status"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	"github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	"github.com/codeclout/AccountEd/session/ports/core/member"
	"github.com/codeclout/AccountEd/session/ports/framework/driven/cloud"
	drivenMemberSession "github.com/codeclout/AccountEd/session/ports/framework/driven/member"
	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"

	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
)

type aws = cloud.CredentialsAWSPort
type dmp = drivenMemberSession.SessionDrivenMemberPort
type mtr = monitoring.Adapter
type scp = member.SessionCoreMemberPort

type cctx = context.Context
type gcpc = *protocol.AdapterServiceClients
type wait = *sync.WaitGroup

type sessionIdResp = pb.EncryptedStringResponse
type sessionIdData = sessiontypes.SessionIdEncryptionOut
type storeMeta = sessiontypes.SessionStoreMetadata
type storeSessResp = dynamov1.PreRegistrationConfirmationResponse

type Adapter struct {
	config             map[string]interface{}
	contextAPILabel    sessiontypes.ContextAPILabel
	contextDrivenLabel sessiontypes.ContextDrivenLabel
	core               scp
	drivenCloud        aws
	drivenMember       dmp
	grpcClient         gcpc
	monitor            mtr
	wg                 wait
}

func NewAdapter(config map[string]interface{}, core scp, cloud aws, dms dmp, grpc gcpc, monitor mtr, wg wait) *Adapter {
	return &Adapter{
		drivenCloud:        cloud,
		config:             config,
		contextAPILabel:    "api_input",
		contextDrivenLabel: "driven_input",
		core:               core,
		grpcClient:         grpc,
		monitor:            monitor,
		drivenMember:       dms,
		wg:                 wg,
	}
}

func (a *Adapter) EncryptSessionId(ctx cctx, awscreds []byte, in *storeMeta, uch chan *sessionIdResp, ech chan error) {
	k, e := a.drivenMember.GetSessionIdKey(ctx, awscreds)
	if e != nil {
		x := errors.Wrapf(e, "api-EncryptSessionId -> core.ProcessSessionIdEncryption(sessionID:%s)", in.SessionID)
		ech <- x
		return
	}

	ctx = context.WithValue(ctx, a.contextDrivenLabel, *k)
	ctx = context.WithValue(ctx, a.contextAPILabel, in.SessionID)

	core, e := a.core.ProcessSessionIdEncryption(ctx)
	if e != nil {
		x := errors.Wrapf(e, "api-EncryptSessionId -> core.ProcessSessionIdEncryption(sessionID:%s)", in.SessionID)
		ech <- x
		return
	}

	go a.storeEncryptedSession(ctx, core, *in, awscreds)

	out := pb.EncryptedStringResponse{
		EncryptedSessionId: *core.CipherText,
	}

	uch <- &out
	return
}

func (a *Adapter) storeEncryptedSession(ctx cctx, in *sessionIdData, meta storeMeta, staticCredentials []byte) {
	a.wg.Add(1)

	tableName, ok := a.config["SessionTableName"].(string)
	if !ok {
		a.monitor.LogGrpcError(ctx, "session table name not set in environment")
	}

	data := dynamov1.PreRegistrationConfirmationRequest{
		AssociatedData:               in.AssociatedData,
		EncryptedSessionID:           *in.CipherText,
		ForwardedIp:                  "",
		HasAutoCorrect:               meta.HasAutoCorrect,
		MemberId:                     meta.MemberID,
		Nonce:                        in.IV,
		SessionServiceAWScredentials: staticCredentials,
		SessionID:                    *in.SessionID,
		SessionTableName:             tableName,
		Ttl:                          1000 * 60 * 15,
	}

	// @TODO - Handle system event from response data -> new member pre confirmation
	client := *a.grpcClient.SessionStorageclient
	_, e := client.StorePreConfirmationRegistrationSession(ctx, &data)

	if e != nil {
		ko, ok := status.FromError(e)

		if !ok {
			panic(fmt.Sprintf("unexpected error %v", e))
		}

		a.monitor.LogGrpcError(ctx, fmt.Sprintf("session store operation failed ->  %v", ko))
		_, _ = client.StorePreConfirmationRegistrationSession(ctx, &data)
	}

	a.wg.Done()
}
