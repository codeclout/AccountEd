package member

import (
	"context"

	"golang.org/x/exp/slog"

	cloudAWS "github.com/codeclout/AccountEd/session/ports/api/cloud"
	"github.com/codeclout/AccountEd/session/ports/api/member"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/pkg/errors"

	sessiontypes "github.com/codeclout/AccountEd/session/session-types"

	awspb "github.com/codeclout/AccountEd/session/gen/aws/v1"
	pb "github.com/codeclout/AccountEd/session/gen/members/v1"
)

var transactionLable = sessiontypes.LogLabel("transaction_id")

type Adapter struct {
	api    member.SessionAPIMemberPort
	awsapi cloudAWS.AWSApiPort
	config map[string]interface{}
	log    *slog.Logger
}

func NewAdapter(config map[string]interface{}, api member.SessionAPIMemberPort, awsapi cloudAWS.AWSApiPort, log *slog.Logger) *Adapter {
	return &Adapter{
		api:    api,
		awsapi: awsapi,
		config: config,
		log:    log,
	}
}

func (a Adapter) GetEncryptedSessionId(ctx context.Context, request *pb.EncryptedStringRequest) (*pb.EncryptedStringResponse, error) {
	arn := a.config["RoleToAssume"].(string)
	id := request.GetSessionId()
	region := a.config["Region"].(string)

	data := sessiontypes.AmazonConfigurationInput{
		ARN:    aws.String(arn),
		Region: aws.String(region),
	}

	ch := make(chan *awspb.AWSConfigResponse, 1)
	echan := make(chan error, 1)
	uch := make(chan *pb.EncryptedStringResponse, 1)

	ctx = context.WithValue(ctx, transactionLable, arn+"|"+region)
	a.awsapi.GetAWSSessionCredentials(ctx, data, ch, echan)

	select {
	case session := <-ch:
		a.log.Log(ctx, 4, string(session.GetAwsCredentials()))
		a.api.EncryptSessionId(ctx, session.AwsCredentials, id, uch, echan)
	}

	select {
	case <-ctx.Done():
		t := ctx.Value(transactionLable)
		a.log.Error("request timeout", transactionLable, t.(string))
		return nil, errors.New("request timeout")

	case out := <-uch:
		t := ctx.Value(transactionLable)
		a.log.Info("success", transactionLable, t.(string))
		return out, nil

	case e := <-echan:
		t := ctx.Value(transactionLable)
		a.log.Error(e.Error(), transactionLable, t.(string))
		return nil, e
	}
}
