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

	containerContext, err := system.GetContainerContext()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n")
	fmt.Printf("\t\t%s%s%s @ %s%s%s\n", color.Green, username, color.Reset, color.Cyan, hostname, color.Reset)
	fmt.Printf(" |\\___/|\t%sTime:%s\t\t%s\n", color.Red, color.Reset, currentTime)
	fmt.Printf("=( ^.^ )=\t%sShell:%s\t\t%s\n", color.Purple, color.Reset, shell)
	fmt.Printf(" (%sö%s) (%sö%s)\t%sOS:%s\t\t%s\n", color.Red, color.Reset, color.Red, color.Reset, color.Yellow, color.Reset, os)
	if containerContext == "container" {
		fmt.Printf("\t\t%sContainer:%s\ttrue\n", color.Blue, color.Reset)
	}
}
