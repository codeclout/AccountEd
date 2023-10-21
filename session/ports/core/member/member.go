package member

import (
	"context"

	dynamov1 "github.com/codeclout/AccountEd/storage/gen/dynamo/v1"
)

type SessionCoreMemberPort interface {
	ProcessTokenValidation(ctx context.Context, in *dynamov1.FetchTokenResponse) (bool, error)
}
