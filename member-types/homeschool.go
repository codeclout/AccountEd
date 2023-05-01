package memberTypes

import "time"

type Gender int8

const (
	Female Gender = iota
	Male
)

type Member struct {
	AccountType         string   `json:"account_type"`
	AuthorizedRoles     []string `json:"authorized_roles"`
	Category            string   `json:"category"`
	County              string   `json:"county" validate:"required"`
	CreatedAt           int64    `json:"created_at"`
	DisplayName         string   `json:"display_name"`
	FirstName           string   `json:"first_name" validate:"required"`
	GroupId             string   `json:"group_id"`
	Id                  string
	IsActive            bool   `json:"is_active"`
	IsMarkedForDeletion bool   `json:"is_marked_for_deletion"`
	IsPending           bool   `json:"is_pending"`
	LastName            string `json:"last_name" validate:"required"`
	Pin                 string `json:"pin" validate:"required"`
	UpdatedAt           int64  `json:"updated_at"`
}

type ParentGuardian struct {
	Member
	HomeSchoolRegistrationId string `json:"home_school_registration_id" bson:"home_school_registration_id" validate:"required"`
	Phone                    string `json:"phone" bson:"phone" validate:"required"`
	Username                 string `json:"username" bson:"username" validate:"required,email"`
	Zipcode                  string `json:"zipcode" bson:"zipcode"`
}

type Student struct {
	Member
	BirthCertificateGender Gender    `json:"birth_certificate_gender"`
	DOB                    time.Time `json:"dob"`
	Email                  string    `json:"email"`
	GradeTypeRequested     uint8     `json:"grade_type_requested" validate:"required"`
	PrincipalId            string    `json:"principal_id"`
}

type HomeSchoolRegisterIn struct {
	ParentGuardians []*ParentGuardian `json:"parent_guardians"`
	Students        []*Student        `json:"students"`
}

type HomeSchoolRegisterOut struct {
	ParentGuardians []*ParentGuardianOut `json:"parent_guardians"`
}

type ParentGuardianOut struct {
	GroupId  string `json:"account_id" validate:"required"`
	Id       string `json:"user_id" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type StudentOut struct {
	ParentAccountId string `json:"parent_account_id" validate:"required"`
}

