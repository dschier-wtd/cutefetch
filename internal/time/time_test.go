package time

import (
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	// Test case 1: Check if the current time is returned in the correct format
	expectedTimeFormat := "02.01.2006 - 15:04"
	currentTime, err := GetCurrentTime()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if currentTime != time.Now().Format(expectedTimeFormat) {
		t.Fatalf("GetCurrentTime() = %s; want %s", currentTime, time.Now().Format(expectedTimeFormat))
	}

	// Test case 2: Check if the returned time is not empty
	if currentTime == "" {
		t.Fatal("Expected non-empty string, but got empty string")
	}

	// Test case 3: Check if the returned time has the correct length
	if len(currentTime) != 18 {
		t.Fatalf("Expected 16 characters, but got %d characters", len(currentTime))
	}

	// Test case 4: Ensure the function returns the same result for multiple invocations within a short time period
	currentTime1, err := GetCurrentTime()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	time.Sleep(time.Millisecond)
	currentTime2, err := GetCurrentTime()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if currentTime1 != currentTime2 {
		t.Fatalf("GetCurrentTime() returned different values for multiple invocations: %s, %s", currentTime1, currentTime2)
	}
}
