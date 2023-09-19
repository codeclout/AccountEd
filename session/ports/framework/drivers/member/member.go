package member

import (
	"context"

	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
)

type SessionDriverMemberPort interface {
	GenerateMemberToken(ctx context.Context, request *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error)
	ValidateMemberToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error)
}
