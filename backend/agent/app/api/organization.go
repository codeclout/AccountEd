package api

import (
	"context"

	ports "github.com/codeclout/AccountEd/gateway/core/organization"
	"github.com/google/uuid"
)

type Adapter struct {
	organization ports.OrganizationPort
}

func NewAdapter(org ports.OrganizationPort) *Adapter {
	return &Adapter{
		organization: org,
	}
}

func (a Adapter) GetOrganizationByID(ctx context.Context, id uuid.UUID) (ports.Details, error) {
	org, e := a.organization.GetOrganization(ctx, id)

	if e != nil {
		return ports.Details{}, e
	}
	return org, nil
}

func (a Adapter) GetOrganizationListByID(ctx context.Context, ids []uuid.UUID) ([]ports.Details, error) {
	orgs, e := a.organization.GetOrganizationBatch(ctx, ids)
	d := make([]ports.Details, len(ids))

	if e != nil {
		return d, e
	}
	return orgs, nil
}

func (a Adapter) PostOrganizationUnit(ctx context.Context, unit ports.Details) error {
	e := a.organization.UpsertOrganizationUnit(ctx, unit)

	if e != nil {
		return e
	}

	return nil
}

func (a Adapter) PublishEvent(ctx context.Context) error {
	e := a.organization.LogOrganizationHistoryEvent(ctx)

	if e != nil {
		return e
	}

	return nil
}

func (a Adapter) PutActivation(ctx context.Context, in ports.ActivateInput) error {
	e := a.organization.ActivateOrganization(ctx, in)

	if e != nil {
		return e
	}

	return nil
}

func (a Adapter) PutDeactivation(ctx context.Context, id uuid.UUID) error {
	e := a.organization.DeactivateOrganization(ctx, id)

	if e != nil {
		return e
	}

	return nil
}
