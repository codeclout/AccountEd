package cloud

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/credentials"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

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

func (a *Adapter) ConvertCredentialsToBytes(ctx context.Context, in *credentials.StaticCredentialsProvider) ([]byte, error) {
	b, e := json.Marshal(in)
	if e != nil {
		return nil, e
	}

	return b, nil
}
