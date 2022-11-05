package account

import (
	"encoding/json"
	"fmt"

	ports "github.com/codeclout/AccountEd/ports/core"
)

// Adapter - Core adapter for the account pkg
//
// Implements - the core accounts port
type Adapter struct {
	log func(level string, msg string)
}

// NewAdapter - Creates a new Adapter for the account pkg
func NewAdapter(logger func(level string, msg string)) *Adapter {
	return &Adapter{
		log: logger,
	}
}

// NewAccountType - Responsible for returning the result of a new account type create.
func (a *Adapter) NewAccountType(id interface{}, name string, timestamp string) (ports.NewAccountTypeOutput, error) {
	return ports.NewAccountTypeOutput{
		AccountType: name,
		CreatedAt:   timestamp,
		ID:          id,
		ModifiedAt:  timestamp,
	}, nil
}

func (a *Adapter) ListAccountTypes(accountTypes []byte) ([]ports.NewAccountTypeOutput, error) {
	var out []ports.NewAccountTypeOutput

	e := json.Unmarshal(accountTypes, &out)

	if e != nil {
		a.log("error", fmt.Sprintf("Invalid Data: %v", e))
		return []ports.NewAccountTypeOutput{}, e
	}

	return out, nil
}
