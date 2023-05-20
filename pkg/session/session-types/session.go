package sessiontypes

import (
  "time"

  "github.com/o1egl/paseto"
)

type APIRequestToken paseto.JSONToken

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

type StatelessAPI struct {
  TokenData APIRequestToken
}

type Transport string
