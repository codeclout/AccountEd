package membertypes

import (
	"time"

	pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
)

type ParentGuardian struct {
	Member
	Phone    string  `json:"phone"`
	Username *string `json:"username" bson:"username" validate:"required,email"`
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
	PrimaryMember *ParentGuardian `json:"parent_guardians"`
}

type HomeSchoolRegisterOut struct {
	PrimaryMember *ParentGuardianOut `json:"parent_guardians"`
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

// EmailValidationIn represents an input object for validating an email address. It contains properties related to the email validation, including
// autocorrect, deliverability, quality score, and various email type checks (e.g., free email, disposable email, role email, catchall email, MX record, and
// SMTP validation). The email type checks utilize protobuf EmailVerificationPayload objects.
type EmailValidationIn struct {
	Email             string                       `json:"email"`
	Autocorrect       string                       `json:"autocorrect"`
	Deliverability    string                       `json:"deliverability"`
	QualityScore      string                       `json:"quality_score"`
	IsValidFormat     *pb.EmailVerificationPayload `json:"is_valid_format"`
	IsFreeEmail       *pb.EmailVerificationPayload `json:"is_free_email"`
	IsDisposableEmail *pb.EmailVerificationPayload `json:"is_disposable_email"`
	IsRoleEmail       *pb.EmailVerificationPayload `json:"is_role_email"`
	IsCatchallEmail   *pb.EmailVerificationPayload `json:"is_catchall_email"`
	IsMxFound         *pb.EmailVerificationPayload `json:"is_mx_found"`
	IsSmtpValid       *pb.EmailVerificationPayload `json:"is_smtp_valid"`
}
