package configuration

import (
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v16/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StackModules struct {
	App cdktf.App
}

type StackDataProvider struct {
	Environment  *string
	Host         *string
	Organization *string
	Provider     *awsprovider.AwsProviderConfig
}
