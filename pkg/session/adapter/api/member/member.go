package member

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	"github.com/codeclout/AccountEd/pkg/session/ports/core/member"
	"github.com/codeclout/AccountEd/pkg/session/ports/framework/driven/cloud"
	member2 "github.com/codeclout/AccountEd/pkg/session/ports/framework/driven/member"

	pb "github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
)

type Adapter struct {
	config       map[string]interface{}
	core         member.SessionCoreMemberPort
	awsdriven    cloud.CredentialsAWSPort
	log          *slog.Logger
	memberdriven member2.SessionDrivenMemberPort
}

func NewAdapter(config map[string]interface{}, core member.SessionCoreMemberPort, awsdriven cloud.CredentialsAWSPort, memberdriven member2.SessionDrivenMemberPort, log *slog.Logger) *Adapter {
	return &Adapter{
		awsdriven:    awsdriven,
		config:       config,
		core:         core,
		log:          log,
		memberdriven: memberdriven,
	}
}

func (a *Adapter) EncryptSessionId(ctx context.Context, awscredentials []byte, id string, uch chan *pb.EncryptedStringResponse, echan chan error) {
	k, e := a.memberdriven.GetSessionIdKey(ctx, awscredentials)
	if e != nil {
		x := errors.Wrapf(e, "api-EncryptSessionId -> core.ProcessSessionIdEncryption(sessionID:%s)", id)
		echan <- x
		return
	}

	encryptedString, e := a.core.ProcessSessionIdEncryption(ctx, id, *k)
	if e != nil {
		x := errors.Wrapf(e, "api-EncryptSessionId -> core.ProcessSessionIdEncryption(sessionID:%s)", id)
		echan <- x
		return
	}

	out := pb.EncryptedStringResponse{
		EncryptedSessionId: *encryptedString.CipherText,
	}

	uch <- &out
	return
}
