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

	f := StandardFields{
		IssuedAt: time.Now(),
		Expiration: time.Now(),
		Audience: "hey",
		Issuer: "hey",
	}

	var x StandardFields

	b.ResetTimer()

	for b.Loop() {
		token, err := encoder.GenerateCustomToken(f)
		if err != nil {
		b.Fatal(err)
		}

		err = encoder.ParseToken(token, &x)
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