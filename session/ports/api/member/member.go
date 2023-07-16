package member

import (
	"context"

	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
)

type SessionAPIMemberPort interface {
	EncryptSessionId(ctx context.Context, awscredentials []byte, id string, uch chan *pb.EncryptedStringResponse, echan chan error)
}
