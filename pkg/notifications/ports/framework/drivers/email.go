package drivers

import (
  "context"

  pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
)

type (
  IsLocal   bool
  IsNonProd bool
  IsProd    bool
)

type EmailDriverPort interface {
  HandleValidateEmailAddress(ctx context.Context, email *pb.ValidateEmailAddressRequest) (*pb.ValidateEmailAddressResponse, error)
}
