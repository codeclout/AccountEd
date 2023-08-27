package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/credentials"
)

type AWSCloudCorePort interface {
	ConvertCredentialsToBytes(ctx context.Context, in *credentials.StaticCredentialsProvider) ([]byte, error)
}
