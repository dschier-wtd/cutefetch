package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"

	color "github.com/dschier-wtd/cutefetch/packages"
)

func getUser() (cuser string) {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Print("Error")
	}
	cuser = currentUser.Username
	return
}

func getHostname() (chost string) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Print("Error")
	}
	chost = hostname
	return
}

func getTime() (ctime string) {
	currentTime := time.Now()
	formatedTime := currentTime.Format("02.01.2006 - 15:04:05")
	ctime = formatedTime
	return
}

func getOS() (cdist string) {
	content, err := os.ReadFile("/etc/os-release")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				distribution := strings.Trim(parts[1], `"`)

				cdist = distribution
				return
			}
		}
	}

	cdist = "Unknown"
	return
}

func main() {

	fmt.Println()
	fmt.Println(" (\\ /) \t\t" + color.Red + "User:\t" + color.Reset +
		getUser() + "@" + getHostname())
	fmt.Println(" ( · ·) \t" + color.Green + "Time:\t" + color.Reset +
		getTime())
	fmt.Println(" c(" + color.Red + "\"" + color.Reset + ")(" + color.Red +
		"\"" + color.Reset + ") \t" + color.Cyan + "OS:\t" + color.Reset +
		getOS())
}
