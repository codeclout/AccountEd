package driver

import (
  "context"
  "errors"
  "time"

  "github.com/go-webauthn/webauthn/webauthn"
  "github.com/o1egl/paseto"
  "golang.org/x/crypto/ed25519"
  "golang.org/x/exp/slog"

  "github.com/codeclout/AccountEd/pkg/monitoring"
  sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

type Adapter struct {
  log       *slog.Logger
  monitor   *monitoring.Adapter
  publicKey *ed25519.PublicKey
  token     *paseto.V2
  webAuth   *webauthn.WebAuthn
}

func NewAdapter(monitor *monitoring.Adapter) (*Adapter, error) {
  return &Adapter{
    log:     monitor.Logger,
    monitor: monitor,
    token:   paseto.NewV2(),
  }, nil
}

func (a *Adapter) CreateToken(ctx context.Context, duration time.Duration, groupid, memberid string) (*string, error) {
  // var out sessiontypes.APIRequestToken
  //
  // id := uuid.New()
  //
  // publicKey, privateKey, e := ed25519.GenerateKey(nil)
  // if e != nil {
  //   a.log.Error("Error creating token ID", e)
  //   return nil, fmt.Errorf("attempt to create public/private key pair for token failed %s", e.Error())
  // }
  //
  // e = a.storePrivateKey(publicKey, privateKey)
  // if e != nil {
  //   a.log.Error("Error creating token ID", e)
  //   return nil, fmt.Errorf("attempt to store private key for token failed %s", e.Error())
  // }
  //
  // session := paseto.JSONToken{
  //   Audience:   groupid,
  //   Issuer:     ctx.Value("issuer").(string),
  //   Jti:        id.String(),
  //   Subject:    memberid,
  //   Expiration: a.monitor.GetTimeStamp().Add(duration),
  //   IssuedAt:   a.monitor.GetTimeStamp(),
  //   NotBefore:  a.monitor.GetTimeStamp(),
  // }
  // session.Set("publicKey", string(publicKey))
  //
  // token, e := a.token.Sign(privateKey, session, nil)
  return nil, nil
}

func (a *Adapter) VerifyTokenData(ctx context.Context) (*sessiontypes.StatelessAPI, error) {
  return nil, errors.New("not implemented")
}
