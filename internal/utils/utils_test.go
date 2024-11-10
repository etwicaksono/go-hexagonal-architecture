package utils

import (
	"fmt"
	"testing"
)

func TestPasswordGenerate(t *testing.T) {
	tests := []struct {
		name      string
		raw       string
		expectErr bool
	}{
		{"ValidPassword", "TestPassword123", false},
		{"ValidPassword", "TestPassword123", false},
		{"EmptyPassword", "", false}, // bcrypt handles empty strings, but they are not recommended in practice
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, err := PasswordGenerate(tt.raw)
			fmt.Println("tt.raw: ", tt.raw)
			fmt.Println("hashedPassword: ", hashedPassword)

			if (err != nil) != tt.expectErr {
				t.Fatalf("expected error: %v, got: %v", tt.expectErr, err)
			}

			// Ensure that hashed password is different from the raw password if hashing succeeded
			if !tt.expectErr && hashedPassword == tt.raw {
				t.Errorf("expected hashed password to be different from raw password")
			}

			// Verify that the generated hash matches the raw password
			if !tt.expectErr {
				if err := PasswordVerify(hashedPassword, tt.raw); err != nil {
					t.Errorf("expected hash verification to succeed, got error: %v", err)
				}
			}
		})
	}
}

func TestPasswordVerify(t *testing.T) {
	tests := []struct {
		name          string
		raw           string
		hashPassword  func(string) (string, error) // function to generate a hash for setup
		inputPassword string
		expectErr     bool
	}{
		{
			name:          "CorrectPassword",
			raw:           "TestPassword123",
			hashPassword:  PasswordGenerate,
			inputPassword: "TestPassword123",
			expectErr:     false,
		},
		{
			name:          "IncorrectPassword",
			raw:           "TestPassword123",
			hashPassword:  PasswordGenerate,
			inputPassword: "WrongPassword123",
			expectErr:     true,
		},
		{
			name:          "EmptyPassword",
			raw:           "",
			hashPassword:  PasswordGenerate,
			inputPassword: "",
			expectErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, err := tt.hashPassword(tt.raw)
			if err != nil {
				t.Fatalf("failed to generate hash: %v", err)
			}

			err = PasswordVerify(hashedPassword, tt.inputPassword)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
		})
	}
}
