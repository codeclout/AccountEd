package server

import (
  "context"

  membertypes "github.com/codeclout/AccountEd/members/member-types"
)

type Authorization interface {
  HandleAuthorizationOptions(ctx context.Context, id membertypes.MemberID)
}
