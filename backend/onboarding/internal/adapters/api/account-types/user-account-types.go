package account

import (
	"fmt"

	accountTypeCorePort "github.com/codeclout/AccountEd/onboarding/internal/ports/core/account-types"
	storageFrameworkPort "github.com/codeclout/AccountEd/onboarding/internal/ports/framework/out/storage"
)

type coreAdapterPort accountTypeCorePort.UserAccountTypeCorePort
type logger func(l string, m string)
type storageAccountTypePort storageFrameworkPort.AccountTypeActionPort

type Adapter struct {
	accountTypeCore coreAdapterPort
	storage         storageAccountTypePort
	log             logger
}

func NewAdapter(act coreAdapterPort, store storageAccountTypePort, l logger) *Adapter {
	return &Adapter{
		accountTypeCore: act,
		storage:         store,
		log:             l,
	}
}

func (a *Adapter) GetAccountTypes(limit *int16) (*[]accountTypeCorePort.NewAccountTypeOutput, error) {

	getAccountTypeListResult, e := a.storage.GetAccountTypes(limit)
	if e != nil {
		return nil, e
	}

	result, e := a.accountTypeCore.ListAccountTypes(getAccountTypeListResult)
	if e != nil {
		return nil, e
	}

	return result, nil
}

func (a *Adapter) FetchAccountType(id *string) (*accountTypeCorePort.NewAccountTypeOutput, error) {
	var out accountTypeCorePort.NewAccountTypeOutput

	b, e := a.storage.GetAccountTypeById(id)

	if e != nil {
		a.log("warn", fmt.Sprintf("Unable to retrieve account type %s: %v", *id, e))
		return &out, e
	}

	r, _ := a.accountTypeCore.FetchAccountType(b)

	return r, nil
}
