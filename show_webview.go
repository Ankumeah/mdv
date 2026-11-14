package main

import (
	"github.com/Ankumeah/webview_go"
)

func showWebview(filename string, html string) {
	view := webview.New(false)
	defer view.Destroy()

	view.SetTitle(filename)
	view.SetSize(480, 320, webview.HintNone)
	view.SetHtml(html)
	view.Bind(
		"setHTML",
		func(text string) {
			view.SetHtml(styleText(text))
		},
	)

	view.Run()
}
