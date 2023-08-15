package driven

import (
	"context"
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

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	"github.com/codeclout/AccountEd/storage/ports/framework/driven"
	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

type storeSessionIn = storageTypes.PreRegistrationSessionAPIin

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

func (a *Adapter) GetDynamoClient(ctx context.Context, creds *credentials.StaticCredentialsProvider, region *string) (*dynamodb.Client, error) {
	// if a.clientCache != nil {
	// 	return a.clientCache, nil
	// }
	var endpoint string

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
		awsconfig.WithCredentialsProvider(*creds),
	)
	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil, storageTypes.ErrorDefaultConfiguration(errors.New("unable to load DynamoDB configuration"))
	}

	// @TODO - store and check client expiration
	session := dynamodb.NewFromConfig(dynamoConfig)
	a.clientCache = session

	return session, nil
}

func (a *Adapter) StoreSession(ctx context.Context, api driven.DynamodbAPI, in storeSessionIn) (*storageTypes.PreRegistrationSessionDrivenOut, error) {
	tableName := in.SessionStorageTableName

	now := time.Now()
	ts := time.Unix(0, now.UnixNano()).UTC()
	ttl := now.Add(15 * time.Minute).Unix()

	data := &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"id":               &types.AttributeValueMemberN{Value: in.SessionID},
			"active":           &types.AttributeValueMemberBOOL{Value: !in.HasAutoCorrect},
			"associated_data":  &types.AttributeValueMemberB{Value: in.AssociatedData},
			"created_at":       &types.AttributeValueMemberN{Value: ts.Format(time.RFC3339)},
			"encrypted_id":     &types.AttributeValueMemberN{Value: in.EncryptedSessionID},
			"forwarded_ip":     &types.AttributeValueMemberN{Value: in.ForwardedIP},
			"has_auto_correct": &types.AttributeValueMemberBOOL{Value: in.HasAutoCorrect},
			"member_id":        &types.AttributeValueMemberN{Value: in.MemberID},
			"modified_at":      &types.AttributeValueMemberN{Value: ts.Format(time.RFC3339)},
			"nonce":            &types.AttributeValueMemberB{Value: in.Nonce},
			"ttl":              &types.AttributeValueMemberN{Value: strconv.FormatInt(ttl, 10)},
		},
		TableName: aws.String(tableName),
	}

	result, e := api.PutItem(ctx, data)
	if e != nil {
		a.monitor.LogGrpcError(ctx, e.Error())
		return nil, status.Errorf(codes.Internal, "unable to save session id -> ", in.SessionID)
	}

	out := storageTypes.PreRegistrationSessionDrivenOut{
		Active:         !in.HasAutoCorrect,
		Attributes:     result.Attributes,
		CreatedAt:      ts,
		ExpiresAt:      ttl,
		MemberId:       in.MemberID,
		ResultMetadata: result.ResultMetadata,
	}

	return &out, nil
}
