package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"

	color "github.com/dschier-wtd/cutefetch/internal"
)

func getUsername() (string, error) {
	username, err := user.Current()
	if err != nil {
		return "", err
	}
	return username.Username, nil
}

func getHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostname, nil
}

func getCurrentTime() string {
	currentTime := time.Now().Format("02.01.2006 - 15:04:05")
	return currentTime
}

func getOS() (string, error) {
	fileContent, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(fileContent), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				osName := strings.Trim(parts[1], `"`)

				return osName, nil
			}
		}
	}

	return "Unknown", nil
}

func main() {
	username, err := getUsername()
	if err != nil {
		fmt.Println(err)
	}

	hostname, err := getHostname()
	if err != nil {
		fmt.Println(err)
	}

	currentTime := getCurrentTime()

	os, err := getOS()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Printf("(\\ /)\t\t%sUser:\t%s%s@%s\n", color.Red, color.Reset, username, hostname)
	fmt.Printf("( · ·)\t\t%sTime:\t%s%s\n", color.Green, color.Reset, currentTime)
	fmt.Printf("c(%s\"%s%s%s%s)\t%sOS:\t%s%s\n", color.Red, color.Reset, "\"", color.Red, color.Reset, color.Cyan, color.Reset, os)
}
