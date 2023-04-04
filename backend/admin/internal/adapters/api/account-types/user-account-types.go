package account_types

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/codeclout/AccountEd/admin/internal/adapters/framework/out/storage"
	accountTypeCorePort "github.com/codeclout/AccountEd/admin/internal/ports/core/account-types"
)

type coreAdapterPort accountTypeCorePort.UserAccountTypeCoreAdminPort
type logger func(level, msg string)
type storageAccountTypePort storage.AccountTypeActionPort

type Adapter struct {
	accountTypeCorePort coreAdapterPort
	log                 logger
	storage             storageAccountTypePort
}

func NewAdapter(atcp coreAdapterPort, store storageAccountTypePort, l logger) *Adapter {
	return &Adapter{
		accountTypeCorePort: atcp,
		log:                 l,
		storage:             store,
	}
}

// CreateAccountType - The account_type field has a unique constraint, therefore an error might occur.
func (a *Adapter) CreateAccountType(name *string) (*accountTypeCorePort.NewAccountTypeOutput, error) {
	var (
		in                  accountTypeCorePort.NewAccountTypeInput
		insertAccountOutput storage.MongoInsertOutput
		out                 accountTypeCorePort.NewAccountTypeOutput
	)

	in.AccountType = *name

	payload, e := json.Marshal(in)
	if e != nil {
		a.log("error", e.Error())
		return &out, e
	}

	insertResult, e := a.storage.InsertAccountType(&payload)
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

	result, e := a.accountTypeCorePort.NewAccountType(&id, name, &t)
	if e != nil {
		return &out, e
	}

	return result, nil
}

func (a *Adapter) RemoveAccountType(id *string) (*accountTypeCorePort.NewAccountTypeOutput, error) {
	var out accountTypeCorePort.NewAccountTypeOutput

	v, e := a.storage.RemoveAccountType(id)
	if e != nil {
		return &out, e
	}

	deleteAccountTypeResult, e := a.accountTypeCorePort.DeleteAccountType(v)
	if e != nil {
		return &out, e
	}

	return deleteAccountTypeResult, nil
}

func (a *Adapter) UpdateAccountType(name, id *string) (*accountTypeCorePort.UpdatedAccountTypeOutput, error) {
	var out accountTypeCorePort.UpdatedAccountTypeOutput

	count, e := a.storage.UpdateAccountType(name, id)
	if e != nil {
		a.log("error", fmt.Sprintf("Error saving update to account type: %v", e))
		return &out, e
	}

	core, e := a.accountTypeCorePort.UpdateAccountType(count, id)
	if e != nil {
		return core, e
	}

	return core, e
}
