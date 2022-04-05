package api

import "github.com/google/uuid"

type PolicyAPI interface {
	AttachPolicy(policy uuid.UUID, target uuid.UUID) (*Policy, error)
	GetPoliciesByOrganizationId(id uuid.UUID) ([]Policy, error)
}

type Target struct {
	ID   uuid.UUID
	Type *OrganizationUnit
}

type Policy struct {
	ID      uuid.UUID
	Targets []Target
}
