package core

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
	notifications "github.com/codeclout/AccountEd/notifications/notification-types"
	"github.com/codeclout/AccountEd/pkg/monitoring"
	"github.com/codeclout/AccountEd/pkg/validations"
)

type Adapter struct {
	config  map[string]interface{}
	monitor monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) setRegistrationPending(delivery, qscore string, hasMX, isDisposable, isRole bool) bool {
	quality, _ := strconv.ParseFloat(qscore, 32)

	return quality < 0.80 || delivery != "DELIVERABLE" || isDisposable || isRole || !hasMX
}

// useAutoCorrect determines if Autocorrect is applicable for the input autocorrectString. It returns true if the length of autocorrectString
// is greater than 0, otherwise it returns false.
func (a *Adapter) useAutoCorrect(autocorrectString string) bool {
	return len(autocorrectString) > 0
}

func (a *Adapter) ProcessEmailValidation(ctx context.Context, in notifications.ValidateEmailOut) (*pb.ValidateEmailAddressResponse, error) {
	var autoCorrectAddress, memberId string

	if a.useAutoCorrect(in.Autocorrect) {
		memberId = ""

		x, e := validations.ValidateEmail(&in.Autocorrect)
		if e != nil {
			const msg = "error validating auto corrected email address"
			a.monitor.LogGrpcError(ctx, fmt.Sprintf(msg+": %s", e.Error()))
			return nil, status.Error(codes.Internal, msg)
		}

		suggestedMemberID, ok := x.Load().(string)
		if !ok {
			const msg = "error loading auto corrected email address"
			a.monitor.LogGrpcError(ctx, fmt.Sprintf(msg+": %s", e.Error()))
			return nil, status.Error(codes.Internal, msg)
		}

		autoCorrectAddress = suggestedMemberID
	} else {
		memberId = in.Email
	}

	out := pb.ValidateEmailAddressResponse{
		AutoCorrect:          autoCorrectAddress,
		MemberId:             memberId,
		ShouldConfirmAddress: a.setRegistrationPending(in.Deliverability, in.QualityScore, in.IsMxFound.GetValue(), in.IsDisposableEmail.GetValue(), in.IsRoleEmail.GetValue()),
		MemberIdPending:      a.useAutoCorrect(in.Autocorrect),
	}

	return &out, nil
}

func (a *Adapter) SendPreRegistrationEmailCore(ctx context.Context, in *notifications.NoReplyEmailIn) (*notifications.NoReplyEmailInput, error) {
	body := "Before you can begin using your new account, it needs to be activated - " +
		"this is a necessary step to verify that the email address you provided is valid and owned by you. " +
		"To proceed with the activation of your account, please click on the link provided below " +
		"or completely copy and paste it into your browser's address bar."
	sessionLink := fmt.Sprintf("%s/members/confirm?t=%s", in.Domain, in.Token)
	welcome := fmt.Sprintf("Welcom to %s", in.Domain)

	msg := fmt.Sprintf("%s\n\n%s\n\n%s\n\nThanks!\nThe "+in.Domain+" staff", welcome, body, sessionLink)

	out := notifications.NoReplyEmailInput{
		Body:    msg,
		Subject: "Please confirm your email address",
	}

	return &out, nil
}
