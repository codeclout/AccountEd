package main

import (
	"strconv"
	"time"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v16/provider"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/sesdomainidentity"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/sesdomainmailfrom"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	now := time.Now()
	t := time.Unix(0, now.UnixNano()).UTC()
	stack := cdktf.NewTerraformStack(scope, &id)

	config := awsprovider.AwsProviderConfig{
		AccessKey: jsii.String("AKIASWBCNCZGYYY2OWQI"),
		AssumeRole: []*awsprovider.AwsProviderAssumeRole{
			&awsprovider.AwsProviderAssumeRole{
				Duration:    jsii.String("45m"),
				RoleArn:     jsii.String("arn:aws:iam::184755754573:role/notifications-sch00l.io-service"),
				SessionName: jsii.String("MySession" + strconv.Itoa(t.Nanosecond())),
			},
		},
		Region:    jsii.String("us-east-2"),
		SecretKey: jsii.String("YYeqRoav111pSkvWP4HpdzxwycSLhp1gdMrlerO8"),
	}

	aws := awsprovider.NewAwsProvider(stack, jsii.String("AWS"), &config)

	domainIdentity := sesdomainidentity.NewSesDomainIdentity(stack, jsii.String("sch00lDomain"), &sesdomainidentity.SesDomainIdentityConfig{
		Provider: aws,
		Domain:   jsii.String("sch00l.io"),
	})

	sesdomainmailfrom.NewSesDomainMailFrom(stack, jsii.String("sesMailFrom"), &sesdomainmailfrom.SesDomainMailFromConfig{
		Domain:         domainIdentity.Domain(),
		MailFromDomain: jsii.String("notifications." + *domainIdentity.Domain()),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	stack := NewMyStack(app, "infrastructure")
	cdktf.NewCloudBackend(stack, &cdktf.CloudBackendConfig{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("sch00l"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String("io-sch00l-notifications")),
	})

	app.Synth()
}
