package member

import (
	"context"

	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
)

type SessionDriverMemberPort interface {
	ValidateMemberToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error)
}
