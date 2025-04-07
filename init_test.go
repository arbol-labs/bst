package bst

import "testing"

func TestNew(t *testing.T) {
	tests := []struct{
		name string
		key1 []byte
		key2 []byte
		expPanic bool
	}{
		{
			name: "valid 16 byte token cong",
			key1: make([]byte, 16),
			key2: []byte("hello world"),
			expPanic: false,
		},
		{
			name: "valid 24 byte token cong",
			key1: make([]byte, 24),
			key2: []byte("hello world"),
			expPanic: false,
		},
		{
			name: "valid 32 byte token cong",
			key1: make([]byte, 32),
			key2: []byte("hello world"),
			expPanic: false,
		},
		{
			name: "invalid  token conf",
			key1: make([]byte, 234),
			key2: []byte("hello world"),
			expPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil && tt.expPanic {
					t.Errorf("Expected panic but got none")
				} 
			}()
			
			New(tt.key1, tt.key2)
		})
	}
}