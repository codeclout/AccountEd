package api

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"

	memberTypes "github.com/codeclout/AccountEd/members/member-types"
	pb "github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
)

func (a *Adapter) processAWSCredentials(ctx context.Context) *aws.Config {
	var creds = &memberTypes.CredentialsAWS{}

	grpcProtocol := *a.grpcProtocol.AWSClient

	arn, ok := a.config["AWSRolePreRegistration"].(string)
	if !ok {
		a.log.Error("pre registration role not set in environment")
		return nil
	}

	region, ok := a.config["AWSRegion"].(string)
	if !ok {
		a.log.Error("region not set in environment")
		return nil
	}

	request := pb.AWSConfigRequest{
		Arn:    arn,
		Region: region,
	}

	sessionPkg, e := grpcProtocol.GetAWSSessionCredentials(ctx, &request)
	if e != nil {
		a.log.Error(e.Error())
		return nil
	}

	if sessionPkg == nil {
		a.log.Error("sessionPkg is nil")
		return nil
	}

	e = json.Unmarshal(sessionPkg.GetAwsSession(), &creds)
	if e != nil {
		a.log.Error(e.Error())
		return nil
	}

	credsProvider := credentials.NewStaticCredentialsProvider(creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken)

	c := aws.Config{
		Region:      a.config["AWSRegion"].(string),
		Credentials: credsProvider,
	}

	return &c
}

func (a *Adapter) getAWSCredentialBytes(ctx context.Context) []byte {
	config := a.processAWSCredentials(ctx)
	creds := config.Credentials

	b, e := json.Marshal(creds)
	if e != nil {
		a.log.Error(e.Error())
		return nil
	}

	return b
}
