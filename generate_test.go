package bst

import (
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

	b.ResetTimer()


	var x StandardFields


	for b.Loop() {
		token, _ := encoder.GenerateToken(f)
		encoder.ParseToken(token, &x)
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

	token, err := encoder.GenerateToken(f)
	if err != nil {
		t.Fatal(err)
	}

	var z StandardFields

	err = encoder.ParseToken(token, &z)
	if err != nil {
		t.Fatal(err)
	}

}