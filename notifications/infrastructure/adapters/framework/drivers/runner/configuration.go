package runner

import (
  "fmt"
  "os"
  "reflect"
  "strconv"
  "time"

  "github.com/aws/constructs-go/constructs/v10"
  "github.com/aws/jsii-runtime-go"
  awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v16/provider"
  "github.com/cdktf/cdktf-provider-aws-go/aws/v16/sesdomainidentity"
  "github.com/cdktf/cdktf-provider-aws-go/aws/v16/sesdomainmailfrom"
  "github.com/hashicorp/terraform-cdk-go/cdktf"
  "golang.org/x/exp/slog"
)

type environment struct {
  AccessKey       string
  Region          string
  RoleArn         string
  SecretAccessKey string
  SessionLabel    string
}

type Adapter struct {
  app cdktf.App
  log *slog.Logger
}

func NewAdapter(app cdktf.App, log *slog.Logger) *Adapter {
  return &Adapter{
    app: app,
    log: log,
  }
}

func (a *Adapter) LoadNotificationsInfrastructureConfig() *map[string]interface{} {
  var out = make(map[string]interface{})
  var s string

  envConfig := environment{
    AccessKey:       os.Getenv("AWS_ACCESS_KEY_ID"),
    Region:          os.Getenv("AWS_REGION"),
    RoleArn:         os.Getenv("AWS_ROLE_TO_ASSUME"),
    SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
    SessionLabel:    os.Getenv("AWS_SESSION_LABEL"),
  }

  val := reflect.ValueOf(&envConfig).Elem()

  for i := 0; i < val.NumField(); i++ {
    out[val.Type().Field(i).Name] = val.Field(i).Interface()
  }

  for k, v := range out {
    switch x := v.(type) {
    case string:
      if x == (s) {
        a.log.Error(fmt.Sprintf("Notification:%s is not defined in the environment", k))
        os.Exit(1)
      }
    default:
      panic("invalid Notification configuration type")
    }
  }

  return &out

}

func (a *Adapter) InitializeInfrastructure() {
  stack := a.NewMyStack(a.app, "infrastructure")

  cdktf.NewCloudBackend(stack, &cdktf.CloudBackendConfig{
    Hostname:     jsii.String("app.terraform.io"),
    Organization: jsii.String("sch00l"),
    Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String("io-sch00l-notifications")),
  })
}

func (a *Adapter) NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
  internal := *a.LoadNotificationsInfrastructureConfig()

  now := time.Now()
  t := time.Unix(0, now.UnixNano()).UTC()

  stack := cdktf.NewTerraformStack(scope, &id)

  config := awsprovider.AwsProviderConfig{
    AccessKey: jsii.String(internal["AccessKey"].(string)),
    AssumeRole: []*awsprovider.AwsProviderAssumeRole{
      &awsprovider.AwsProviderAssumeRole{
        Duration:    jsii.String("45m"),
        RoleArn:     jsii.String(internal["RoleArn"].(string)),
        SessionName: jsii.String(internal["SessionLabel"].(string) + strconv.Itoa(t.Nanosecond())),
      },
    },
    Region:    jsii.String("us-east-2"),
    SecretKey: jsii.String(internal["SecretAccessKey"].(string)),
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
