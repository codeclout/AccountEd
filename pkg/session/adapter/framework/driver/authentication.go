package driver

import (
  "context"
  "os"
  "time"

  "github.com/go-webauthn/webauthn/protocol"
  "github.com/go-webauthn/webauthn/webauthn"

  sessiontypes "github.com/codeclout/AccountEd/pkg/session/session-types"
)

func (a *Adapter) HandleAuthenticationOptions(ctx context.Context, id string) sessiontypes.AuthenticationOptions {
}

func useDebug() bool {
  if os.Getenv("ENVIRONMENT") == "prod" {
    return false
  }
  return true
}

func (a *Adapter) HandleRegistrationOptions(ctx context.Context, id string) sessiontypes.RegistrationOptions {

  a.webAuth.Config = &webauthn.Config{
    RPID:                   ctx.Value("relying_party_id").(string),
    RPDisplayName:          ctx.Value("relying_party_display_name").(string),
    RPOrigins:              []string{"get from config"},
    AttestationPreference:  protocol.PreferNoAttestation,
    AuthenticatorSelection: protocol.AuthenticatorSelection{},
    Debug:                  useDebug(),
    EncodeUserIDAsString:   false,
    Timeouts: webauthn.TimeoutsConfig{
      Login: struct {
        Enforce    bool
        Timeout    time.Duration
        TimeoutUVD time.Duration
      }{Enforce: true, Timeout: ctx.Value("login_timeout").(time.Duration), TimeoutUVD: 0},
    },
  }

}
