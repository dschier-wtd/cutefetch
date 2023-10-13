package system

import (
	"os"
	"os/user"
	"strings"
)

func GetUsername() (string, error) {
	username, err := user.Current()
	if err != nil {
		return "", err
	}
	return username.Username, nil
}

func GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostname, nil
}

func GetOS() (string, error) {
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
