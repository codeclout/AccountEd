package cloud

import (
	"context"

	pb "github.com/codeclout/AccountEd/session/gen/aws/v1"
)

type AWSDriverPort interface {
	GetAWSSessionCredentials(ctx context.Context, request *pb.AWSConfigRequest) (*pb.AWSConfigResponse, error)
}
