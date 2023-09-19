package member

import (
	"context"

	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
	sessionTypes "github.com/codeclout/AccountEd/session/session-types"
)

type cc = context.Context

type GenerateTokenRequest = pb.GenerateTokenRequest
type GenerateTokenResponse = pb.GenerateTokenResponse
type ValidateTokenPayload = sessionTypes.ValidateTokenPayload
type ValidateTokenResponse = pb.ValidateTokenResponse

type SessionAPIMemberPort interface {
	CreateMemberToken(ctx cc, awscreds []byte, in *sessionTypes.NewTokenPayload, tch chan *GenerateTokenResponse, ech chan error)
	ValidateMemberToken(ctx cc, awscreds []byte, in *ValidateTokenPayload, tch chan *ValidateTokenResponse, ech chan error)
}
