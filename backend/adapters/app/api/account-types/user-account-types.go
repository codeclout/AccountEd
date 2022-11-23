package account

import (
	"encoding/json"
	"fmt"
	"time"

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
	return &Adapter{accountTypeCore: act, db: db, log: logger}
}

// CreateAccountType - The account_type field has a unique constraint, therefore an error might occur.
func (a *Adapter) CreateAccountType(name string) (ports.NewAccountTypeOutput, error) {
	var in ports.NewAccountTypeInput
	t := time.Unix(time.Now().Unix(), 0).Format(time.RFC3339)

	in.AccountType = name

	payload, e := json.Marshal(in)

	did, ex := a.db.InsertAccountType(payload)

	if ex != nil {
		a.log("error", fmt.Sprintf("Account type creation failed: %v", ex))
		return ports.NewAccountTypeOutput{}, ex
	}

	result, e := a.accountTypeCore.NewAccountType(did.InsertedID, name, t)

	if e != nil {
		a.log("error", fmt.Sprintf("Core account type processing failed: %v", e))
		return ports.NewAccountTypeOutput{}, e
	}

	return result, nil
}

func (a *Adapter) GetAccountTypes(v int64) ([]ports.NewAccountTypeOutput, error) {
	b, e := a.db.GetAccountTypes(v)

	if e != nil {
		a.log("error", fmt.Sprintf("Error retrieving account types: %v", e))
		return []ports.NewAccountTypeOutput{}, e
	}

	return a.accountTypeCore.ListAccountTypes(b)
}

func (a *Adapter) RemoveAccountType(id string) (ports.NewAccountTypeOutput, error) {
	v, e := a.db.RemoveAccountType(id)
	p, e := a.accountTypeCore.DeleteAccountType(v)

	if e != nil {
		a.log("error", fmt.Sprintf("Error removing account type: %v", e))
		return p, e
	}

	return p, nil
}

func (a *Adapter) UpdateAccountType(in []byte) (ports.NewAccountTypeOutput, error) {
	_, e := a.db.UpdateAccountType(in)

	if e != nil {
		a.log("error", fmt.Sprintf("Error saving update to account type: %v", e))
		return ports.NewAccountTypeOutput{}, e
	}

	r, e := a.db.GetAccountTypeById(in)
	o, e := a.accountTypeCore.UpdateAccountType(r)

	fmt.Println(r)

	return o, nil
}
