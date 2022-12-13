package account

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/codeclout/AccountEd/adapters/framework/out/db"
	ports "github.com/codeclout/AccountEd/ports/core/account-types"
	dbport "github.com/codeclout/AccountEd/ports/framework/out/db"
)

type l func(l string, m string)

type Adapter struct {
	accountTypeCore ports.UserAccountTypeCorePort
	db              dbport.UserAccountTypeDbPort
	log             l
}

func NewAdapter(act ports.UserAccountTypeCorePort, db dbport.UserAccountTypeDbPort, logger l) *Adapter {
	return &Adapter{
		accountTypeCore: act,
		db:              db,
		log:             logger,
	}
}

// CreateAccountType - The account_type field has a unique constraint, therefore an error might occur.
func (a *Adapter) CreateAccountType(name *string) (ports.NewAccountTypeOutput, error) {
	var in ports.NewAccountTypeInput
	var out ports.NewAccountTypeOutput

	t := time.Unix(time.Now().Unix(), 0).Format(time.RFC3339)

	in.AccountType = *name

	payload, e := json.Marshal(in)

	if e != nil {
		a.log("error", e.Error())
		return out, e
	}

	did, e := a.db.InsertAccountType(payload)

	if e != nil {
		a.log("error", fmt.Sprintf("Account type creation failed: %v", e))
		return out, e.(db.WriteError)
	}

	result, e := a.accountTypeCore.NewAccountType(did.InsertedID, *name, t)

	if e != nil {
		a.log("error", fmt.Sprintf("Core account type processing failed: %v", e))
		return out, e
	}

	return *result, nil
}

func (a *Adapter) GetAccountTypes(v int64) ([]ports.NewAccountTypeOutput, error) {
	var out []ports.NewAccountTypeOutput

	b, e := a.db.GetAccountTypes(v)

	if e != nil {
		a.log("error", fmt.Sprintf("Error retrieving account types: %v", e))
		return out, e
	}

	r, _ := a.accountTypeCore.ListAccountTypes(b)

	return *r, nil
}

func (a *Adapter) RemoveAccountType(id string) (ports.NewAccountTypeOutput, error) {
	v, e := a.db.RemoveAccountType(id)

	if e != nil {
		a.log("error", e.Error())
		return ports.NewAccountTypeOutput{}, e
	}

	p, e := a.accountTypeCore.DeleteAccountType(v)

	if e != nil {
		a.log("error", fmt.Sprintf("Error removing account type: %v", e))
		return *p, e
	}

	return *p, nil
}

func (a *Adapter) UpdateAccountType(in []byte) (ports.NewAccountTypeOutput, error) {
	var out ports.NewAccountTypeOutput

	_, e := a.db.UpdateAccountType(in)

	if e != nil {
		a.log("error", fmt.Sprintf("Error saving update to account type: %v", e))
		return out, e
	}

	r, _ := a.db.GetAccountTypeById(in)
	o, _ := a.accountTypeCore.UpdateAccountType(r)

	return o, nil
}

func (a *Adapter) FetchAccountType(id []byte) (ports.NewAccountTypeOutput, error) {
	var out ports.NewAccountTypeOutput

	b, e := a.db.GetAccountTypeById(id)

	if e != nil {
		a.log("warn", fmt.Sprintf("Unable to retrieve account type %s: %v", id, e))
		return out, e
	}

	r, _ := a.accountTypeCore.FetchAccountType(b)

	return *r, nil
}
