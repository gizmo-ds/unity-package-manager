package ui

import (
	"upm/types"

	"github.com/zserge/webview"
)

func Window(info *types.PackageInfo) {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Hello")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://lumina.moe/")
	w.Run()
}
