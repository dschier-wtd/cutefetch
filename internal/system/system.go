package system

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func GetUsername() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.Username, nil
}

func GetHostname() (string, error) {
	hostname, err := os.Hostname()
	return hostname, err
}

func GetOS() (string, error) {
	fileContent, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(fileContent), "\n")
	for _, line := range lines {
		osName := strings.TrimPrefix(line, "PRETTY_NAME=")
		osName = strings.Trim(osName, `"`)
		return osName, nil
	}

	return "Unknown", nil
}

func GetShell() (string, error) {
	shellPath := os.Getenv("SHELL")

	if shellPath == "" {
		return "Unknown", nil
	}

	shell := filepath.Base(shellPath)

	return shell, nil
}
