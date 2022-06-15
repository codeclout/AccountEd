package ports

import ports "github.com/codeclout/AccountEd/ports/core"

type AccountAPIPort interface {
	CreateAccountType(in string) (ports.NewAccountTypeOutput, error)
}
