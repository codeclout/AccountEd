package cloud

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/codeclout/AccountEd/pkg/monitoring"
)

type StaticCredentialsProvider = credentials.StaticCredentialsProvider

type Adapter struct {
	config  map[string]interface{}
	monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) ConvertCredentialsToBytes(ctx context.Context, in *StaticCredentialsProvider) ([]byte, error) {
	if _, e := in.Retrieve(ctx); e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil, status.Error(codes.FailedPrecondition, "invalid cloud credentials")
	}

	b, e := json.Marshal(in)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, status.Error(codes.Internal, e.Error())
	}

	return b, nil
}
