package main

import (
	"sync"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	serverProtocolAdapter "github.com/codeclout/AccountEd/pkg/server/adapters/framework/drivers/protocol"
	memberAdapterCore "github.com/codeclout/AccountEd/session/adapter/core/member"
	memberAdapterDriven "github.com/codeclout/AccountEd/session/adapter/framework/driven/member"
	memberPortDriven "github.com/codeclout/AccountEd/session/ports/framework/driven/member"

	memberAdapterApi "github.com/codeclout/AccountEd/session/adapter/api/member"
	memberAdapterDriver "github.com/codeclout/AccountEd/session/adapter/framework/drivers/member"
	memberPortApi "github.com/codeclout/AccountEd/session/ports/api/member"
	memberPortCore "github.com/codeclout/AccountEd/session/ports/core/member"
	"github.com/codeclout/AccountEd/session/ports/framework/drivers/member"

	"github.com/codeclout/AccountEd/session/adapter/api/cloud"
	cloudAdapterCore "github.com/codeclout/AccountEd/session/adapter/core/cloud"
	cloudAdapterDriven "github.com/codeclout/AccountEd/session/adapter/framework/driven/cloud"
	cloudAdapterDriver "github.com/codeclout/AccountEd/session/adapter/framework/drivers/cloud"
	grpcProtocol "github.com/codeclout/AccountEd/session/adapter/framework/drivers/protocols"
	configuration "github.com/codeclout/AccountEd/session/adapter/framework/drivers/server"
	cloudPortApi "github.com/codeclout/AccountEd/session/ports/api/cloud"
	cloudPortCore "github.com/codeclout/AccountEd/session/ports/core/cloud"
	cloudPortDriven "github.com/codeclout/AccountEd/session/ports/framework/driven/cloud"
	cloudPortDriver "github.com/codeclout/AccountEd/session/ports/framework/drivers/cloud"
	"github.com/codeclout/AccountEd/session/ports/framework/drivers/protocols"
)

func main() {
	var (
		wg sync.WaitGroup

		awsAPIAdapter    cloudPortApi.AWSApiPort
		awsDriverAdapter cloudPortDriver.AWSDriverPort
		awsCoreAdapter   cloudPortCore.AWSCloudCorePort
		awsDrivenAdapter cloudPortDriven.CredentialsAWSPort

		memberDriverAdapter member.SessionDriverMemberPort
		memberApiAdapter    memberPortApi.SessionAPIMemberPort
		memberCoreAdapter   memberPortCore.SessionCoreMemberPort
		memberDrivenAdapter memberPortDriven.SessionDrivenMemberPort

		grpcProtocolAdapter protocols.GRPCProtocolPort
	)

	monitor := monitoring.NewAdapter()

	sessionConfiguration := configuration.NewAdapter("./config.hcl", *monitor)
	internalConfig := sessionConfiguration.LoadStaticConfig()

	awsCoreAdapter = cloudAdapterCore.NewAdapter(*internalConfig, *monitor)
	awsDrivenAdapter = cloudAdapterDriven.NewAdapter(*internalConfig, *monitor)
	awsAPIAdapter = cloud.NewAdapter(awsCoreAdapter, awsDrivenAdapter, *monitor)
	awsDriverAdapter = cloudAdapterDriver.NewAdapter(*internalConfig, awsAPIAdapter, *monitor)

	gRPCAdapterClient := serverProtocolAdapter.NewGrpcAdapter(*internalConfig, *monitor, &wg)
	go gRPCAdapterClient.InitializeClientsForSession()
	defer gRPCAdapterClient.StopProtocolListener()

	memberCoreAdapter = memberAdapterCore.NewAdapter(*internalConfig, monitor)
	memberDrivenAdapter = memberAdapterDriven.NewAdapter(*internalConfig, *monitor)
	memberApiAdapter = memberAdapterApi.NewAdapter(*internalConfig, memberCoreAdapter, awsDrivenAdapter, memberDrivenAdapter, gRPCAdapterClient, *monitor, &wg)
	memberDriverAdapter = memberAdapterDriver.NewAdapter(*internalConfig, memberApiAdapter, awsAPIAdapter, *monitor)

	grpcProtocolAdapter = grpcProtocol.NewAdapter(*internalConfig, awsDriverAdapter, memberDriverAdapter, *monitor)
	grpcProtocolAdapter.Run()
}
