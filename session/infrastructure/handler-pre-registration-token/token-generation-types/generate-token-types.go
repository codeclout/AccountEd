package token_generation_types

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

type GenerateTokenRequest struct {
	HasAutoCorrect bool   `json:"has_auto_correct"`
	MemberId       string `json:"member_id"`
	TokenId        string `json:"token_id"`
}

type GenerateTokenResponse struct {
	Token string `json:"token,omitempty"`
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

type TokenCreateOut struct {
	Token string
	*TokenPayload
	TTL time.Duration
}
