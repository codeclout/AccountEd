package out

import (
  "context"
  "fmt"

  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/config"
  "github.com/aws/aws-sdk-go-v2/credentials"
  "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoActions struct {
  Client         *dynamodb.Client
  StorageAdapter *Adapter
  TableName      string
}

func (a *Adapter) NewDynamoDb(creds *aws.Config) *DynamoActions {
  identity, _ := creds.Credentials.Retrieve(context.TODO())

  cfg, e := config.LoadDefaultConfig(context.TODO(),
    config.WithRegion(a.RuntimeConfig["AwsRegion"].(string)),

    config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
      func(service, region string, options ...interface{}) (aws.Endpoint, error) {
        return aws.Endpoint{URL: a.RuntimeConfig["DynamoDbEndpoint"].(string)}, nil
      })),

    config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
      Value: aws.Credentials{
        AccessKeyID:     identity.AccessKeyID,
        SecretAccessKey: identity.SecretAccessKey,
        SessionToken:    identity.SessionToken,
      },
    }),
  )

  if e != nil {
    a.log("fatal", e.Error())
  }

  return &DynamoActions{
    Client:         dynamodb.NewFromConfig(cfg),
    StorageAdapter: a,
    TableName:      a.RuntimeConfig["DynamoDbTableName"].(string),
  }
}

func (d *DynamoActions) Initialize() {
  _, e := d.Client.DescribeTable(context.TODO(),
    &dynamodb.DescribeTableInput{
      TableName: aws.String(d.StorageAdapter.RuntimeConfig["DynamoDbTableName"].(string))})

  if e != nil {
    d.StorageAdapter.log("fatal", fmt.Sprintf("unable to connect to dynamo: %s", e.Error()))
  }

  d.StorageAdapter.log("info", "successfully connected to dynamodb")
}

func (d *DynamoActions) CloseConnection() {}
