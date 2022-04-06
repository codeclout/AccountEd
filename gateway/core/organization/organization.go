package ports

import (
	"time"

	"github.com/google/uuid"
)

type Port interface {
	CreateOrganization(input CreateOrganizationInput) (*CreateOrganizationResult, error)
	CreateOrganizationUnit(input CreateOrganizationUnitInput) (*CreateOrganizationUnitResult, error)
	Organization(input uuid.UUID) (Details, error)
	OrganizationUnit(input uuid.UUID) (OrganizationUnit, error)
}

type PoliciesPort interface {
	CreatePolicy(input CreatePolicyInput) (*CreatePolicyOutput, error)
}

type CreatePolicyInput struct {
	Content     string
	Description string
	Name        string
	Type        string
}

type CreatePolicyOutput struct {
	ID   uuid.UUID
	Name string
}

type CreateOrganizationInput struct {
	Date  time.Time
	Email string
	Name  string
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

type Details struct {
	ID    uuid.UUID
	Name  string
	Units []OrganizationUnit
}

type OrganizationUnit struct {
	ID   uuid.UUID
	Name string
}
