package time

import "time"

func GetCurrentTime() string {
	currentTime := time.Now().Format("02.01.2006 - 15:04:05")
	return currentTime
}
