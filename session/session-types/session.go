package sessiontypes

import (
	"time"
)

type AuthenticationAlgorithm int8

type AuthenticationCredential struct {
	Authenticator
	CredentialType      string                    `json:"credential_type"`
	ID                  string                    `json:"id"`
	SupportedAlgorithms []AuthenticationAlgorithm `json:"supported_algorithms"`
	Transports          []Transport               `json:"transports"`
}

type Authenticator struct {
	ResidentKey        string `json:"resident_key"`
	MemberVerification string `json:"member_verification"`
}

type AuthenticationOptions struct {
	FederatedLogin    bool `json:"federated_login"`
	LoginLink         bool `json:"login_link"`
	UsernamePinTOTP   bool `json:"legacy_password"`
	WebAuthentication bool `json:"web_authentication"`
}

type DefaultRouteDuration int

type RegistrationOptions struct {
	AttestationType           string                     `json:"attestation_type"`
	AuthenticationCredentials []AuthenticationCredential `json:"authentication_credentials"`
	MemberDisplayName         string                     `json:"member_display_name"`
	MemberID                  string                     `json:"member_id"`
	MemberLegalName           string                     `json:"member_legal_name"`
	RelyingPartyID            string                     `json:"relying_party_id"`
	RelyingPartyName          string                     `json:"relying_party_name"`
	Timeout                   time.Duration              `json:"timeout"`
}

type SessionStoreMetadata struct {
	HasAutoCorrect  bool
	MemberID        string
	TokenIdentifier string
}

type ValidateTokenPayload struct {
	Token string `json:"token"`
}

type Transport string
