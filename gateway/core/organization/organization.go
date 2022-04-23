package ports

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Port interface {
	ActivateOrganization(ctx context.Context, in ActivateInput) error
	DeactivateOrganization(ctx context.Context, id uuid.UUID) error
	GetOrganization(ctx context.Context, id uuid.UUID) (Details, error)
	GetOrganizationBatch(ctx context.Context, ids []uuid.UUID) ([]Details, error)
	LogOrganizationHistoryEvent(ctx context.Context) error
	UpsertOrganizationUnit(ctx context.Context, unit Details) error
}
type ActivateInput struct {
	Date  time.Time
	Email string
	ID    uuid.UUID
	Name  string
}
type Details struct {
	ID    uuid.UUID
	Name  string
	Units []OrganizationUnit
}
type CreateOrganizationUnitInput struct {
	Name     string
	ParentID string
}

type CreateOrganizationUnitResult struct {
	ID   uuid.UUID
	Name string
}

type CreateOrganizationResult struct {
	ID uuid.UUID
}

type OrganizationUnit struct {
	ID   uuid.UUID
	Name string
}
