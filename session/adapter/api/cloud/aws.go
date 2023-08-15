package cloud

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	core "github.com/codeclout/AccountEd/session/ports/core/cloud"
	"github.com/codeclout/AccountEd/session/ports/framework/driven/cloud"
	sessiontypes "github.com/codeclout/AccountEd/session/session-types"

	pb "github.com/codeclout/AccountEd/session/gen/aws/v1"
)

type Adapter struct {
	core    core.AWSCloudCorePort
	driven  cloud.CredentialsAWSPort
	monitor monitoring.Adapter
}

func NewAdapter(core core.AWSCloudCorePort, driven cloud.CredentialsAWSPort, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		core:    core,
		driven:  driven,
		monitor: monitor,
	}
}

func (a *Adapter) GetAWSSessionCredentials(ctx context.Context, in sessiontypes.AmazonConfigurationInput, out chan *pb.AWSConfigResponse, echan chan error) {
	// a.core.GetServiceIdMetadata(ctx)

	sessionCredentials, e := a.driven.AssumeRoleCredentials(ctx, in.RoleArn, in.Region)
	if e != nil {
		x := errors.Wrapf(e, "api-GetAWSSessionCredentials -> driven.AssumeRoleCredentials(arn:%s,region%s)", *in.RoleArn, *in.Region)
		echan <- x
		return
	}

	b, e := json.Marshal(sessionCredentials)
	if e != nil {
		echan <- errors.Wrap(e, "api-GetAWSSessionCredentials -> json.Marshal(sessionCredentials)")
		return
	}

	response := pb.AWSConfigResponse{AwsCredentials: b}
	out <- &response

	return
}
