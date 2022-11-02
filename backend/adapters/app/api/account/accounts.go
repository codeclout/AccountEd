package account

import (
	"encoding/json"
	"log"
	"time"

	ports "github.com/codeclout/AccountEd/ports/core"
	dbport "github.com/codeclout/AccountEd/ports/framework/out/db"
)

type Adapter struct {
	account ports.AccountPort
	db      dbport.AccountDbPort
}

func NewAdapter(act ports.AccountPort, db dbport.AccountDbPort) *Adapter {
	return &Adapter{account: act, db: db}
}

// CreateAccountType - The account_type field has a unique constraint, therefore an error might occur.
func (a *Adapter) CreateAccountType(name string) (ports.NewAccountTypeOutput, error) {
	var in ports.NewAccountTypeInput
	t := time.Unix(time.Now().Unix(), 0).Format(time.RFC3339)

	in.AccountType = name
	in.CreatedAt = t
	in.ModifiedAt = t

	payload, e := json.Marshal(in)

	did, ex := a.db.InsertAccountType("account_type", payload)

	if ex != nil {
		log.Printf("Account type creation failed: %v", ex)
		return ports.NewAccountTypeOutput{}, ex
	}

	result, e := a.account.NewAccountType(did.InsertedID, name, t)

	if e != nil {
		return ports.NewAccountTypeOutput{}, e
	}

	return result, nil
}
