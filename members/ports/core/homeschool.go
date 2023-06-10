package core

import (
	"context"

	mt "github.com/codeclout/AccountEd/members/member-types"
)

type HomeschoolCore interface {
	PreRegister(ctx context.Context, in mt.EmailValidationIn) (*mt.PrimaryMemberStartRegisterOut, error)
}
