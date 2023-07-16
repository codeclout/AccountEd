package cloud

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	core "github.com/codeclout/AccountEd/session/ports/core/cloud"
	"github.com/codeclout/AccountEd/session/ports/framework/driven/cloud"
	sessiontypes "github.com/codeclout/AccountEd/session/session-types"

	pb "github.com/codeclout/AccountEd/session/gen/aws/v1"
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

	response := pb.AWSConfigResponse{AwsCredentials: b}
	out <- &response

	return
}
