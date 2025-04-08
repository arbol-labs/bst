package bst

import (
	"crypto/cipher"
	"hash"
	"strings"
	"time"
)


type Token struct {
	hash hash.Hash
	block cipher.Block
	gcm cipher.AEAD
	builder strings.Builder
}


type StandardFields struct {
	IssuedAt time.Time  `json:"iat"`
	Expiration time.Time `json:"exp"`
	Audience string `json:"aud"`
	Issuer string `json:"is"`
}