package core

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
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

func (a *Adapter) PreRegistrationConfirmationCore(ctx context.Context) (*dynamov1.PreRegistrationConfirmationResponse, error) {
	api, ok := ctx.Value("api_input").(storageTypes.PreRegistrationSessionAPIin)
	if !ok {
		a.monitor.LogGrpcError(ctx, fmt.Sprintf("invalid core api -> PreRegistrationConfirmationCore(%v)", ctx))
		return nil, status.Error(codes.InvalidArgument, "api input is invalid")
	}

	driven := ctx.Value("driven_output").(*storageTypes.PreRegistrationSessionDrivenOut)
	if !ok {
		a.monitor.LogGrpcError(ctx, fmt.Sprintf("invalid core driven -> PreRegistrationConfirmationCore(%v)", ctx))
		return nil, status.Error(codes.InvalidArgument, "api input is invalid")
	}

	out := dynamov1.PreRegistrationConfirmationResponse{
		Active:    api.HasAutoCorrect,
		CreatedAt: timestamppb.New(driven.CreatedAt),
		ExpiresAt: driven.ExpiresAt,
		MemberId:  driven.MemberId,
	}

	return &out, nil
}
