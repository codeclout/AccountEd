package ports

import (
	ports "github.com/codeclout/AccountEd/gateway/core/organization"
	"github.com/google/uuid"
)

type OrganizationAPIPort interface {
	GetOrganizationDetails(id uuid.UUID) (ports.Details, error)
	GetOrganizationUnit(id uuid.UUID) (ports.OrganizationUnit, error)
}
