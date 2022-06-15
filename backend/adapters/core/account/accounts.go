package account

import (
	ports "github.com/codeclout/AccountEd/ports/core"
)

// Adapter - Core adapter for the account pkg
//
// Implements - the core accounts port
type Adapter struct {
}

// NewAdapter - Creates a new Adapter for the account pkg
func NewAdapter() *Adapter {
	return &Adapter{}
}

// NewAccountType - Responsible for returning the result of a new account type create.
// The account_type field has a unique constraint, therefore an error might occur.
func (a Adapter) NewAccountType(insertId interface{}, name string) (ports.NewAccountTypeOutput, error) {
	return ports.NewAccountTypeOutput{
		AccountType: name,
		ID:          insertId,
	}, nil
}
