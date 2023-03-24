package account_types

import (
  "encoding/json"
  "fmt"

  accountTypePortCore "github.com/codeclout/AccountEd/admin/internal/ports/core/account-types"
)

type logger func(level, msg string)

type Adapter struct {
  log logger
}

func NewAdapter(l logger) *Adapter {
  return &Adapter{
    l,
  }
}

// NewAccountType - Responsible for returning the result of a new account type create.
func (a *Adapter) NewAccountType(id *string, name *string, timestamp *string) (*accountTypePortCore.NewAccountTypeOutput, error) {
  return &accountTypePortCore.NewAccountTypeOutput{
    AccountType: *name,
    CreatedAt:   *timestamp,
    ID:          *id,
    ModifiedAt:  *timestamp,
  }, nil
}

func (a *Adapter) DeleteAccountType(in *[]byte) (*accountTypePortCore.NewAccountTypeOutput, error) {
  var out accountTypePortCore.NewAccountTypeOutput

  e := json.Unmarshal(*in, &out)
  if e != nil {
    a.log("error", e.Error())
    return &out, e
  }

  return &out, nil
}

func (a *Adapter) UpdateAccountType(updatedCount *int64, id *string) (*accountTypePortCore.UpdatedAccountTypeOutput, error) {

  if *updatedCount < int64(1) {
    a.log("error", fmt.Sprintf("record id %s failed to update", *id))
    return &accountTypePortCore.UpdatedAccountTypeOutput{ID: *id, Status: false}, fmt.Errorf("record id %s failed to update", *id)
  }

  return &accountTypePortCore.UpdatedAccountTypeOutput{ID: *id, Status: true}, nil
}
