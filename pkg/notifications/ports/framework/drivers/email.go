package drivers

import (
  "context"

  pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
)

type EmailDriverPort interface {
  ValidateEmailAddress(ctx context.Context, email *pb.ValidateEmailAddressRequest) (*pb.ValidateEmailAddressResponse, error)
}
