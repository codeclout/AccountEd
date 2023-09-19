package api

import (
	"context"

	notification "github.com/codeclout/AccountEd/notifications/notification-types"

	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
)

type cc = context.Context

type NoReplyEmailIn = notification.NoReplyEmailIn
type ValidateEmailIn = notification.ValidateEmailIn

type EmailApiPort interface {
	SendPreRegistrationEmailAPI(ctx cc, in *NoReplyEmailIn, ch chan *pb.NoReplyEmailNotificationResponse, ech chan error)
	ValidateEmailAddress(ctx context.Context, in *ValidateEmailIn, ch chan *pb.ValidateEmailAddressResponse, ech chan error)
}
