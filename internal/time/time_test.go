package time

import (
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	// Test case 1: Check if the current time is returned in the correct format
	expected := time.Now().Format("02.01.2006 - 15:04:05")
	actual := GetCurrentTime()
	if expected != actual {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}

	// Test case 2: Check if the returned time is not empty
	actual = GetCurrentTime()
	if actual == "" {
		t.Errorf("Expected non-empty string, but got empty string")
	}
}
