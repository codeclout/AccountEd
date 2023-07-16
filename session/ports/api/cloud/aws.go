package cloud

import (
	"golang.org/x/net/context"

	sessiontypes "github.com/codeclout/AccountEd/session/session-types"

	pb "github.com/codeclout/AccountEd/session/gen/aws/v1"
)

type AWSApiPort interface {
	GetAWSSessionCredentials(ctx context.Context, in sessiontypes.AmazonConfigurationInput, out chan *pb.AWSConfigResponse, echan chan error)
}
