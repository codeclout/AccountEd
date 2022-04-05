package api

import "github.com/google/uuid"

type OrganizationAPIPort interface {
	GetOrganizationDetails(id uuid.UUID) (Details, error)
	GetOrganizationUnit(id uuid.UUID) (OrganizationUnit, error)
}

type Details struct {
	ID    uuid.UUID
	Name  string
	Units []OrganizationUnit
}

type OrganizationUnit struct {
	ID   uuid.UUID
	Name string
}
