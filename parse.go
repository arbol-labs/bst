package bst

import (
	"crypto/hmac"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/arbol-labs/bst/pkg/variables"
)

// this function is used to validate tokens with cutom fields. it takes in the token and the destination struct, it will decode the token fields into the token
func (t *Token) ParseToken(token string, dst *any) error {
	sections := strings.Split(token, ".")

	if len(sections) != 3 {
		return fmt.Errorf("invalid token")
	}

	if sections[0] != variables.CustomFieldsTokenType {
		return fmt.Errorf("invalid token type")
	}

	
	decodedHash, err := hex.DecodeString(sections[2])
	if err != nil {
		return fmt.Errorf("failed to decode hash section")
	}

	decodedCipher, err := hex.DecodeString(sections[1])
	if err != nil {
		return err
	}
	
	hash := generateHash(variables.CustomFieldsTokenType, decodedCipher, t.hash)

	if !hmac.Equal(decodedHash, hash) {
		return fmt.Errorf("token has been tampered with")
	}

	nonce := decodedCipher[:t.gcm.NonceSize()]
	ct := decodedCipher[t.gcm.NonceSize():]

	data, err := t.gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return err
	}

	return json.Unmarshal(data ,dst)
}

func (t *Token) ValidateTTLToken(token string) (error) {
	sections := strings.Split(token, ".")

	if len(sections) != 3 {
		return fmt.Errorf("invalid token")
	}

	if sections[0] != variables.TtlToken {
		return fmt.Errorf("invalid token type")
	}

	decodedTTL, err := hex.DecodeString(sections[1])
	if err != nil {
		return err
	}

	decodedHash, err := hex.DecodeString(sections[2])
	if err != nil {
		return fmt.Errorf("failed to decode hash section")
	}

	hash := generateHash(variables.TtlToken, decodedTTL, t.hash)
	if !hmac.Equal(decodedHash, hash) {
		return fmt.Errorf("token has been tampered with")
	}

	if !hmac.Equal(decodedHash, hash) {
		return fmt.Errorf("token has been tampered with")
	}


	ti, err := time.Parse(time.RFC3339, string(decodedTTL))
	if err != nil {
		return err
	}

	if time.Now().After(ti) {
		return fmt.Errorf("token has expired")
	}

	return nil

}