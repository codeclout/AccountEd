package member

import (
	"context"

	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

type sessionIdResponse = pb.EncryptedStringResponse
type storeMetadata = sessiontypes.SessionStoreMetadata

type SessionAPIMemberPort interface {
	EncryptSessionId(ctx context.Context, awscreds []byte, in *storeMetadata, uch chan *sessionIdResponse, ech chan error)
}
