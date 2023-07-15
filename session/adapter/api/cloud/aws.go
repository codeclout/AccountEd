package cloud

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	pb "github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
	core "github.com/codeclout/AccountEd/pkg/session/ports/core/cloud"
	"github.com/codeclout/AccountEd/pkg/session/ports/framework/driven/cloud"
	sessiontypes "github.com/codeclout/AccountEd/pkg/session/session-types"
)

type Adapter struct {
	core   core.AWSCloudCorePort
	driven cloud.CredentialsAWSPort
	log    *slog.Logger
}

func NewAdapter(core core.AWSCloudCorePort, driven cloud.CredentialsAWSPort, log *slog.Logger) *Adapter {
	return &Adapter{
		core:   core,
		driven: driven,
		log:    log,
	}
}

func (a *Adapter) GetAWSSessionCredentials(ctx context.Context, in sessiontypes.AmazonConfigurationInput, out chan *pb.AWSConfigResponse, echan chan error) {
	// a.core.GetServiceIdMetadata(ctx)

	session, e := a.driven.AssumeRoleCredentials(ctx, in.ARN, in.Region)
	if e != nil {
		x := errors.Wrapf(e, "api-GetAWSSessionCredentials -> driven.AssumeRoleCredentials(arn:%s,region%s)", *in.ARN, *in.Region)
		echan <- x
		return
	}

	b, e := json.Marshal(session.Credentials)
	if e != nil {
		echan <- errors.Wrap(e, "api-GetAWSSessionCredentials -> json.Marshal(session)")
		return
	}

	response := pb.AWSConfigResponse{AwsSession: b}
	out <- &response

	return
}
