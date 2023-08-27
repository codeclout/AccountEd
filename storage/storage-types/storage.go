package storage_types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
)

type ErrorCredentialsRetrieval error
type ErrorInvalidConfiguration error
type ErrorDefaultConfiguration error
type ErrorSessionExists error

type DefaultRouteDuration int

type PreRegistrationSessionAPIin struct {
	AssociatedData            []byte
	EncryptedSessionID        string
	ForwardedIP               string
	HasAutoCorrect            bool
	MemberID                  string
	Nonce                     []byte
	SessionID                 string
	SessionServiceCredentials *credentials.StaticCredentialsProvider
	SessionStorageTableName   string
	TTL                       int32
}

type PreRegistrationSessionDrivenOut struct {
	Active     bool
	Attributes map[string]types.AttributeValue
	CreatedAt  time.Time
	ExpiresAt  int64

	MemberId       string
	ResultMetadata middleware.Metadata
}
