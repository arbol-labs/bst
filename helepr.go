package bst

import (
	"hash"
)


 
func generateHash(tokenType string, data []byte, h hash.Hash) (hash []byte) {
	h.Write([]byte(tokenType))
	h.Write(data)
	hash = h.Sum(nil)
	h.Reset()

	return hash
}