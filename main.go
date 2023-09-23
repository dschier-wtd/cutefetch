package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"time"

	color "github.com/dschier-wtd/cutefetch/packages"
)

func getUser() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Print("Error")
		return ""
	}
	return currentUser.Username
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("Error:", err)
	}
	return hostname
}

func getTime() string {
	currentTime := time.Now()
	formatedTime := currentTime.Format("02.01.2006 - 15:04:05")
	return formatedTime
}

func getOS() string {
	content, err := os.ReadFile("/etc/os-release")

	if err != nil {
		log.Println("Error:", err)
		return "Unknown"
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				distribution := strings.Trim(parts[1], `"`)

				return distribution
			}
		}
	}

	return "Unknown"
}

func main() {

	fmt.Println()
	fmt.Println("(\\ /) \t\t" + color.Red + "User:\t" + color.Reset +
		getUser() + "@" + getHostname())
	fmt.Println("( · ·) \t\t" + color.Green + "Time:\t" + color.Reset +
		getTime())
	fmt.Println("c(" + color.Red + "\"" + color.Reset + ")(" + color.Red +
		"\"" + color.Reset + ") \t" + color.Cyan + "OS:\t" + color.Reset +
		getOS())
}
