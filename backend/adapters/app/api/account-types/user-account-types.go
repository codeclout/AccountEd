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
	account ports.UserAccountTypeCorePort
	db      dbport.UserAccountTypeDbPort
	log     l
}

func NewAdapter(act ports.UserAccountTypeCorePort, db dbport.UserAccountTypeDbPort, logger l) *Adapter {
	return &Adapter{account: act, db: db, log: logger}
}

// CreateAccountType - The account_type field has a unique constraint, therefore an error might occur.
func (a *Adapter) CreateAccountType(name string) (ports.NewAccountTypeOutput, error) {
	var in ports.NewAccountTypeInput
	t := time.Unix(time.Now().Unix(), 0).Format(time.RFC3339)

	in.AccountType = name
	in.CreatedAt = t
	in.ModifiedAt = t

	payload, e := json.Marshal(in)

	did, ex := a.db.InsertAccountType(payload)

	if ex != nil {
		a.log("error", fmt.Sprintf("Account type creation failed: %v", ex))
		return ports.NewAccountTypeOutput{}, ex
	}

	result, e := a.account.NewAccountType(did.InsertedID, name, t)

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

	return a.account.ListAccountTypes(b)
}
