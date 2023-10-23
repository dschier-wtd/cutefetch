package main

import (
	"fmt"

	"github.com/dschier-wtd/cutefetch/internal/color"
	"github.com/dschier-wtd/cutefetch/internal/system"
	"github.com/dschier-wtd/cutefetch/internal/time"
)

func main() {
	username, err := system.GetUsername()
	if err != nil {
		fmt.Println(err)
	}

	hostname, err := system.GetHostname()
	if err != nil {
		fmt.Println(err)
	}

	currentTime := time.GetCurrentTime()

	os, err := system.GetOS()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n")
	fmt.Printf("(\\ /)\t\t%sUser:\t%s%s@%s\n", color.Red, color.Reset, username, hostname)
	fmt.Printf("( · ·)\t\t%sTime:\t%s%s\n", color.Green, color.Reset, currentTime)
	fmt.Printf("c(%s\"%s)(%s\"%s)\t\t%sOS:\t%s%s\n", color.Red, color.Reset, color.Red, color.Reset, color.Cyan, color.Reset, os)
	fmt.Printf("\n")
}
