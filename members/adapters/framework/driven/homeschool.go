package driven

import (
	"context"

	"github.com/pkg/errors"

	memberTypes "github.com/codeclout/AccountEd/members/member-types"
	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	monitoring "github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"
)

type emailClient = pb.EmailNotificationServiceClient
type pmrStart = memberTypes.PrimaryMemberStartRegisterIn

type Adapter struct {
	monitor monitoring.Adapter
}

func NewAdapter(monitor monitoring.Adapter) *Adapter {
	return &Adapter{monitor: monitor}
}

func (a *Adapter) ValidateEmailAddress(ctx context.Context, data *pmrStart, emailClient *emailClient) (*pb.ValidateEmailAddressResponse, error) {
	var workflowLabel = memberTypes.ContextPreRegistrationWorkflowLabel("ValidateEmailAddress")

	ctx = context.WithValue(ctx, workflowLabel, "called")

	if v := *emailClient; v == nil {
		const msg = "nil notifications gRPC client"
		a.monitor.LogGenericError(msg)
		return nil, errors.New(msg)
	}

	if x := data; x == nil {
		const msg = "nil primary member email address"
		a.monitor.LogGenericError(msg)
		return nil, errors.New(msg)
	}

	client := *emailClient
	response, e := client.ValidateEmailAddress(ctx, &pb.ValidateEmailAddressRequest{Address: *data.Username})

	if e != nil {
		a.monitor.LogGrpcError(ctx, *data.Username)
		return nil, errors.Wrapf(e, "registerAccountAPI -> core.PreRegister(%v)", *data)
	}

	return response, nil
}
