package main

import (
	"sync"

	memberAdapterDriven "github.com/codeclout/AccountEd/pkg/session/adapter/framework/driven/member"
	memberPortDriven "github.com/codeclout/AccountEd/pkg/session/ports/framework/driven/member"

	memberAdapterApi "github.com/codeclout/AccountEd/pkg/session/adapter/api/member"
	memberAdapterCore "github.com/codeclout/AccountEd/pkg/session/adapter/core/member"
	memberAdapterDriver "github.com/codeclout/AccountEd/pkg/session/adapter/framework/drivers/member"
	memberPortApi "github.com/codeclout/AccountEd/pkg/session/ports/api/member"
	memberPortCore "github.com/codeclout/AccountEd/pkg/session/ports/core/member"
	"github.com/codeclout/AccountEd/pkg/session/ports/framework/drivers/member"

	"github.com/codeclout/AccountEd/pkg/monitoring"
	"github.com/codeclout/AccountEd/pkg/session/adapter/api/cloud"
	cloudAdapterCore "github.com/codeclout/AccountEd/pkg/session/adapter/core/cloud"
	cloudAdapterDriven "github.com/codeclout/AccountEd/pkg/session/adapter/framework/driven/cloud"
	cloudAdapterDriver "github.com/codeclout/AccountEd/pkg/session/adapter/framework/drivers/cloud"
	grpcProtocol "github.com/codeclout/AccountEd/pkg/session/adapter/framework/drivers/protocols"
	configuration "github.com/codeclout/AccountEd/pkg/session/adapter/framework/drivers/server"
	cloudPortApi "github.com/codeclout/AccountEd/pkg/session/ports/api/cloud"
	cloudPortCore "github.com/codeclout/AccountEd/pkg/session/ports/core/cloud"
	cloudPortDriven "github.com/codeclout/AccountEd/pkg/session/ports/framework/driven/cloud"
	cloudPortDriver "github.com/codeclout/AccountEd/pkg/session/ports/framework/drivers/cloud"
	"github.com/codeclout/AccountEd/pkg/session/ports/framework/drivers/protocols"
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
	wg.Add(1)
	go monitor.Initialize(&wg)

	sessionConfiguration := configuration.NewAdapter(monitor.Logger)
	internalConfig := sessionConfiguration.LoadSessionConfig()

	awsCoreAdapter = cloudAdapterCore.NewAdapter(monitor.Logger)
	memberCoreAdapter = memberAdapterCore.NewAdapter(*internalConfig, monitor)
	awsDrivenAdapter = cloudAdapterDriven.NewAdapter(*internalConfig, monitor.GetTimeStamp, monitor.Logger)
	awsAPIAdapter = cloud.NewAdapter(awsCoreAdapter, awsDrivenAdapter, monitor.Logger)
	awsDriverAdapter = cloudAdapterDriver.NewAdapter(*internalConfig, awsAPIAdapter, monitor.Logger)

	memberDrivenAdapter = memberAdapterDriven.NewAdapter(*internalConfig, monitor.Logger)
	memberApiAdapter = memberAdapterApi.NewAdapter(*internalConfig, memberCoreAdapter, awsDrivenAdapter, memberDrivenAdapter, monitor.Logger)
	memberDriverAdapter = memberAdapterDriver.NewAdapter(*internalConfig, memberApiAdapter, awsAPIAdapter, monitor.Logger)

	grpcProtocolAdapter = grpcProtocol.NewAdapter(*internalConfig, awsDriverAdapter, memberDriverAdapter, monitor.Logger)
	grpcProtocolAdapter.Run()
}
