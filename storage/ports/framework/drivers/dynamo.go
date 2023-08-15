package drivers

import (
	"context"

	pb "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
)

type confirmedRegReq = pb.StoreConfirmedRegistrationRequest
type confirmedRegResp = pb.StoreConfirmedRegistrationResponse
type preConfirmRegReq = pb.PreRegistrationConfirmationRequest
type preConfirmRegResp = pb.PreRegistrationConfirmationResponse

type DynamoDbDriverPort interface {
	StorePreConfirmationRegistrationSession(ctx context.Context, request *preConfirmRegReq) (*preConfirmRegResp, error)
	StoreConfirmedRegistration(context.Context, *confirmedRegReq) (*confirmedRegResp, error)
}
