package membertypes

const (
	Female Gender = iota + 1
	Male
)

type Gender int8
type ConfigurationPath string
type ContextAPILabel string
type ContextDrivenLabel string

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

type MemberErrorOut struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
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

type PrimaryMemberConfirmationOut struct {
	IsPrimaryMemberConfirmed bool `json:"isConfirmed"`
}

type PrimaryMemberStartRegisterIn struct {
	MemberID *string `json:"member_id"`
	Pin      *string `json:"pin"`
}

type ValidatedEmailResonse struct { //nolint:maligned
	AutoCorrect         string `json:"auto_correct"`
	MemberID            string `json:"member_id"`
	RegistrationPending bool   `json:"registration_pending"`
	TokenID             string `json:"token_id"`
	Token               string `json:"token"`
	UsernamePending     bool   `json:"username_pending"`
}
