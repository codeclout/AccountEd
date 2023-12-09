package main

import (
	"sync"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	serverProtocolAdapter "github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	"github.com/codeclout/AccountEd/storage/adapters/api"
	"github.com/codeclout/AccountEd/storage/adapters/core"
	"github.com/codeclout/AccountEd/storage/adapters/framework/driven"
	"github.com/codeclout/AccountEd/storage/adapters/framework/drivers"
	"github.com/codeclout/AccountEd/storage/adapters/framework/drivers/protocol"
	"github.com/codeclout/AccountEd/storage/adapters/framework/drivers/server"
	storageApiPort "github.com/codeclout/AccountEd/storage/ports/api"
	storageCorePort "github.com/codeclout/AccountEd/storage/ports/core"
	storageDrivenPort "github.com/codeclout/AccountEd/storage/ports/framework/driven"
	storageDriverPort "github.com/codeclout/AccountEd/storage/ports/framework/drivers"
)

var staticConfigurationPath = "./config.hcl"

func main() {
	var (
		dynamodbAPIAdapter    storageApiPort.DynamoDbApiPort
		dynamodbCoreAdapter   storageCorePort.DynamoDbCorePort
		dynamodbDrivenAdapter storageDrivenPort.DynamoDbDrivenPort
		dynamodbDriverAdapter storageDriverPort.DynamoDbDriverPort

		wg sync.WaitGroup
	)

	monitor := monitoring.NewAdapter()

	storageConfiguration := server.NewAdapter(*monitor, staticConfigurationPath)
	config := *storageConfiguration.LoadStaticConfig()

	gRPCAdapter := serverProtocolAdapter.NewGrpcAdapter(config, *monitor, &wg)
	go gRPCAdapter.InitializeClientsForStorage()
	defer gRPCAdapter.StopProtocolListener()

	dynamodbDrivenAdapter = driven.NewAdapter(config, *monitor, &wg)
	dynamodbCoreAdapter = core.NewAdapter(config, *monitor)
	dynamodbAPIAdapter = api.NewAdapter(config, dynamodbCoreAdapter, gRPCAdapter, dynamodbDrivenAdapter, *monitor)
	dynamodbDriverAdapter = drivers.NewAdapter(config, dynamodbAPIAdapter, *monitor, &wg)

	gRPCprotocol := protocol.NewAdapter(config, dynamodbDriverAdapter, *monitor, &wg)
	gRPCprotocol.StorageRun()
}
