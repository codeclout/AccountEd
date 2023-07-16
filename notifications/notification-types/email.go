package notification_types

import (
	"github.com/google/uuid"

	pb "github.com/codeclout/AccountEd/notifications/gen/v1"
)

type URL string
type EmailAddress string
type EmailList []string
type ErrorEmailVerificationProcessor error
type ErrorStaticConfig error
type SessionID string

type EmailDrivenIn struct {
	EmailAddress string     `json:"email_address"`
	Endpoint     string     `json:"endpoint"`
	SessionID    *uuid.UUID `json:"session_id"`
}

type NoReplyEmailIn struct {
	AWSCredentials []byte
	Domain         string
	FromAddress    string
	SessionID      string
	ToAddress      []string
}

type NoReplyEmailInput struct {
	Body    string
	Subject string
}

type NoReplyEmailOut struct {
	MessageID string
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
