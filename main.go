//go:generate goversioninfo -o=unity-package-manager.syso
package main

import (
	"log"
	"upm/cmd"
)

func init() {
	log.SetFlags(log.Llongfile)
	cmd.Execute()
}

func main() {}
