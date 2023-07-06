package member

import (
	"context"

	pb "github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
)

type SessionAPIMemberPort interface {
	EncryptSessionId(ctx context.Context, awscredentials []byte, id string, uch chan *pb.EncryptedStringResponse, echan chan error)
}
