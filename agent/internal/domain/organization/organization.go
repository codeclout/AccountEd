package organization

import (
	"github.com/codeclout/AccountEd/gateway/domain"
	"github.com/google/uuid"
)

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a Adapter) CreateOrganization(in domain.CreateOrganizationInput) (*domain.CreateOrganizationResult, error) {
	return _, nil
}

func (a Adapter) CreateOrganizationUnit(in domain.CreateOrganizationUnitInput) (*domain.CreateOrganizationUnitResult, error) {
	return _, nil
}

func (a Adapter) GetOrganizationDetails(in uuid.UUID) (domain.Details, error) {
	return _, nil
}

func (a Adapter) GetOrganizationUnits(in uuid.UUID) ([]domain.OrganizationUnit, error) {
	return _, nil
}
