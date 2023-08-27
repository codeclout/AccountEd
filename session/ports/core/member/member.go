package member

import (
	"context"

	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

type SessionCoreMemberPort interface {
	ProcessSessionIdEncryption(ctx context.Context, in *string, id string) (*sessiontypes.SessionIdEncryptionOut, error)
	ProcessSessionIdDecryption(associatedData, key []byte, cipherIn *string) ([]byte, error)
}
