package driver

import (
	"context"
	"time"
)

type SessionPort interface {
	CreateToken(ctx context.Context, duration time.Duration, groupid, memberid string) (*string, error)
	// VerifyTokenData(ctx context.Context) (*sessiontypes.StatelessAPI, error)
}
