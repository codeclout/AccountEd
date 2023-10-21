package member

import (
	"context"

	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
	sessionTypes "github.com/codeclout/AccountEd/session/session-types"
)

type cc = context.Context

type ValidateTokenPayload = sessionTypes.ValidateTokenPayload
type ValidateTokenResponse = pb.ValidateTokenResponse

type SessionAPIMemberPort interface {
	ValidateMemberToken(ctx cc, awscreds []byte, in *ValidateTokenPayload, tch chan *ValidateTokenResponse, ech chan error)
}
