// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package main

import (
	"os"

	server "github.com/fast-mq/server/core"
	"github.com/urfave/cli"
)

// 2020年06月02日内容
// 包   守护进程
// 马士兵  写完这个
// 优化代码
func main() {
	// 实例化cli
	app := cli.NewApp()
	// Name可以设定应用的名字
	app.Name = "Fast Message Queue"
	app.Usage = "is light weight message queue middleware 🚀."
	// Version可以设定应用的版本号
	app.Version = server.Verison
	// Commands 命令程序
	app.Commands = []cli.Command{
		{
			// 命令的名字
			Name: "start",
			// 命令的缩写，就是不输入language只输入lang也可以调用命令
			Aliases: []string{"up"},
			// 命令的用法注释，这里会在输入 程序名 -help的时候显示命令的使用方法
			Usage: "FastMQ program start command.",
			// 命令的处理函数
			Action: func(c *cli.Context) error {
				//c.Args().First()
				server.EchoInfo()
				return nil
			},
		},
		{
			// 命令的名字
			Name: "stop",
			// 命令的缩写，就是不输入language只输入lang也可以调用命令
			Aliases: []string{"down"},
			// 命令的用法注释，这里会在输入 程序名 -help的时候显示命令的使用方法
			Usage: "FastMQ program stop command.",
			// 命令的处理函数
			Action: func(c *cli.Context) error {
				//c.Args().First()
				server.EchoInfo()
				return nil
			},
		},
	}
	// 接受os.Args并启动程序
	app.Run(os.Args)
	//fmt.Println(language)
}
