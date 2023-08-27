package cloud

import (
	"context"

	"github.com/pkg/errors"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	coreCloud "github.com/codeclout/AccountEd/session/ports/core/cloud"
	"github.com/codeclout/AccountEd/session/ports/framework/driven/cloud"
	sessiontypes "github.com/codeclout/AccountEd/session/session-types"

	pb "github.com/codeclout/AccountEd/session/gen/aws/v1"
)

type Adapter struct {
	core    coreCloud.AWSCloudCorePort
	driven  cloud.CredentialsAWSPort
	monitor monitoring.Adapter
}

func NewAdapter(core coreCloud.AWSCloudCorePort, driven cloud.CredentialsAWSPort, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		core:    core,
		driven:  driven,
		monitor: monitor,
	}
}

func (a *Adapter) GetAWSSessionCredentials(ctx context.Context, in sessiontypes.AmazonConfigurationInput, out chan *pb.AWSConfigResponse, ech chan error) {
	if in.RoleArn == nil || in.Region == nil {
		ech <- errors.New("nil RoleArn or Region provided")
		return
	}

	sessionCredentials, e := a.driven.AssumeRoleCredentials(ctx, in.RoleArn, in.Region)
	if e != nil {
		x := errors.Wrapf(e, "api-GetAWSSessionCredentials -> driven.AssumeRoleCredentials(arn:%s,region%s)", *in.RoleArn, *in.Region)
		ech <- x
		return
	}

	core, e := a.core.ConvertCredentialsToBytes(ctx, sessionCredentials)
	if e != nil {
		ech <- e
		return
	}

	response := pb.AWSConfigResponse{AwsCredentials: core}
	out <- &response

	return
}
