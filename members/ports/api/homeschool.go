package api

import (
  "context"

  mt "github.com/codeclout/AccountEd/members/member-types"
)

type HomeschoolAPI interface {
  PreRegisterPrimaryMember(ctx context.Context, data *mt.PrimaryMemberStartRegisterIn, ch chan *mt.PrimaryMemberStartRegisterOut, ech chan error)
}
