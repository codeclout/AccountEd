package adapters

import (
	ports "github.com/codeclout/AccountEd/gateway/core/organization"
	"github.com/google/uuid"
)

type Adapter struct{}

// NewAdapter - creates a new Adapter and returns a pointer to an Adapter struct literal
func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a Adapter) CreateOrganization(in ports.CreateOrganizationInput) (*ports.CreateOrganizationResult, error) {
	return _, nil
}

func (a Adapter) CreateOrganizationUnit(in ports.CreateOrganizationUnitInput) (*ports.CreateOrganizationUnitResult, error) {
	return _, nil
}

func (a Adapter) Organization(in uuid.UUID) (ports.Details, error) {
	return _, nil
}

func (a Adapter) OrganizationUnit(in uuid.UUID) (ports.OrganizationUnit, error) {
	return _, nil
}
