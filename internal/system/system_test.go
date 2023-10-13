package system

import (
	"testing"
)

func TestGetUsername(t *testing.T) {
	username, err := GetUsername()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if username == "" {
		t.Errorf("Expected non-empty string, but got empty string")
	}
}

func TestGetHostname(t *testing.T) {
	hostname, err := GetHostname()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if hostname == "" {
		t.Errorf("Expected non-empty string, but got empty string")
	}
}

func TestGetOS(t *testing.T) {
	os, err := GetOS()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if os == "" {
		t.Errorf("Expected non-empty string, but got empty string")
	}
}
