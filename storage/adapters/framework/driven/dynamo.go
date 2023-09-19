package driven

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	"github.com/codeclout/AccountEd/storage/ports/framework/driven"
	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

type cc = context.Context

type FetchTokenIn = storageTypes.FetchTokenIn
type FetchTokenOut = storageTypes.FetchTokenResult
type TokenStorePayload = storageTypes.TokenStorePayload
type TokenStoreResult = storageTypes.TokenStoreResult

type Adapter struct {
	DynamoClient driven.DynamodbAPI
	clientCache  *dynamodb.Client
	config       map[string]interface{}
	monitor      monitoring.Adapter
	waitGroup    *sync.WaitGroup
}

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter, wg *sync.WaitGroup) *Adapter {
	return &Adapter{
		config:    config,
		monitor:   monitor,
		waitGroup: wg,
	}
}

func (a *Adapter) GetDynamoClient(ctx context.Context, in credentials.StaticCredentialsProvider, region *string) (*dynamodb.Client, error) {
	var endpoint string

	dynamoCredentials := &credentials.StaticCredentialsProvider{Value: in.Value}

	isDemo, ok := a.config["DemoMode"].(bool)
	if !ok {
		a.monitor.LogGenericError("ambiguous application mode")
		return nil, storageTypes.ErrorInvalidConfiguration(errors.New("configuration error: DemoMode"))
	}

	if isDemo {
		uri, ok := a.config["DynamoDemo"].(string)
		if !ok {
			a.monitor.LogGenericError("dynamodb endpoint not configured")
			return nil, storageTypes.ErrorInvalidConfiguration(errors.New("configuration error: DynamoEndpoint"))
		}

		endpoint = uri
	}

	dynamoConfig, e := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(*region),
		awsconfig.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, awsregion string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: endpoint}, nil
			})),
		awsconfig.WithCredentialsProvider(dynamoCredentials),
	)

	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil, storageTypes.ErrorDefaultConfiguration(errors.New("unable to load DynamoDB configuration"))
	}

	session := dynamodb.NewFromConfig(dynamoConfig)
	a.clientCache = session

	return session, nil
}

func (a *Adapter) GetTokenItem(ctx context.Context, client driven.DynamodbAPI, in FetchTokenIn) (*FetchTokenOut, error) {
	data := &dynamodb.GetItemInput{
		Key:       map[string]types.AttributeValue{"token": &types.AttributeValueMemberS{Value: in.Token}},
		TableName: aws.String(in.TableName),
	}

	resp, e := client.GetItem(ctx, data)
	if e != nil {
		const msg = "error occurred while trying to get item"

		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, status.Error(codes.Internal, msg)
	}

	out := FetchTokenOut{}

	e = attributevalue.UnmarshalMap(resp.Item, &out)
	if e != nil {
		const msg = "failed to unmarshal GetItem response to struct, %v"

		a.monitor.LogGrpcError(ctx, fmt.Sprintf(msg, e))
		return nil, status.Error(codes.Internal, fmt.Sprintf(msg, e))
	}

	return &out, nil
}

func (a *Adapter) StoreToken(ctx cc, client driven.DynamodbAPI, in *TokenStorePayload) (*TokenStoreResult, error) {
	ts := a.getTimeStamp()
	ttl := strconv.FormatInt(int64(in.Ttl), 10)

	data := &dynamodb.PutItemInput{
		ConditionExpression: aws.String("attribute_not_exists(member_id)"),
		Item: map[string]types.AttributeValue{
			"active":           &types.AttributeValueMemberBOOL{Value: !in.HasAutoCorrect},
			"created_at":       &types.AttributeValueMemberS{Value: ts.Format(time.RFC3339)},
			"has_auto_correct": &types.AttributeValueMemberBOOL{Value: in.HasAutoCorrect},
			"member_id":        &types.AttributeValueMemberS{Value: in.MemberId},
			"modified_at":      &types.AttributeValueMemberS{Value: ts.Format(time.RFC3339)},
			"public_key":       &types.AttributeValueMemberS{Value: in.PublicKey},
			"token":            &types.AttributeValueMemberS{Value: in.Token},
			"token_id":         &types.AttributeValueMemberS{Value: in.TokenId},
			"ttl":              &types.AttributeValueMemberN{Value: ttl},
		},
		TableName: aws.String(in.SessionTableName),
	}

	result, e := client.PutItem(ctx, data)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())

		var conditionalCheckFailed *types.ConditionalCheckFailedException
		if errors.As(e, &conditionalCheckFailed) {
			var ErrorSessionExists storageTypes.ErrorSessionExists = errors.New("session already exists")
			return nil, status.Error(codes.AlreadyExists, ErrorSessionExists.Error())
		}

		return nil, status.Errorf(codes.Internal, "unable to save token id -> %s", in.TokenId)
	}

	expiresAt := a.getExpiresAt(ttl)

	out := storageTypes.TokenStoreResult{
		Active:         !in.HasAutoCorrect,
		CreatedAt:      ts,
		ExpiresAt:      expiresAt,
		HasAutoCorrect: in.HasAutoCorrect,
		Token:          in.Token,
		ResultMetadata: result.ResultMetadata,
	}

	return &out, nil
}

func (a *Adapter) getExpiresAt(ttl string) time.Time {
	expiresAt, e := strconv.ParseInt(ttl, 10, 64)
	if e != nil {
		a.monitor.LogGenericError(fmt.Sprintf("unable to process expiresAt -> %s", e.Error()))
		return time.Unix(0, 0)
	}

	t := time.Unix(expiresAt, 0)

	return t
}

func (a *Adapter) getTimeStamp() time.Time {
	now := time.Now()
	ts := time.Unix(0, now.UnixNano()).UTC()

	return ts
}
