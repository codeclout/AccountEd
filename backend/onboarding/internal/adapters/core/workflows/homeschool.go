package workflows

import (
	"context"

	"github.com/codeclout/AccountEd/onboarding/internal"
	"github.com/codeclout/AccountEd/onboarding/internal/ports/core/workflows"
)

func (a *Adapter) LogIn(ctx context.Context, in workflows.LogIn) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) Register(ctx context.Context, in *internal.HomeSchoolRegisterIn) (*internal.HomeSchoolRegisterOut, error) {
	var data []interface{}

	t := a.mongoActions.GetTimeStamp()

	for _, v := range in.ParentGuardians {
		v.CreatedAt = t
		v.IsActive = false
		v.IsMarkedForDeletion = false
		v.IsPending = true
		v.UpdatedAt = t

		_ = append(data, v)
	}

	_, e := a.storageOut.Register(ctx, &data)
	if e != nil {
		return nil, e
	}

	return &internal.HomeSchoolRegisterOut{ParentGuardians: []*internal.ParentGuardianOut{}}, nil
}

func (a *Adapter) UsersByAccountId(ctx context.Context) {}

func (a *Adapter) UserById(ctx context.Context) {}

func (a *Adapter) UserByUsername(ctx context.Context) {}
