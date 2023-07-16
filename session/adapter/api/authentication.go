package api

import (
  "context"
  "errors"

  sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

type Adapter struct{}

func NewAdapter() *Adapter {
  return &Adapter{}
}

func (a *Adapter) GetAuthenticationOptions(ctx context.Context, ch chan<- sessiontypes.AuthenticationOptions, id string) error {
  return errors.New("not implemented")
}

func (a *Adapter) GetRegistrationOptions(ctx context.Context, ch chan<- sessiontypes.RegistrationOptions, id string) error {
  return errors.New("not implemented")
}
