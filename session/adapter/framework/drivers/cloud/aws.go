package cloud

import (
	"context"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	pb "github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
	"github.com/codeclout/AccountEd/pkg/session/ports/api/cloud"
	sessiontypes "github.com/codeclout/AccountEd/pkg/session/session-types"
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

func (a *Adapter) GetAWSSessionCredentials(ctx context.Context, request *pb.AWSConfigRequest) (*pb.AWSConfigResponse, error) {
	arn := request.GetArn()
	region := request.GetRegion()

	data := sessiontypes.AmazonConfigurationInput{
		ARN:    &arn,
		Region: &region,
	}

	ch := make(chan *pb.AWSConfigResponse, 1)
	errorch := make(chan error, 1)

	ctx = context.WithValue(ctx, transactionIdLogLabel, arn+"|"+region)
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
