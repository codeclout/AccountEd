package member

import (
	"context"
	"fmt"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	cloudAWS "github.com/codeclout/AccountEd/session/ports/api/cloud"
	"github.com/codeclout/AccountEd/session/ports/api/member"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/pkg/errors"

	sessiontypes "github.com/codeclout/AccountEd/session/session-types"

	awspb "github.com/codeclout/AccountEd/session/gen/aws/v1"
	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
)

type mtr = monitoring.Adapter
type cloudApi = cloudAWS.AWSApiPort
type memberApi = member.SessionAPIMemberPort

type Adapter struct {
	api     memberApi
	aws     cloudApi
	config  map[string]interface{}
	monitor mtr
}

func NewAdapter(config map[string]interface{}, api memberApi, awsapi cloudApi, monitor mtr) *Adapter {
	return &Adapter{
		api:     api,
		aws:     awsapi,
		config:  config,
		monitor: monitor,
	}
}

func (a Adapter) GetEncryptedSessionId(ctx context.Context, request *pb.EncryptedStringRequest) (*pb.EncryptedStringResponse, error) {
	arn, ok := a.config["RoleToAssume"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'RoleToAssume' in config")
	}

	region, ok := a.config["Region"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'Region' in config")
	}

	sessionIdRequestRole := sessiontypes.AmazonConfigurationInput{
		RoleArn: aws.String(arn),
		Region:  aws.String(region),
	}

	ch := make(chan *awspb.AWSConfigResponse, 1)
	ech := make(chan error, 1)
	uch := make(chan *pb.EncryptedStringResponse, 1)

	ctx = context.WithValue(ctx, a.monitor.LogLabelTransactionID, arn+"|"+region)

	apiData := sessiontypes.SessionStoreMetadata{
		HasAutoCorrect: request.GetHasAutoCorrect(),
		MemberID:       request.GetMemberId(),
		SessionID:      request.GetSessionId(),
	}

	a.aws.GetAWSSessionCredentials(ctx, sessionIdRequestRole, ch, ech)

	select {
	case session := <-ch:
		a.api.EncryptSessionId(ctx, session.AwsCredentials, &apiData, uch, ech)
	}

	select {
	case <-ctx.Done():
		const msg = "get encrypted session id request timeout"
		a.monitor.LogGrpcError(ctx, msg)
		return nil, errors.New(msg)

	case out := <-uch:
		a.monitor.LogGrpcInfo(ctx, "success")
		return out, nil

	case e := <-ech:
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, e
	}
}
