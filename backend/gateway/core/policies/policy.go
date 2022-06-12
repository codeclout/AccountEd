package ports

import "github.com/google/uuid"

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
