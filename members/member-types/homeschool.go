package membertypes

import "time"

type ParentGuardian struct {
	Member
	Phone    string `json:"phone"`
	Username string `json:"username" bson:"username" validate:"required,email"`
}

type ParentGuardianCompliance struct {
	HomeSchoolRegistrationID string `json:"home_school_registration_id" bson:"home_school_registration_id" validate:"required"`
	Phone                    string `json:"phone" bson:"phone" validate:"required"`
	Zipcode                  string `json:"zipcode" bson:"zipcode"`
}

type Student struct {
	Member
	Email              string `json:"email"`
	GradeTypeRequested uint8  `json:"grade_type_requested" validate:"required"`
	HasSpecialNeeds    bool   `json:"has_special_needs"`
	PrincipalID        string `json:"principal_id"`
}

type StudentCompliance struct {
	BirthCertificateGender Gender    `json:"birth_certificate_gender"`
	DOB                    time.Time `json:"dob"`
}

type HomeSchoolRegisterIn struct {
	ParentGuardians []*ParentGuardian `json:"parent_guardians"`
	Students        []*Student        `json:"students"`
}

type HomeSchoolRegisterOut struct {
	ParentGuardians []*ParentGuardianOut `json:"parent_guardians"`
}

type ParentGuardianOut struct {
	AuthorizedRoles []string `json:"authorized_roles"`
	GroupID         string   `json:"account_id" validate:"required"`
	ID              string   `json:"user_id" validate:"required"`
	Username        string   `json:"username" validate:"required"`
}

type StudentOut struct {
	ParentAccountID string `json:"parent_account_id" validate:"required"`
}

