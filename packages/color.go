package color

import "runtime"

var Reset = "\033[0;0m"
var Red = "\033[1;31m"
var Green = "\033[1;32m"
var Yellow = "\033[1;33m"
var Blue = "\033[1;34m"
var Purple = "\033[1;35m"
var Cyan = "\033[1;36m"
var Gray = "\033[1;37m"
var White = "\033[1;97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}
