package member

import (
	"context"

	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
)

type SessionDriverMemberPort interface {
	GetEncryptedSessionId(ctx context.Context, request *pb.EncryptedStringRequest) (*pb.EncryptedStringResponse, error)
}
