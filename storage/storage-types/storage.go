package storage_types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/smithy-go/middleware"
)

type ErrorInvalidConfiguration error
type ErrorDefaultConfiguration error
type ErrorSessionExists error

type FetchTokenIn struct {
	Credentials credentials.StaticCredentialsProvider
	TableName   string
	Token       string
}

type FetchTokenResult struct {
	Active         bool   `dynamodbav:"active"`
	HasAutoCorrect bool   `dynamodbav:"has_auto_correct"`
	MemberID       string `dynamodbav:"member_id"`
	PublicKey      string `dynamodbav:"public_key"`
	Token          string `dynamodbav:"token"`
	TokenId        string `dynamodbav:"token_id"`
}

type TokenStoreResult struct {
	Active         bool
	CreatedAt      time.Time
	ExpiresAt      time.Time
	HasAutoCorrect bool
	Token          string
	ResultMetadata middleware.Metadata
}

type TokenStorePayload struct {
	AWSRegion        *string                               `json:"aws_region"`
	Credentials      credentials.StaticCredentialsProvider `json:"sessionServiceAWScredentials"`
	HasAutoCorrect   bool                                  `json:"has_auto_correct"`
	MemberId         string                                `json:"member_id"`
	PublicKey        string                                `json:"publicKey"`
	SessionTableName string                                `json:"sessionTableName"`
	Token            string                                `json:"token"`
	TokenId          string                                `json:"tokenId"`
	Ttl              int32                                 `json:"ttl"`
}
