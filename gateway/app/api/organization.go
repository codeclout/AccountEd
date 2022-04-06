package ports

import (
	"github.com/codeclout/AccountEd/agent/core/organization"
	"github.com/codeclout/AccountEd/gateway/core/organization"
	"github.com/google/uuid"
)

type OrganizationAPIPort interface {
	GetOrganizationDetails(id uuid.UUID) (organization.Details, error)
	GetOrganizationUnit(id uuid.UUID) (organization.OrganizationUnit, error)
}
