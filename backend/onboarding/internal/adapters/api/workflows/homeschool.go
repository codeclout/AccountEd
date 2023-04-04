package workflows

import (
	"context"

	"github.com/codeclout/AccountEd/onboarding/internal"
	"github.com/codeclout/AccountEd/onboarding/internal/ports/core/workflows"
)

type homeSchoolWorkflowPort workflows.HomeSchoolCorePort
type logger func(l, m string)

type Adapter struct {
	homeSchoolWorkflow homeSchoolWorkflowPort
	log                logger
}

func NewAdapter(workflow homeSchoolWorkflowPort, l logger) *Adapter {
	return &Adapter{
		homeSchoolWorkflow: workflow,
		log:                l,
	}
}

func (a *Adapter) LoginHomeSchool(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) RegisterHomeSchool(ctx context.Context, out chan internal.HomeSchoolRegisterOut, in internal.HomeSchoolRegisterIn) error {
	data, e := a.homeSchoolWorkflow.Register(ctx, &in)
	if e != nil {
		return e
	}

	out <- *data
	return nil
}
