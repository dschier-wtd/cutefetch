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

	currentTime, err := time.GetCurrentTime()
	if err != nil {
		fmt.Println(err)
	}

	os, err := system.GetOS()
	if err != nil {
		fmt.Println(err)
	}

	shell, err := system.GetShell()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n")
	fmt.Printf("\t\t\t%s%s%s @ %s%s%s\n", color.Green, username, color.Reset, color.Cyan, hostname, color.Reset)
	fmt.Printf("(\\ /)\t\t%sTime:\t%s%s\n", color.Red, color.Reset, currentTime)
	fmt.Printf("( · ·)\t\t%sShell:\t%s%s\n", color.Purple, color.Reset, shell)
	fmt.Printf("c(%s\"%s)(%s\"%s)\t\t%sOS:\t%s%s\n", color.Red, color.Reset, color.Red, color.Reset, color.Yellow, color.Reset, os)
}
