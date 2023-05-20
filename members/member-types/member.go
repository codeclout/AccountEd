package membertypes

import "github.com/google/uuid"

const (
  Female Gender = iota + 1
  Male
)

type Gender int8

type Member struct {
  AccountType         string   `json:"account_type"`
  AuthorizedRoles     []string `json:"authorized_roles"`
  County              string   `json:"county" validate:"required"`
  CreatedAt           int64    `json:"created"`
  DisplayName         string   `json:"display_name"`
  GroupID             string   `json:"group_id"`
  ID                  MemberID `json:"id"`
  Image               any      `json:"image"` // FixMe
  IsActive            bool     `json:"is_active"`
  IsMarkedForDeletion bool     `json:"is_marked_for_deletion"`
  IsPending           bool     `json:"is_pending"`
  IsVerified          bool     `json:"is_verified"`
  LegalFirstName      *string  `json:"legal_first_namefirst_name" validate:"required"`
  LegalLastName       *string  `json:"legal_last_namelast_name" validate:"required"`
  MemberType          string   `json:"member_type"`
  Pin                 *string  `json:"pin" validate:"required"`
  UpdatedAt           int64    `json:"updated"`
}

type MemberID string

type MemberGroup struct {
  CreatedAt int64  `json:"created_at"`
  ID        string `json:"id"`
  UpdatedAt int64  `json:"updated_at"`
}

type MemberSession struct{}

type MemberType struct {
  CreatedAt int64  `json:"created_at"`
  ID        string `json:"id"`
  UpdatedAt int64  `json:"updated_at"`
}

type PrimaryMemberStartRegisterIn struct {
  Username *string `json:"username" bson:"username" validate:"required,email"`
}

type PrimaryMemberStartRegisterOut struct {
  Username       *string `json:"username" bson:"username" validate:"required,email"`
  VerificationID uuid.UUID
}
