package main

import (
	"fmt"
	"log"
	"os"
	"upm/ui"
)

var (
	gitHash string
	version string
)

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("UnityPackageManager version %s, build %s\n", version, gitHash)
		return
	}
	if len(args) > 1 {
		info := OpenPackage(args[1])
		if info != nil {
			ui.Window(info)
		}
	}
}
