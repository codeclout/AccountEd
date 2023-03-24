package account_types

import (
	"encoding/json"
	"errors"
	"fmt"

	accountTypePortCore "github.com/codeclout/AccountEd/onboarding/internal/ports/core/account-types"
)

type Adapter struct {
	log func(level string, msg string)
}

// NewAdapter - Creates a new Adapter for the account pkg
func NewAdapter(logger func(level string, msg string)) *Adapter {
	return &Adapter{
		log: logger,
	}
}

func (a *Adapter) ListAccountTypes(accountTypes *[]byte) (*[]accountTypePortCore.NewAccountTypeOutput, error) {
	var out []accountTypePortCore.NewAccountTypeOutput

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

func (a *Adapter) FetchAccountType(in *[]byte) (*accountTypePortCore.NewAccountTypeOutput, error) {
	var out accountTypePortCore.NewAccountTypeOutput

	e := json.Unmarshal(*in, &out)

	if e != nil {
		a.log("error", e.Error())
		return &out, e
	}

	return &out, nil
}
