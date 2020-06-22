package utils

import (
	"path/filepath"

	"github.com/go-toast/toast"
)

const (
	AppID = "UnityPackageManager"
)

var iconfile string

func init() {
	iconfile, _ = filepath.Abs("icon/default.png")
}

func Toast(msg string, title ...string) {
	t := AppID
	if len(title) > 0 {
		t = title[0]
	}
	notification := toast.Notification{
		AppID:   AppID,
		Title:   t,
		Message: msg,
		Icon:    iconfile,
	}
	_ = notification.Push()
}
