package account

import (
	"encoding/json"
	"fmt"
	"time"

	port "github.com/codeclout/AccountEd/ports/core/account-types"
	storagePort "github.com/codeclout/AccountEd/ports/framework/out/storage"
)

type l func(l string, m string)

type Adapter struct {
	accountTypeCore port.UserAccountTypeCorePort
	db              storagePort.AccountTypeActionPort
	log             l
}

func NewAdapter(act port.UserAccountTypeCorePort, db storagePort.AccountTypeActionPort, logger l) *Adapter {
	return &Adapter{
		accountTypeCore: act,
		db:              db,
		log:             logger,
	}
}

// CreateAccountType - The account_type field has a unique constraint, therefore an error might occur.
func (a *Adapter) CreateAccountType(name *string) (port.NewAccountTypeOutput, error) {
	var in port.NewAccountTypeInput
	var out port.NewAccountTypeOutput

	t := time.Unix(time.Now().Unix(), 0).Format(time.RFC3339)

	in.AccountType = *name

	payload, e := json.Marshal(in)

	if e != nil {
		a.log("error", e.Error())
		return out, e
	}

	insertId, e := a.db.InsertAccountType(payload)

	if e != nil {
		a.log("error", fmt.Sprintf("Account type creation failed: %v", e))
		return out, e
	}

	result, e := a.accountTypeCore.NewAccountType(insertId.InsertedID, *name, t)

	if e != nil {
		a.log("error", fmt.Sprintf("Core account type processing failed: %v", e))
		return out, e
	}

	return *result, nil
}

func (a *Adapter) GetAccountTypes(v int64) ([]port.NewAccountTypeOutput, error) {
	var out []port.NewAccountTypeOutput

	b, e := a.db.GetAccountTypes(v)

	if e != nil {
		a.log("error", fmt.Sprintf("Error retrieving account types: %v", e))
		return out, e
	}

	r, _ := a.accountTypeCore.ListAccountTypes(b)

	return *r, nil
}

func (a *Adapter) RemoveAccountType(id string) (port.NewAccountTypeOutput, error) {
	v, e := a.db.RemoveAccountType(id)

	if e != nil {
		a.log("error", e.Error())
		return port.NewAccountTypeOutput{}, e
	}

	p, e := a.accountTypeCore.DeleteAccountType(v)

	if e != nil {
		a.log("error", fmt.Sprintf("Error removing account type: %v", e))
		return *p, e
	}

	return *p, nil
}

func (a *Adapter) UpdateAccountType(in []byte) (port.NewAccountTypeOutput, error) {
	var out port.NewAccountTypeOutput

	_, e := a.db.UpdateAccountType(in)

	if e != nil {
		a.log("error", fmt.Sprintf("Error saving update to account type: %v", e))
		return out, e
	}

	r, _ := a.db.GetAccountTypeById(in)
	o, _ := a.accountTypeCore.UpdateAccountType(r)

	return o, nil
}

func (a *Adapter) FetchAccountType(id []byte) (port.NewAccountTypeOutput, error) {
	var out port.NewAccountTypeOutput

	b, e := a.db.GetAccountTypeById(id)

	if e != nil {
		a.log("warn", fmt.Sprintf("Unable to retrieve account type %s: %v", id, e))
		return out, e
	}

	r, _ := a.accountTypeCore.FetchAccountType(b)

	return *r, nil
}
