package bst

import (
	"crypto/aes"
	"crypto/cipher"
)

// either takes in a 16, 24 or 32 byte key, will panic if fails
func New(encyptionKey []byte, signingKey []byte) (*Token) {
	if len(encyptionKey) != 16 && len(encyptionKey) != 24 && len(encyptionKey) != 32 {
		panic("invalid key length")
	}

	block, err := aes.NewCipher(encyptionKey)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	return &Token{
		block: block,
		gcm: gcm,
	}
}