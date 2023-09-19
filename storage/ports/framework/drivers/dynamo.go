package drivers

import (
	"context"

	pb "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
)

type confirmedRegReq = pb.StoreConfirmedRegistrationRequest
type confirmedRegResp = pb.StoreConfirmedRegistrationResponse

type DynamoDbDriverPort interface {
	FetchToken(ctx context.Context, request *pb.FetchTokenRequest) (*pb.FetchTokenResponse, error)
	StorePublicToken(ctx context.Context, request *pb.TokenStoreRequest) (*pb.TokenStoreResponse, error)
	StoreConfirmedRegistration(context.Context, *confirmedRegReq) (*confirmedRegResp, error)
}
