// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

//                    ::
//                   :;J7, :,                        ::;7:
//                   ,ivYi, ,                       ;LLLFS:
//                   :iv7Yi                       :7ri;j5PL
//                  ,:ivYLvr                    ,ivrrirrY2X,
//                  :;r@Wwz.7r:                :ivu@kexianli.
//                 :iL7::,:::iiirii:ii;::::,,irvF7rvvLujL7ur
//                ri::,:,::i:iiiiiii:i:irrv177JX7rYXqZEkvv17
//             ;i:, , ::::iirrririi:i:::iiir2XXvii;L8OGJr71i
//           :,, ,,:   ,::ir@mingyi.irii:i:::j1jri7ZBOS7ivv,
//              ,::,    ::rv77iiiriii:iii:i::,rvLq@huhao.Li
//          ,,      ,, ,:ir7ir::,:::i;ir:::i:i::rSGGYri712:
//        :::  ,v7r:: ::rrv77:, ,, ,:i7rrii:::::, ir7ri7Lri
//       ,     2OBBOi,iiir;r::        ,irriiii::,, ,iv7Luur:
//     ,,     i78MBBi,:,:::,:,  :7FSL: ,iriii:::i::,,:rLqXv::
//     :      iuMMP: :,:::,:ii;2GY7OBB0viiii:i:iii:i:::iJqL;::
//    ,     ::::i   ,,,,, ::LuBBu BBBBBErii:i:i:i:i:i:i:r77ii
//   ,       :       , ,,:::rruBZ1MBBqi, :,,,:::,::::::iiriri:
//  ,               ,,,,::::i:  @arqiao.       ,:,, ,:::ii;i7:
// :,       rjujLYLi   ,,:::::,:::::::::,,   ,:i,:,,,,,::i:iii
// ::      BBBBBBBBB0,    ,,::: , ,:::::: ,      ,,,, ,,:::::::
// i,  ,  ,8BMMBBBBBBi     ,,:,,     ,,, , ,   , , , :,::ii::i::
// :      iZMOMOMBBM2::::::::::,,,,     ,,,,,,:,,,::::i:irr:i:::,
// i   ,,:;u0MBMOG1L:::i::::::  ,,,::,   ,,, ::::::i:i:iirii:i:i:
// :    ,iuUuuXUkFu7i:iii:i:::, :,:,: ::::::::i:i:::::iirr7iiri::
// :     :rk@Yizero.i:::::, ,:ii:::::::i:::::i::,::::iirrriiiri::,
//  :      5BMBBBBBBSr:,::rv2kuii:::iii::,:i:,, , ,,:,:i@petermu.,
//       , :r50EZ8MBBBBGOBBBZP7::::i::,:::::,: :,:,::i;rrririiii::
//           :jujYY7LS0ujJL7r::,::i::,::::::::::::::iirirrrrrrr:ii:
//        ,:  :@kevensun.:,:,,,::::i:i:::::,,::::::iir;ii;7v77;ii;i,
//        ,,,     ,,:,::::::i:iiiii:i::::,, ::::iiiir@xingjief.r;7:i,
//     , , ,,,:,,::::::::iiiiiiiiii:,:,:::::::::iiir;ri7vL77rrirri::
//      :,, , ::::::::i:::i:::i:i::,,,,,:,::i:i:::iir;@Secbone.ii:::

package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/fast-mq/server/internal/app"
	"github.com/fast-mq/server/internal/conf"
	"github.com/higker/logker"
	"github.com/urfave/cli"
)

var (
	// ErrStartUp START UP ERROR
	ErrStartUp = errors.New("daemon start error")
	// whether enable  backgrund running
	isBackgrund bool
	processID   int = -1
	pidPath         = "fast_mq.pid"
)

var (
	// Start program start command.
	Start = cli.Command{
		// Command name
		Name: "start",
		// Command abbreviation
		Aliases: []string{"up"},
		// Command usage notes, here will show how to use the command when you enter the program name -help
		Usage: "FastMQ program start command✅.\n",
		// Command processing function
		Action: func(c *cli.Context) error {
			app.EchoInfo()
			if isBackgrund {
				logker.Info("%s", "FastMQ is enable daemon mode running.")
				daemon()
				logker.Info("FastMQ PID: %d running", processID)
				// 延迟3秒钟因为exec在调用bash需要一点时间
				time.Sleep(time.Second * 3)
				return nil
			}
			logker.Info("%s", "FastMQ serve start up.")
			mq := app.NewMQServer(conf.LoadOption())
			if err := mq.Start(); err != nil {
				return err
			}
			//Daemon(0, 0)
			return nil
		},
	}
	// Stop program stop command.
	Stop = cli.Command{

		Name: "stop",

		Aliases: []string{"down"},

		Usage: "FastMQ program stop command❎.\n",

		Action: func(c *cli.Context) error {
			// STOP
			kill()
			return nil
		},
	}
)

func init() {
	Start.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "daemon, d",
			Usage:       "Whether the program is running in the background.",
			Destination: &isBackgrund,
		},
	}
}

func daemon() {
	cmd := exec.Command("bash", "-c", "./fast_mq up")
	// 将其他命令传入生成出的进程
	// cmd.Stdin = os.Stdin // 给新进程设置文件描述符，可以重定向到文件中
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		panic(err)
		//os.Exit(1)
	}
	if pid, _ := ioutil.ReadFile(pidPath); string(pid) != "" {
		logker.Warning("%s", "FastMQ is running.")
		time.Sleep(time.Second * 2)
		os.Exit(0)
	}
	ioutil.WriteFile(pidPath, []byte(fmt.Sprintf("%d", cmd.Process.Pid)), 0666)
	isBackgrund = false
	processID = cmd.Process.Pid
}
func kill() {
	pid, err := ioutil.ReadFile(pidPath)
	if err == nil {
		exec.Command("bash", "-c", "kill -9 "+string(pid)).Start()
		logker.Warning("%s", "FastMQ close successful.👋")
		ioutil.WriteFile(pidPath, []byte(nil), 0666)
		time.Sleep(5 * time.Second)
		return
	}
	panic("FastMQ pid file not fond.！！！")
}

// Daemon 守护进程
func Daemon(nochdir, noclose int) (int, error) {
	// already a daemon
	if syscall.Getppid() == 1 {
		/* Change the file mode mask */
		syscall.Umask(0)

		if nochdir == 0 {
			os.Chdir("/")
		}

		return 0, nil
	}

	files := make([]*os.File, 3, 6)
	if noclose == 0 {
		nullDev, err := os.OpenFile("/dev/null", 0, 0)
		if err != nil {
			return 1, err
		}
		files[0], files[1], files[2] = nullDev, nullDev, nullDev
	} else {
		files[0], files[1], files[2] = os.Stdin, os.Stdout, os.Stderr
	}

	dir, _ := os.Getwd()
	sysattrs := syscall.SysProcAttr{Setsid: true}
	attrs := os.ProcAttr{Dir: dir, Env: os.Environ(), Files: files, Sys: &sysattrs}

	proc, err := os.StartProcess(os.Args[0], os.Args, &attrs)
	if err != nil {
		return -1, fmt.Errorf("can't create process %s: %s", os.Args[0], err)
	}
	proc.Release()
	os.Exit(0)

	return 0, nil
}
