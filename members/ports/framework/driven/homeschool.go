package driven

import (
	"context"

	memberTypes "github.com/codeclout/AccountEd/members/member-types"
	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
)

type emailClient = pb.EmailNotificationServiceClient
type pmrStart = memberTypes.PrimaryMemberStartRegisterIn

type HomeschoolDrivenPort interface {
	ValidateEmailAddress(ctx context.Context, data *pmrStart, emailClient *emailClient) (*pb.ValidateEmailAddressResponse, error)
}
