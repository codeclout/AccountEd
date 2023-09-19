package sessiontypes

import (
	"time"

	"aidanwoods.dev/go-paseto"
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

type TokenCreateOut struct {
	Token string
	*TokenPayload
	TTL time.Duration
}

type SessionStoreMetadata struct {
	HasAutoCorrect  bool
	MemberID        string
	TokenIdentifier string
}

type NewTokenPayload struct {
	HasAutoCorrect bool   `json:"has_auto_correct"`
	MemberId       string `json:"member_id"`
	TokenId        string `json:"token_id"`
}

type TokenPayload struct {
	ExpiresAt     time.Time `json:"expires_at"`
	ID            string    `json:"id"`
	IssuedAt      time.Time `json:"issued_at"`
	MemberID      string    `json:"member_id"`
	Private       paseto.V4AsymmetricSecretKey
	Public        paseto.V4AsymmetricPublicKey
	PublicEncoded string
}

type ValidateTokenPayload struct {
	Token string `json:"token"`
}

type Transport string
