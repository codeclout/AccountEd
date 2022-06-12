package ports

import (
	"context"

	model "github.com/codeclout/AccountEd/gateway/core/organization"
	"github.com/google/uuid"
)

type OrganizationRepository interface {
	ActivateOrganization(id uuid.UUID) error
	DeactivateOrganization(id uuid.UUID) error
	GetOrganization(id uuid.UUID) (model.Details, error)
	GetOrganizationBatch(ids []uuid.UUID) ([]model.Details, error)
	LogOrganizationHistoryEvent(ctx context.Context) error
	UpsertOrganizationUnit(unit model.Details) error
}
