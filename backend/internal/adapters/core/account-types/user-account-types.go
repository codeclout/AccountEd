package account

import (
	"encoding/json"
	"errors"
	"fmt"

	ports "github.com/codeclout/AccountEd/internal/ports/core/account-types"
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
func (a *Adapter) NewAccountType(id *string, name *string, timestamp *string) (*ports.NewAccountTypeOutput, error) {
	return &ports.NewAccountTypeOutput{
		AccountType: *name,
		CreatedAt:   *timestamp,
		ID:          *id,
		ModifiedAt:  *timestamp,
	}, nil
}

func (a *Adapter) ListAccountTypes(accountTypes *[]byte) (*[]ports.NewAccountTypeOutput, error) {
	var out []ports.NewAccountTypeOutput

	e := json.Unmarshal(*accountTypes, &out)
	if e != nil {
		a.log("error", fmt.Sprintf("core - invalid account type list: %v", e))
		return nil, e
	}

	if len(out) == 0 {
		a.log("error", "0 account types exist")
		return nil, errors.New("0 account types are in the system")
	}

	return &out, nil
}

func (a *Adapter) DeleteAccountType(in *[]byte) (*ports.NewAccountTypeOutput, error) {
	var out ports.NewAccountTypeOutput

	e := json.Unmarshal(*in, &out)
	if e != nil {
		a.log("error", e.Error())
		return &out, e
	}

	return &out, nil
}

func (a *Adapter) UpdateAccountType(updatedCount *int64, id *string) (*ports.UpdatedAccountTypeOutput, error) {

	if *updatedCount < int64(1) {
		a.log("error", fmt.Sprintf("record id %s failed to update", *id))
		return &ports.UpdatedAccountTypeOutput{ID: *id, Status: false}, fmt.Errorf("record id %s failed to update", *id)
	}

	return &ports.UpdatedAccountTypeOutput{ID: *id, Status: true}, nil
}

func (a *Adapter) FetchAccountType(in *[]byte) (*ports.NewAccountTypeOutput, error) {
	var out ports.NewAccountTypeOutput

	e := json.Unmarshal(*in, &out)

	if e != nil {
		a.log("error", e.Error())
		return &out, e
	}

	return &out, nil
}
