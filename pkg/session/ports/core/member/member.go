package member

import (
	"context"

	sessiontypes "github.com/codeclout/AccountEd/pkg/session/session-types"
)

type SessionCoreMemberPort interface {
	ProcessSessionIdEncryption(ctx context.Context, id, key string) (*sessiontypes.SessionIdEncryptionOut, error)
	ProcessSessionIdDecryption(associatedData, key []byte, cipherIn *string) ([]byte, error)
}
