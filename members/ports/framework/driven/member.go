package driven

import (
	"context"

	memberTypes "github.com/codeclout/AccountEd/members/member-types"
	"github.com/codeclout/AccountEd/notifications/gen/email/v1"
	membersessionv1 "github.com/codeclout/AccountEd/session/gen/members/v1"
)

type cc = context.Context
type ec = emailv1.EmailNotificationServiceClient
type pm = memberTypes.PrimaryMemberStartRegisterIn
type pc = memberTypes.PrimaryMemberConfirmationOut
type mmo = memberTypes.MemberErrorOut

type EmailValidationResponse = emailv1.ValidateEmailAddressResponse

type MemberDrivenPort interface {
	ConfirmEmailAddressValidation(ctx cc, sessionClient *membersessionv1.MemberSessionClient, token string) (*pc, *mmo)
	ValidateEmailAddress(ctx cc, data *pm, emailClient *ec) (*EmailValidationResponse, *mmo)
}
