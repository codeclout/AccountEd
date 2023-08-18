package stack

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/dataawsregion"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/dynamodbtable"
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v16/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type config = map[string]interface{}

func NewDynamoDBStorage(internal config, scope constructs.Construct, id string) cdktf.TerraformStack {
	accesskey, ok := internal["AccessKey"].(string)
	if !ok {
		panic("Access Key not set in environment")
	}

	now := time.Now()
	t := time.Unix(0, now.UnixNano()).UTC()

	stack := cdktf.NewTerraformStack(scope, &id)

	config := awsprovider.AwsProviderConfig{
		AccessKey: jsii.String(accesskey),
		AssumeRole: []*awsprovider.AwsProviderAssumeRole{
			&awsprovider.AwsProviderAssumeRole{
				Duration:    jsii.String("45m"),
				RoleArn:     jsii.String(internal["RoleToAssume"].(string)),
				SessionName: jsii.String(internal["SessionLabel"].(string) + strconv.Itoa(t.Nanosecond())),
			},
		},
		Region:    jsii.String("us-east-2"),
		SecretKey: jsii.String(internal["SecretAccessKey"].(string)),
	}

	aws := awsprovider.NewAwsProvider(stack, jsii.String("AWS"), &config)
	dataawsregion.NewDataAwsRegion(stack, jsii.String("region"), &dataawsregion.DataAwsRegionConfig{
		Provider: aws,
		Endpoint: nil,
		Id:       nil,
		Name:     nil,
	})

	dynamodbtable.NewDynamodbTable(stack, jsii.String("sessions_table"), &dynamodbtable.DynamodbTableConfig{
		Lifecycle: &cdktf.TerraformResourceLifecycle{
			IgnoreChanges: []string{"read_capacity", "write_capacity"},
		},
		Provider:     aws,
		Provisioners: nil,
		Name:         jsii.String("io-sch00l-sessions"),
		Attribute: []map[string]string{
			{"name": "id", "type": "S"},
			{"name": "active", "type": "BOOL"},
			{"name": "associated_data", "type": "B"},
			{"name": "created_at", "type": "S"},
			{"name": "encrypted_id", "type": "S"},
			{"name": "forwarded_ip", "type": "S"},
			{"name": "has_auto_correct", "type": "BOOL"},
			{"name": "member_id", "type": "S"},
			{"name": "modified_at", "type": "S"},
			{"name": "nonce", "type": "B"},
			{"name": "ttl", "type": "N"},
		},
		BillingMode:          jsii.String(string(types.BillingModePayPerRequest)),
		GlobalSecondaryIndex: nil,
		HashKey:              jsii.String("id"),
		Id:                   jsii.String("sessions"),
		LocalSecondaryIndex:  nil,
		RangeKey:             jsii.String("member_id"),
		Replica:              nil,
		ServerSideEncryption: nil,
		StreamEnabled:        nil,
		StreamViewType:       nil,
		TableClass:           jsii.String(string(types.TableClassStandard)),
		Tags:                 nil,
		TagsAll:              nil,
		Timeouts:             nil,
		Ttl: &dynamodbtable.DynamodbTableTtl{
			AttributeName: jsii.String("ttl"),
			Enabled:       true,
		},
	})

	return stack
}
