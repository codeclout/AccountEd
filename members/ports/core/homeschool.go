package core

import (
	"context"

	mt "github.com/codeclout/AccountEd/members/member-types"
)

type HomeschoolCore interface {
	PreRegister(ctx context.Context) (*mt.PrimaryMemberStartRegisterOut, error)
}
