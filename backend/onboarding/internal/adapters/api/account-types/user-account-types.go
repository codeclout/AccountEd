package account

import (
	"context"

	"github.com/codeclout/AccountEd/onboarding/internal"
	accountTypeCorePort "github.com/codeclout/AccountEd/onboarding/internal/ports/core/account-types"
)

type coreAdapterPort accountTypeCorePort.UserAccountTypeCorePort
type logger func(l string, m string)

type Adapter struct {
	accountTypeCore coreAdapterPort
	log             logger
}

func NewAdapter(act coreAdapterPort, l logger) *Adapter {
	return &Adapter{
		accountTypeCore: act,
		log:             l,
	}
}

func (a *Adapter) GetAccountTypes(ctx context.Context, limit int16, ch chan *[]internal.AccountTypeOut) error {

	result, e := a.accountTypeCore.ListAccountTypes(ctx, limit)
	if e != nil {
		return e
	}

	ch <- result

	return nil
}

func (a *Adapter) FetchAccountType(ctx context.Context, in internal.AccountTypeIn, ch chan *internal.AccountTypeOut) error {

	result, e := a.accountTypeCore.FetchAccountType(ctx, in)
	if e != nil {
		return e
	}

	ch <- result
	return nil
}
