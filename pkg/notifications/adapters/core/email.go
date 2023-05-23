package core

import "context"

type Adapter struct{}

func NewAdapter() *Adapter {}

func (a *Adapter) ProcessEmailValidation(ctx context.Context) {}
