package api

import (
  "context"

  pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
)

type EmailApiPort interface {
  ValidateEmailAddress(ctx context.Context, address string, ch chan *pb.ValidateEmailAddressResponse, errorch chan error)
}
