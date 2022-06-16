package account

import (
	"log"

	port "github.com/codeclout/AccountEd/ports/core"
	dbport "github.com/codeclout/AccountEd/ports/framework/out/db"
)

type Adapter struct {
	account port.AccountPort
	db      dbport.AccountDbPort
}

func NewAdapter(act port.AccountPort, db dbport.AccountDbPort) *Adapter {
	return &Adapter{account: act, db: db}
}

func (a Adapter) CreateAccountType(name string) (port.NewAccountTypeOutput, error) {
	insertId, ex := a.db.InsertAccountType("account_type", name)

	if ex != nil {
		log.Printf("Account type creation failed: %v", ex)
		return port.NewAccountTypeOutput{}, ex
	}

	result, e := a.account.NewAccountType(insertId.InsertedID, name)

	if e != nil {
		return port.NewAccountTypeOutput{}, e
	}

	return result, nil
}
