package cloud

import "context"

type AWSCloudCorePort interface {
  GetServiceIdMetadata(ctx context.Context)
}
