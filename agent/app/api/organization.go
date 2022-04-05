package api

import (
	"github.com/codeclout/AccountEd/gateway/domain/organization"
	"github.com/google/uuid"
)

type Adapter struct {
	org organization.Port
}

func (a Adapter) GetOrganizationDetails(id uuid.UUID) (organization.OrgDetails, error) {
	return a.org.Organization(id)
}

func (a Adapter) GetOrganizationUnits(id uuid.UUID) (organization.OrgUnit, error) {
	return a.org.OrganizationUnit(id)
}
