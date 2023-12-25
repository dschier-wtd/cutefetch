package system

import (
	"testing"
)

func TestGetUsername(t *testing.T) {
	username, err := GetUsername()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if username == "" {
		t.Fatalf("Expected non-empty string, but got empty string")
	}
}

func TestGetHostname(t *testing.T) {
	hostname, err := GetHostname()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if hostname == "" {
		t.Fatalf("Expected non-empty string, but got empty string")
	}
}

func TestGetOS(t *testing.T) {
	os, err := GetOS()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if os == "" {
		t.Fatalf("Expected non-empty string, but got empty string")
	}
}

func TestGetShell(t *testing.T) {
	shell, err := GetShell()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if shell == "" {
		t.Fatalf("Expected non-empty string, but got empty string")
	}
}
