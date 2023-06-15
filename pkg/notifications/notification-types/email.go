package notification_types

import (
	"github.com/google/uuid"

	pb "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
)

type EmailAddress string
type ErrorEmailVerificationProcessor error
type ErrorStaticConfig error

type EmailDrivenIn struct {
	EmailAddress string     `json:"email_address"`
	Endpoint     string     `json:"endpoint"`
	SessionID    *uuid.UUID `json:"session_id"`
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
