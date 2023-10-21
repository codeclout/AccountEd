package member

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	cloudAWS "github.com/codeclout/AccountEd/session/ports/api/cloud"
	"github.com/codeclout/AccountEd/session/ports/api/member"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/pkg/errors"

	sessionTypes "github.com/codeclout/AccountEd/session/session-types"

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

func (a *Adapter) getRoleARN() (*string, error) {
	arn, ok := a.config["RoleToAssume"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'RoleToAssume' in config")
	}

	return aws.String(arn), nil
}

func (a *Adapter) getAWSRegion() (*string, error) {
	region, ok := a.config["Region"].(string)
	if !ok {
		return nil, fmt.Errorf("missing or invalid 'Region' in config")
	}

	return aws.String(region), nil
}

func (a *Adapter) getAWSCredentialInput() (*sessionTypes.AmazonConfigurationInput, error) {
	arn, e := a.getRoleARN()
	if e != nil {
		return nil, status.Error(codes.Unavailable, e.Error())
	}

	region, e := a.getAWSRegion()
	if e != nil {
		return nil, status.Error(codes.Unavailable, e.Error())
	}

	metadata := sessionTypes.AmazonConfigurationInput{
		RoleArn: arn,
		Region:  region,
	}

	return &metadata, nil
}

func (a *Adapter) processTokenValidation(in *pb.ValidateTokenRequest) (*sessionTypes.ValidateTokenPayload, error) {
	var s string

	if in.GetToken() == (s) {
		return nil, status.Error(codes.InvalidArgument, "invalid token validation request")
	}

	return &sessionTypes.ValidateTokenPayload{Token: in.GetToken()}, nil
}

func (a *Adapter) ValidateMemberToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	metadata, e := a.getAWSCredentialInput()
	if e != nil {
		return nil, status.Error(codes.FailedPrecondition, e.Error())
	}

	ch := make(chan *awspb.AWSConfigResponse, 1)
	ech := make(chan error, 1)
	uch := make(chan *pb.ValidateTokenResponse, 1)

	a.aws.GetAWSSessionCredentials(ctx, *metadata, ch, ech)

	select {
	case session := <-ch:
		apiData, e := a.processTokenValidation(request)
		if e != nil {
			return nil, e
		}

		a.api.ValidateMemberToken(ctx, session.GetAwsCredentials(), apiData, uch, ech)
	}

	select {
	case <-ctx.Done():
		const msg = "decryption operation request timeout"
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
