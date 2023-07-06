package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/pkg/errors"

	pb "github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
	sessiontypes "github.com/codeclout/AccountEd/pkg/session/session-types"
)

func (a *Adapter) GetEncryptedSessionId(ctx context.Context, request *pb.EncryptedStringRequest) (*pb.EncryptedStringResponse, error) {
	arn := a.config["RoleToAssume"].(string)
	id := request.GetSessionId()
	key := request.GetKey()
	region := a.config["Region"].(string)

	data := sessiontypes.AmazonConfigurationInput{
		ARN:    aws.String(arn),
		Region: aws.String(region),
	}

	ch := make(chan *pb.AWSConfigResponse, 1)
	echan := make(chan error, 1)
	uch := make(chan *pb.EncryptedStringResponse, 1)

	ctx = context.WithValue(ctx, transactionIdLogLabel, arn+"|"+region)
	a.api.GetAWSSessionCredentials(ctx, data, ch, echan)

	select {
	case <-ctx.Done():
		t := ctx.Value(transactionIdLogLabel)
		a.log.Error("request timeout", transactionIdLogLabel, t.(string))
		return nil, errors.New("request timeout")

	case awsSession := <-ch:
		a.api.EncryptSessionId(ctx, awsSession, id, key, uch, echan)

	case out := <-uch:
		t := ctx.Value(transactionIdLogLabel)
		a.log.Info("success", transactionIdLogLabel, t.(string))
		return out, nil

	case e := <-echan:
		t := ctx.Value(transactionIdLogLabel)
		a.log.Error(e.Error(), transactionIdLogLabel, t.(string))
		return nil, e
	}
}
