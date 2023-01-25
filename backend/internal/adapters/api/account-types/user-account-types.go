package account

import (
	"encoding/json"
	"fmt"
	"strconv"

	port "github.com/codeclout/AccountEd/internal/ports/core/account-types"
	storagePort "github.com/codeclout/AccountEd/internal/ports/framework/out/storage"
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
func (a *Adapter) CreateAccountType(name *string) (*port.NewAccountTypeOutput, error) {
	var (
		in                  port.NewAccountTypeInput
		insertAccountOutput storagePort.InsertAccountOutput
		out                 port.NewAccountTypeOutput
	)

	in.AccountType = *name

	payload, e := json.Marshal(in)
	if e != nil {
		a.log("error", e.Error())
		return &out, e
	}

	insertResult, e := a.db.InsertAccountType(&payload)
	if e != nil {
		return &out, e
	}

	e = json.Unmarshal(*insertResult, &insertAccountOutput)
	if e != nil {
		a.log("error", fmt.Sprintf("Account type insert, output operation failed: %v", e))
		return &out, e
	}

	id := string(*insertAccountOutput.InsertId)
	t := strconv.FormatInt(int64(*insertAccountOutput.TimeStamp), 10)

	result, e := a.accountTypeCore.NewAccountType(&id, name, &t)
	if e != nil {
		return &out, e
	}

	return result, nil
}

func (a *Adapter) GetAccountTypes(limit *int16) (*[]port.NewAccountTypeOutput, error) {

	getAccountTypeListResult, e := a.db.GetAccountTypes(limit)
	if e != nil {
		return nil, e
	}

	result, e := a.accountTypeCore.ListAccountTypes(getAccountTypeListResult)
	if e != nil {
		return nil, e
	}

	return result, nil
}

func (a *Adapter) RemoveAccountType(id *string) (*port.NewAccountTypeOutput, error) {
	var out port.NewAccountTypeOutput

	v, e := a.db.RemoveAccountType(id)
	if e != nil {
		return &out, e
	}

	deleteAccountTypeResult, e := a.accountTypeCore.DeleteAccountType(v)
	if e != nil {
		return &out, e
	}

	return deleteAccountTypeResult, nil
}

func (a *Adapter) UpdateAccountType(name, id *string) (*port.UpdatedAccountTypeOutput, error) {
	var out port.UpdatedAccountTypeOutput

	count, e := a.db.UpdateAccountType(name, id)
	if e != nil {
		a.log("error", fmt.Sprintf("Error saving update to account type: %v", e))
		return &out, e
	}

	core, e := a.accountTypeCore.UpdateAccountType(count, id)
	if e != nil {
		return core, e
	}

	return core, e
}

func (a *Adapter) FetchAccountType(id *string) (*port.NewAccountTypeOutput, error) {
	var out port.NewAccountTypeOutput

	b, e := a.db.GetAccountTypeById(id)

	if e != nil {
		a.log("warn", fmt.Sprintf("Unable to retrieve account type %s: %v", *id, e))
		return &out, e
	}

	r, _ := a.accountTypeCore.FetchAccountType(b)

	return r, nil
}
