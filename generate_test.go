package bst

import (
	"fmt"
	"testing"
	"time"
)

func  BenchmarkRoundTrip(b *testing.B) {
	key1 := make([]byte, 32)
	key2 := make([]byte, 32)
	encoder := New(key1, key2)


	b.ResetTimer()

	for b.Loop() {
		token, err := encoder.GenerateTTLToken(time.Now().Add(1 * time.Hour))
		if err != nil {
		b.Fatal(err)
		}

		err = encoder.ValidateTTLToken(token)
		if err != nil {
			b.Fatal(err)
		}
	}
}


func  TestRoundTrip(t *testing.T) {
	key1 := make([]byte, 32)
	key2 := make([]byte, 32)
	encoder := New(key1, key2)

	f := StandardFields{
		IssuedAt: time.Now(),
		Expiration: time.Now(),
		Audience: "hey",
		Issuer: "hey",
	}

	token, err := encoder.GenerateCustomToken(f)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(token)
}