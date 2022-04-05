package organization

import (
	"github.com/codeclout/AccountEd/gateway/domain/organization"
	"github.com/google/uuid"
)

type Adapter struct{}

// NewAdapter - creates a new Adapter and returns a pointer to an Adapter struct literal
func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a Adapter) CreateOrganization(in organization.CreateOrganizationInput) (*organization.CreateOrganizationResult, error) {
	return _, nil
}

func (a Adapter) CreateOrganizationUnit(in organization.CreateOrganizationUnitInput) (*organization.CreateOrganizationUnitResult, error) {
	return _, nil
}

func (a Adapter) Organization(in uuid.UUID) (organization.OrgDetails, error) {
	return _, nil
}

func (a Adapter) OrganizationUnit(in uuid.UUID) (organization.OrgUnit, error) {
	return _, nil
}
