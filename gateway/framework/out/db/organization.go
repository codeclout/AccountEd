package ports

import (
	"context"

	model "github.com/codeclout/AccountEd/gateway/core/organization"
	"github.com/google/uuid"
)

type OrganizationRepository interface {
	ActivateOrganization(ctx context.Context, id uuid.UUID) error
	DeactivateOrganization(ctx context.Context, id uuid.UUID) error
	GetOrganization(ctx context.Context, id uuid.UUID) (model.Details, error)
	GetOrganizationBatch(ctx context.Context, ids []uuid.UUID) ([]model.Details, error)
	LogOrganizationHistoryEvent(ctx context.Context, event model.OrganizationEvent) error
	UpsertOrganizationUnit(ctx context.Context, unit model.Details) error
}
