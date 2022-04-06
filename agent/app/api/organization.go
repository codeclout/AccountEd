package api

import (
	ports "github.com/codeclout/AccountEd/gateway/core/organization"
	"github.com/google/uuid"
)

type Adapter struct {
	org ports.Port
}

func NewAdapter(a ports.Port) *Adapter {
	return &Adapter{org: a}
}

func (a Adapter) GetOrganizationDetails(id uuid.UUID) (ports.Details, error) {
	return a.org.Organization(id)
}

func (a Adapter) GetOrganizationUnit(id uuid.UUID) (ports.OrganizationUnit, error) {
	return a.org.OrganizationUnit(id)
}
