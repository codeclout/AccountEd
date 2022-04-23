package api

import (
	"context"

	ports "github.com/codeclout/AccountEd/gateway/core/organization"
	"github.com/google/uuid"
)

type Adapter struct{}

// NewAdapter - creates a new Adapter and returns a pointer to an Adapter struct literal
func NewAdapter() *Adapter {
	return &Adapter{}
}

// ActivateOrganization - Activates an existing organization from an inactive or pending state
func (a Adapter) ActivateOrganization(ctx context.Context, in ports.ActivateInput) error {
	return nil
}

// DeactivateOrganization - Deactivates an existing organization from an active or pending state
func (a Adapter) DeactivateOrganization(ctx context.Context, id uuid.UUID) error {
	return nil
}

// GetOrganization - Get an organizations' details by its' identifier
func (a Adapter) GetOrganization(ctx context.Context, id uuid.UUID) (ports.Details, error) {
	return ports.Details{}, nil
}

// GetOrganizationBatch - Get a slice of organization details via their identifiers
func (a Adapter) GetOrganizationBatch(ctx context.Context, ids []uuid.UUID) ([]ports.Details, error) {
	d := make([]ports.Details, 0, 100)

	return d, nil
}

// LogOrganizationHistoryEvent - An asynchronous organization event logger
func (a Adapter) LogOrganizationHistoryEvent(ctx context.Context) error {
	return nil
}

// UpsertOrganizationUnit - Insert or Update organization details
func (a Adapter) UpsertOrganizationUnit(ctx context.Context, unit ports.Details) error {
	return nil
}
