package membertypes

const (
	Female Gender = iota + 1
	Male
)

type ErrorCoreDataInvalid error
type PreRegistrationAutoCorrect error

type Gender int8
type ConfigurationPath string
type ContextAPILabel string
type ContextDrivenLabel string
type LogLabel string

type AutoCorrectErrorOutput struct {
	SuggestedEmailAddress string `json:"suggestedEmailAddress"`
}

type Member struct {
	AccountType         string   `json:"accountType"`
	AuthorizedRoles     []string `json:"authorizedRoles"`
	County              string   `json:"county" validate:"required"`
	CreatedAt           int64    `json:"created"`
	DisplayName         string   `json:"displayName"`
	GroupID             string   `json:"groupID"`
	ID                  MemberID `json:"id"`
	Image               any      `json:"image"` // FixMe
	IsActive            bool     `json:"isActive"`
	IsMarkedForDeletion bool     `json:"IsMarkedForDeletion"`
	IsPending           bool     `json:"isPending"`
	IsVerified          bool     `json:"isVerified"`
	LegalFirstName      *string  `json:"legalFirstName" validate:"required"`
	LegalLastName       *string  `json:"legalLastName" validate:"required"`
	MemberType          string   `json:"memberType"`
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
	AutoCorrect         string `json:"autoCorrect" bson:"auto_correct"`
	MemberID            string `json:"memberID" bson:"member_id"`
	RegistrationPending bool   `json:"registration_pending" bson:"registration_pending"`
	SessionID           string `json:"sessionID" bson:"session_id"`
	UsernamePending     bool   `json:"usernamePending" bson:"username_pending"`
}
