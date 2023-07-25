package api

import (
	"context"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	pb "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	"github.com/codeclout/AccountEd/storage/ports/core"
)

type Adapter struct {
	config  map[string]interface{}
	core    core.DynamoDbCorePort
	monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, core core.DynamoDbCorePort, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		core:    core,
		monitor: monitor,
	}
}

func (a *Adapter) PreRegistrationConfirmationApi(ctx context.Context, sessionID string, ch chan *pb.PreRegistrationConfirmationResponse, ech chan error) {
	a.core.PreRegistrationConfirmationCore(ctx)

}
