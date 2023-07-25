package api

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/credentials"

	pb "github.com/codeclout/AccountEd/session/gen/aws/v1"
)

func (a *Adapter) processAWSCredentials(ctx context.Context) *credentials.StaticCredentialsProvider {
	var creds credentials.StaticCredentialsProvider

	grpcProtocol := *a.grpcProtocol.AWS_SessionClient

	arn, ok := a.config["AWSRolePreRegistration"].(string)
	if !ok {
		a.monitor.LogGenericError("pre registration role not set in environment")
		return nil
	}

	region, ok := a.config["AWSRegion"].(string)
	if !ok {
		a.monitor.LogGenericError("region not set in environment")
		return nil
	}

	request := pb.AWSConfigRequest{
		RoleArn: arn,
		Region:  region,
	}

	sessionPkg, e := grpcProtocol.GetAWSSessionCredentials(ctx, &request)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil
	}

	if sessionPkg == nil {
		a.monitor.LogGenericError("sessionPkg is nil")
		return nil
	}

	e = json.Unmarshal(sessionPkg.GetAwsCredentials(), &creds)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil
	}

	return &creds
}

func (a *Adapter) getAWSCredentialBytes(ctx context.Context) []byte {
	creds := a.processAWSCredentials(ctx)

	b, e := json.Marshal(creds)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil
	}

	return b
}
