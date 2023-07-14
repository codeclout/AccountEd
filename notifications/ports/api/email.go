package api

import (
  "context"

  notification "github.com/codeclout/AccountEd/pkg/notifications/notification-types"

  pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
)

type EmailApiPort interface {
	SendPreRegistrationEmailAPI(ctx context.Context, in *notification.NoReplyEmailIn, ch chan *pb.NoReplyEmailNotificationResponse, errorch chan error)
	ValidateEmailAddress(ctx context.Context, address string, ch chan *pb.ValidateEmailAddressResponse, errorch chan error)
}
