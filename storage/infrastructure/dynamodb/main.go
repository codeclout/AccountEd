package main

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	"github.com/codeclout/AccountEd/storage/adapters/framework/drivers/server"
	"github.com/codeclout/AccountEd/storage/infrastructure/dynamodb/stack"
)

func main() {
	var staticConfigurationPath = "./config.hcl"
	monitor := monitoring.NewAdapter()

	storageconfig := server.NewAdapter(*monitor, staticConfigurationPath)
	baseconfig := *storageconfig.LoadStorageConfig()

	infraconfig := stack.NewAdapter(*monitor)
	config := infraconfig.LoadStorageInfrastructureConfig(baseconfig)

	app := cdktf.NewApp(nil)
	dynamoStack := stack.NewDynamoDBStorage(*config, app, "io-sch00l-storage-dynamodb")

	cdktf.NewCloudBackend(dynamoStack, &cdktf.CloudBackendConfig{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("sch00l"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String("io-sch00l-storage")),
	})

	app.Synth()
}
