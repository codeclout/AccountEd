package cloud

import (
	"context"

	pb "github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
)

type AWSDriverPort interface {
	GetAWSSessionCredentials(ctx context.Context, request *pb.AWSConfigRequest) (*pb.AWSConfigResponse, error)
}
