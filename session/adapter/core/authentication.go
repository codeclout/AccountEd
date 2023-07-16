package core

import (
  "context"
  "errors"

  sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
  return &Adapter{}
}

func (a *Adapter) AuthenticationOptions(ctx context.Context, id string) (*sessiontypes.AuthenticationOptions, error) {
  return nil, errors.New("not implemented")
}

func (a *Adapter) RegistrationOptions(ctx context.Context, id string) (*sessiontypes.RegistrationOptions, error) {
  return nil, errors.New("not implemented")
}
