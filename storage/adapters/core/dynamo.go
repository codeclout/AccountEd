package core

import (
	"context"
	"reflect"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

type TokenStoreResult = storageTypes.TokenStoreResult

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

func (a *Adapter) ProcessFetchedToken(ctx context.Context, in storageTypes.FetchTokenResult) (*dynamov1.FetchTokenResponse, error) {
	if reflect.DeepEqual(in, storageTypes.FetchTokenResult{}) {
		const msg = "received zero state decrypted message"

		a.monitor.LogGrpcError(ctx, msg)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	out := dynamov1.FetchTokenResponse{
		Active:    in.Active,
		MemberId:  in.MemberID,
		PublicKey: in.PublicKey,
		Token:     in.Token,
		TokenId:   in.TokenId,
	}

	return &out, nil
}

func (a *Adapter) ProcessStoredToken(ctx context.Context, in *TokenStoreResult) (*dynamov1.TokenStoreResponse, error) {
	if reflect.DeepEqual(in, storageTypes.TokenStoreResult{}) {
		const msg = "received zero state encrypted session"
		a.monitor.LogGrpcError(ctx, msg)
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	out := dynamov1.TokenStoreResponse{
		Active:    in.Active,
		CreatedAt: timestamppb.New(in.CreatedAt),
		ExpiresAt: timestamppb.New(in.ExpiresAt),
		Token:     in.Token,
	}

	return &out, nil
}
