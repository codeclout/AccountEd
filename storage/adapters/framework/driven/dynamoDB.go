package driven

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pkg/errors"

	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
	storageTypes "github.com/codeclout/AccountEd/storage/storage-types"
)

type client interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(options *dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type Adapter struct {
	DynamoClient client
	config       map[string]interface{}
	monitor      monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter, wg *sync.WaitGroup) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) getDynamoClient(ctx context.Context, creds *credentials.StaticCredentialsProvider, region *string) (*dynamodb.Client, error) {
	endpoint, ok := a.config["DynamoEndpoint"]
	if !ok {
		a.monitor.LogGenericError("dynamodb endpoint not configured")
		return nil, storageTypes.ErrorInvalidConfiguration(errors.New("configuration error: DynamoEndpoint"))
	}

	dynamoConfig, e := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithRegion(*region),
		awsconfig.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, awsregion string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: endpoint.(string)}, nil
			})),
		awsconfig.WithCredentialsProvider(*creds),
	)

	if e != nil {
		a.monitor.LogGenericError(e.Error())
		return nil, storageTypes.ErrorDefaultConfiguration(errors.New("unable to load DynamoDB configuration"))
	}

	// @TODO - cache client
	// @TODO - store and check client expiration
	client := dynamodb.NewFromConfig(dynamoConfig)
	return client, nil
}

func (a *Adapter) StoreSession(ctx context.Context) error {
	return nil
}
