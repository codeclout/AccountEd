package cloud

import (
	"golang.org/x/net/context"

	pb "github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
	sessiontypes "github.com/codeclout/AccountEd/pkg/session/session-types"
)

type AWSApiPort interface {
	GetAWSSessionCredentials(ctx context.Context, in sessiontypes.AmazonConfigurationInput, out chan *pb.AWSConfigResponse, echan chan error)
}
