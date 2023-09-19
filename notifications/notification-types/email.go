package notification_types

import (
	pb "github.com/codeclout/AccountEd/notifications/gen/email/v1"
)

type ProcessEmailValidationOut struct {
	AutoCorrect          string `json:"auto_correct" bson:"auto_correct"`
	MemberID             string `json:"member_id" bson:"member_id"`
	ShouldConfirmAddress bool   `json:"should_confirm_address" bson:"should_confirm_address"`
	MemberIdPending      bool   `json:"member_id_pending" bson:"member_id_pending"`
}

type NoReplyEmailIn struct {
	AWSCredentials []byte
	Domain         string
	FromAddress    string
	Token          string
	ToAddress      []string
}

type NoReplyEmailInput struct {
	Body    string
	Subject string
}

type NoReplyEmailOut struct {
	MessageID string
}

type ValidateEmailIn struct {
	Address           string
	ProcessorDomain   string
	ProcessorEndpoint string
	ProcessorKey      string
}

type ValidateEmailOut struct {
	Autocorrect       string                       `json:"autocorrect"`
	Deliverability    string                       `json:"deliverability"`
	Email             string                       `json:"email"`
	IsCatchallEmail   *pb.EmailVerificationPayload `json:"is_catchall_email"`
	IsDisposableEmail *pb.EmailVerificationPayload `json:"is_disposable_email"`
	IsFreeEmail       *pb.EmailVerificationPayload `json:"is_free_email"`
	IsMxFound         *pb.EmailVerificationPayload `json:"is_mx_found"`
	IsRoleEmail       *pb.EmailVerificationPayload `json:"is_role_email"`
	IsSMTPValid       *pb.EmailVerificationPayload `json:"is_smtp_valid"`
	IsValidFormat     *pb.EmailVerificationPayload `json:"is_valid_format"`
	QualityScore      string                       `json:"quality_score"`
}
