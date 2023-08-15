package member

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/codeclout/AccountEd/pkg/monitoring/adapters/framework/drivers"

	"github.com/pkg/errors"

	sessiontypes "github.com/codeclout/AccountEd/session/session-types"
)

type Adapter struct {
	config  map[string]interface{}
	monitor *drivers.Adapter
}

func NewAdapter(config map[string]interface{}, monitor *drivers.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		monitor: monitor,
	}
}

func (a *Adapter) ProcessSessionIdEncryption(ctx context.Context) (*sessiontypes.SessionIdEncryptionOut, error) {
	key, ok := ctx.Value(sessiontypes.ContextDrivenLabel("driven_input")).(string)
	if !ok {
		return nil, errors.New("invalid session id key")
	}

	id, ok := ctx.Value(sessiontypes.ContextAPILabel("api_input")).(string)
	if !ok {
		return nil, errors.New("invalid session id")
	}

	keyBytes := []byte(key)
	if len(keyBytes) != 32 {
		return nil, errors.New("key length must equal 32 bytes")
	}

	associatedData := []byte(a.monitor.GetTimeStamp().String())

	internalKey, e := aes.NewCipher([]byte(key))
	if e != nil {
		return nil, e
	}

	gcm, e := cipher.NewGCM(internalKey)
	if e != nil {
		return nil, e
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, e = io.ReadFull(rand.Reader, nonce); e != nil {
		return nil, e
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(id), associatedData)
	cipherOut := base64.URLEncoding.EncodeToString(ciphertext)

	out := sessiontypes.SessionIdEncryptionOut{
		AssociatedData: associatedData,
		CipherText:     &cipherOut,
		IV:             nonce,
		SessionID:      &id,
	}

	return &out, nil
}

func (a *Adapter) ProcessSessionIdDecryption(associatedData, key []byte, cipherIn *string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext, _ := base64.StdEncoding.DecodeString(*cipherIn)

	// Nonce size should be the same as Block size
	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]

	// Open (decrypt) the ciphertext and authenticate it with the given associated data
	plaintext, err := gcm.Open(nil, nonce, ciphertext, associatedData)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
