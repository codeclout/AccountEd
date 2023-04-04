package account_types

import (
	"context"

	"github.com/codeclout/AccountEd/onboarding/internal"
	"github.com/codeclout/AccountEd/onboarding/internal/ports/framework/out/storage"
)

type logger func(level, msg string)
type storageOut storage.AccountTypeActionPort

type Adapter struct {
	log                 logger
	frameworkOutStorage storage.AccountTypeActionPort
}

func NewAdapter(s storageOut, l logger) *Adapter {
	return &Adapter{
		log:                 l,
		frameworkOutStorage: s,
	}
}

func (a *Adapter) ListAccountTypes(ctx context.Context, limit int16) (*[]internal.AccountTypeOut, error) {

	ts, e := a.frameworkOutStorage.GetAccountTypes(ctx, limit)
	if e != nil {
		return nil, e
	}

	return ts, nil
}

func (a *Adapter) FetchAccountType(ctx context.Context, in internal.AccountTypeIn) (*internal.AccountTypeOut, error) {

	ts, e := a.frameworkOutStorage.GetAccountTypeById(ctx, in)
	if e != nil {
		return nil, e
	}

	return ts, nil
}
