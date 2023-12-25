package time

import (
	"fmt"
	"time"
)

func GetCurrentTime() (string, error) {
	currentTime := time.Now().Format("02.01.2006 - 15:04")

	if currentTime == "" {
		return "", fmt.Errorf("current time is empty")
	}

	return currentTime, nil
}
