package cloud

import (
	"context"

	"github.com/pkg/errors"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	awspb "github.com/codeclout/AccountEd/session/gen/aws/v1"

	"github.com/codeclout/AccountEd/session/ports/api/cloud"
	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

type Adapter struct {
	api     cloud.AWSApiPort
	config  map[string]interface{}
	monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, api cloud.AWSApiPort, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		api:     api,
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) GetAWSSessionCredentials(ctx context.Context, request *awspb.AWSConfigRequest) (*awspb.AWSConfigResponse, error) {
	region := request.GetRegion()
	roleArn := request.GetRoleArn()

	data := sessiontypes.AmazonConfigurationInput{
		RoleArn: &roleArn,
		Region:  &region,
	}

	ch := make(chan *awspb.AWSConfigResponse, 1)
	ech := make(chan error, 1)

	ctx = context.WithValue(ctx, a.monitor.LogLabelTransactionID, roleArn+"|"+region)
	a.api.GetAWSSessionCredentials(ctx, data, ch, ech)

	select {
	case <-ctx.Done():
		a.monitor.LogGrpcError(ctx, "request timeout")
		return nil, errors.New("request timeout")

	case out := <-ch:
		a.monitor.LogGrpcInfo(ctx, "success")
		return out, nil

	case e := <-ech:
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, e
	}
}
