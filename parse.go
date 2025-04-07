package bst

import "encoding/hex"


func (t *Token) ParseToken(token string, dst any) error {
	decodedHex, err := hex.DecodeString(token)
	if err != nil {
		return err
	}

	nonce := decodedHex[:t.gcm.NonceSize()]
	ct := decodedHex[t.gcm.NonceSize():]

	data, err := t.gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return err
	}

	return json.Unmarshal(data ,&dst)
}