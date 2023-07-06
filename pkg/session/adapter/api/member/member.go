package cloud

import (
  "context"

  pb "github.com/codeclout/AccountEd/pkg/session/gen/v1/sessions"
)

func (a *Adapter) EncryptSessionId(ctx context.Context, creds []byte, id, key string, uch chan *pb.EncryptedStringResponse, echan chan error) {
  
}
