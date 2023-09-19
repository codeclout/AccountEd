package member

import (
	"context"

	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
)

type SessionCoreMemberPort interface {
	ProcessTokenCreation(ctx context.Context, in *sessiontypes.TokenPayload) (*sessiontypes.TokenCreateOut, error)
	ProcessTokenValidation(ctx context.Context, in *dynamov1.FetchTokenResponse) (bool, error)
}
