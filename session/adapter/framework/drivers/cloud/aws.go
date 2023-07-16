package cloud

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	awspb "github.com/codeclout/AccountEd/session/gen/aws/v1"

	"github.com/codeclout/AccountEd/session/ports/api/cloud"
	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

var defaultRouteDuration = sessiontypes.DefaultRouteDuration(2000)
var transactionIdLogLabel = sessiontypes.LogLabel("transactionID")

type Adapter struct {
	api    cloud.AWSApiPort
	config map[string]interface{}
	log    *slog.Logger
}

func NewAdapter(config map[string]interface{}, api cloud.AWSApiPort, log *slog.Logger) *Adapter {
	return &Adapter{
		api:    api,
		config: config,
		log:    log,
	}
}

func (a *Adapter) GetAWSSessionCredentials(ctx context.Context, request *awspb.AWSConfigRequest) (*awspb.AWSConfigResponse, error) {
	region := request.GetRegion()
	roleArn := request.GetRoleArn()

	data := sessiontypes.AmazonConfigurationInput{
		ARN:    &roleArn,
		Region: &region,
	}

	ch := make(chan *awspb.AWSConfigResponse, 1)
	errorch := make(chan error, 1)

	ctx = context.WithValue(ctx, transactionIdLogLabel, roleArn+"|"+region)
	a.api.GetAWSSessionCredentials(ctx, data, ch, errorch)

	select {
	case <-ctx.Done():
		t := ctx.Value(transactionIdLogLabel)
		a.log.Error("request timeout", transactionIdLogLabel, t.(string))
		return nil, errors.New("request timeout")

	case out := <-ch:
		t := ctx.Value(transactionIdLogLabel)
		a.log.Info("success", transactionIdLogLabel, t.(string))
		return out, nil

	case e := <-errorch:
		t := ctx.Value(transactionIdLogLabel)
		a.log.Error(e.Error(), transactionIdLogLabel, t.(string))
		return nil, e
	}
}
