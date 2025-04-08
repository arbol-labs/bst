package bst

import (
	"crypto/rand"
	"encoding/hex"

	"time"

	"github.com/arbol-labs/bst/pkg/variables"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary


func (t *Token) GenerateCustomToken(fields any) (token string, err error) {
	data, err := json.Marshal(fields)
	if err != nil {
		return "", err
	}

	var nonce = make([]byte, t.gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}
	encrypted := t.gcm.Seal(nil, nonce, data, nil)

	nonce = append(nonce, encrypted...)

	hash := generateHash(variables.CustomFieldsTokenType, nonce, t.hash)

	t.builder.WriteString(variables.CustomFieldsTokenType)
	t.builder.WriteRune('.')
	t.builder.WriteString(hex.EncodeToString(nonce))
	t.builder.WriteRune('.')
	t.builder.WriteString(hex.EncodeToString(hash))

	token = t.builder.String()

	t.builder.Reset()

	return token, nil
}





// <type>.<ttl>.<hash>
// no encryption
// this function will generate a token which only conatains a ttl, best used for simple and lightweight tokens
func (t *Token) GenerateTTLToken(ttl time.Time) (token string, err error) {
	ttlStr := ttl.Format(time.RFC3339)

	tokenBytes := []byte(ttlStr)

	t.builder.WriteString(variables.TtlToken)
	t.builder.WriteRune('.')
	t.builder.WriteString(hex.EncodeToString(tokenBytes))

	hash := generateHash(variables.TtlToken, tokenBytes, t.hash)
	t.builder.WriteRune('.')
	t.builder.WriteString(hex.EncodeToString(hash))
	token = t.builder.String()

	t.builder.Reset()

	return token, nil
}