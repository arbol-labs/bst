package bst

import (
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

	var z StandardFields

	token, err := encoder.GenerateCustomToken(f)
	if err != nil {
		t.Fatal(err)
	}

	err = encoder.ParseToken(token, &z)
	if err != nil {
		t.Fatal(err)
	}

	if z.IssuedAt.Unix() != f.IssuedAt.Unix() {
		t.Errorf("IssuedAt mismatch: got %v, want %v", z.IssuedAt, f.IssuedAt)
	}
	if z.Expiration.Unix() != f.Expiration.Unix() {
		t.Errorf("Expiration mismatch: got %v, want %v", z.Expiration, f.Expiration)
	}
	if z.Audience != f.Audience {
		t.Errorf("Audience mismatch: got %v, want %v", z.Audience, f.Audience)
	}
	if z.Issuer != f.Issuer {
		t.Errorf("Issuer mismatch: got %v, want %v", z.Issuer, f.Issuer)
	}
}