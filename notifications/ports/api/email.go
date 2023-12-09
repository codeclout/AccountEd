package api

import (
	"context"

	notification "github.com/codeclout/AccountEd/notifications/notification-types"

	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
)

type ValidateEmailIn = notification.ValidateEmailIn

type EmailApiPort interface {
	ValidateEmailAddress(ctx context.Context, in *ValidateEmailIn, ch chan *pb.ValidateEmailAddressResponse, ech chan error)
}
