package utils

import (
	"log"
	"path/filepath"

	"github.com/gen2brain/beeep"
)

const (
	AppID = "UnityPackageManager"
)

var iconfile string

func init() {
	iconfile, _ = filepath.Abs("resource/icon/default.png")
}

func Toast(msg string, title ...string) {
	t := AppID
	if len(title) > 0 {
		t = title[0]
	}
	if err := beeep.Notify(t, msg, iconfile); err != nil {
		log.Println(err.Error())
	}
}
