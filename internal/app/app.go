// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package app

const (
	// Verison app version
	Verison = "0.0.1"
	// Website INFO
	Website = "Website: https://mq.ibyte.me"
	// AppCliName application name
	AppCliName = "Fast Message Queue"
	// AppCliUsage application usage
	AppCliUsage = "is light weight message queue middleware 🚀."
	// ListenPerfix listen Address perfix
	ListenPerfix = "Serve listen:"
	// PasswordPerfix password perfix info
	PasswordPerfix = "Your serve admin Password:"
	// ConfigFile config file path info
	ConfigFile = "Using config file: ~/.config/code-server/config.ini"
)

// Bootstrap FastMQ boot & application entrance
// func Bootstrap() {
// 	appCli := cli.NewApp()

// 	appCli.Name = "Fast Message Queue"

// 	appCli.Usage = "is light weight message queue middleware 🚀."

// 	appCli.Version = app.Verison

// 	appCli.Commands = []cli.Command{
// 		cmd.Start,
// 		cmd.Stop,
// 	}

// 	err := appCli.Run(os.Args)
// 	if err != nil {
// 		logker.Error("application init error:%v", err)
// 	}
// }
