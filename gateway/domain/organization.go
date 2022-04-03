package domain

import (
	"time"

	"github.com/google/uuid"
)

type Organization interface {
	CreateOrganization(input CreateOrganizationInput) (*CreateOrganizationResult, error)
	CreateOrganizationUnit(input CreateOrganizationUnitInput) (*CreateOrganizationUnitResult, error)
	GetOrganizationDetails(id uuid.UUID) (Details, error)
	GetOrganizationUnits(id uuid.UUID) ([]OrganizationUnit, error)
}

type Policies interface {
	AttachPolicy(policy uuid.UUID, target uuid.UUID) (*Policy, error)
	CreatePolicy(input CreatePolicyInput) (*CreatePolicyOutput, error)
	GetPoliciesByOrganizationId(id uuid.UUID) ([]Policy, error)
}

type Team interface {
}

type Project interface {
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

type Target struct {
	ID   uuid.UUID
	Type *OrganizationUnit
}

type Policy struct {
	ID      uuid.UUID
	Targets []Target
}
