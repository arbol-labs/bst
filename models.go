package bst

import (
	"crypto/cipher"
	"time"
)


type Token struct {
	block cipher.Block
	gcm cipher.AEAD
}


type StandardFields struct {
	IssuedAt time.Time  `json:"iat"`
	Expiration time.Time `json:"exp"`
	Audience string `json:"aud"`
	Issuer string `json:"is"`
}