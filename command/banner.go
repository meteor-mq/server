// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package command

import (
	"github.com/fatih/color"
)

const (
	verison   = "0.0.1"
	bannerStr = "\n" +
		"   ___         _             ____ \n" +
		"  / __\\_ _ ___| |_  /\\/\\    /___ \\\n" +
		" / _\\/ _` / __| __|/    \\  //  / /\n" +
		"/ / | (_| \\__ \\ |_/ /\\/\\ \\/ \\_/ / \n" +
		"\\/   \\__,_|___/\\__\\/    \\/\\___,_\\ " + "%s"
)

// echoInfo output Application info
func echoInfo() {
	bg := color.New(color.BgBlue)
	bg.Add(color.FgBlack)
	color.Red(bannerStr, bg.Sprintf("V: %s", verison)+"\n")
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
	bg := color.New(color.BgGreen)
	bg.Add(color.FgBlack)
	bg.Print("「 INFO 」")
	color.White(item)
}
