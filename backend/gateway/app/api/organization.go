package ports

import (
	"context"

	model "github.com/codeclout/AccountEd/backend/gateway/core/organization"
	"github.com/google/uuid"
)

type OrganizationAPI interface {
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (model.Details, error)
	GetOrganizationListByID(ctx context.Context, ids []uuid.UUID) ([]model.Details, error)
	PostOrganizationUnit(ctx context.Context, unit model.Details) error
	PublishEvent(ctx context.Context) error
	PutActivation(ctx context.Context, in model.ActivateInput) error
	PutDeactivation(ctx context.Context, id uuid.UUID) error
}
