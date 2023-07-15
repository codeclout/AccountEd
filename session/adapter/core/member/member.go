package member

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/codeclout/AccountEd/pkg/monitoring"

	"github.com/pkg/errors"
	"golang.org/x/exp/slog"

	sessiontypes "github.com/codeclout/AccountEd/pkg/session/session-types"
)

type Adapter struct {
	config  map[string]interface{}
	log     *slog.Logger
	monitor *monitoring.Adapter
}

func NewAdapter(config map[string]interface{}, monitor *monitoring.Adapter) *Adapter {
	return &Adapter{
		config:  config,
		log:     monitor.Logger,
		monitor: monitor,
	}
}

func (a *Adapter) ProcessSessionIdEncryption(ctx context.Context, id, key string) (*sessiontypes.SessionIdEncryptionOut, error) {
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
