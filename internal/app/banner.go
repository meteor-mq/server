// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package app

import "github.com/fatih/color"

const (
	// Verison app version
	Verison   = "0.0.1"
	bannerStr = "\n" +
		"   ___         _             ____ \n" +
		"  / __\\_ _ ___| |_  /\\/\\    /___ \\\n" +
		" / _\\/ _` / __| __|/    \\  //  / /\n" +
		"/ / | (_| \\__ \\ |_/ /\\/\\ \\/ \\_/ / \n" +
		"\\/   \\__,_|___/\\__\\/    \\/\\___,_\\ " + "%s"
)

// EchoInfo output Application info
func EchoInfo() {
	color.Red(bannerStr, colors(color.BgBlue, color.FgBlack).Sprintf("V: %s", Verison)+"\n")
	greenInfo("Using config file ~/.config/code-server/config.yaml")
}

// info  Wrote default config file to ~/.config/code-server/config.yaml
// info  Using config file ~/.config/code-server/config.yaml
// info  Using user-data-dir ~/.local/share/code-server
// info  code-server 3.3.1 6f1309795e1cb930edba68cdc7c3dcaa01da0ab3
// info  HTTP server listening on http://127.0.0.1:8080
// info      - Using password from ~/.config/code-server/config.yaml
// info      - To disable use `--auth none`
// info    - Not serving HTTPS

func greenInfo(item string) {
	colors(color.BgGreen, color.FgBlack).Print("「 INFO 」")
	color.White(item)
}
func colors(back, fg color.Attribute) *color.Color {
	return color.New(back).Add(fg)
}
