// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

// ###############################$
// #########################&%$###$
// #################@%%##&|!|&####$
// ##############$|!!%$|!!!|@#####$
// ###########%!!!!!!!!;;!%#######$
// ########%!!!;;;!!;;;;!$########$
// #####&|!!;;;;;;;;;;!!||!$######$
// ###&|!!;;;;;;;;;;;!!!!!&#######$
// ##&!!;;;;;;;;;;;;;;;!|@########$
// ##%!;;;:'``':;;;;;;!%@#########$
// ##$!;;;:`..`:;;;;!!$###########$
// ##@|!;;;::::;;;!!|@############$
// ####%!!;;;;;;!!|&##############$
// ######&%!!!!|%@################$
// ###############################$
// LOGO IMAGE : https://avatars0.githubusercontent.com/u/48612456?s=200&v=4

package app

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/fast-mq/server/internal/conf"
	"github.com/higker/logker"
)

var (
	// data buffer
	buffers [1024]byte
)

// MQServer is message queue.
type MQServer struct {
	// Bind IP
	Address string
	// Bind Port 0-65535
	Port uint16
	// Allow access IP
	AllowIP []string
	// Auth password
	Password string
}

// NewMQServer create server
func NewMQServer(option *conf.Option) *MQServer {
	return &MQServer{
		Address: option.Address,
		Port: func() uint16 {
			parseUint, err := strconv.ParseUint(option.Port, 10, 16)
			if err != nil {
				logker.Error("bind port error:", err)
				//os.Exit(1)
			}
			return uint16(parseUint)
		}(),
		AllowIP:  option.AllowIP,
		Password: option.Password,
	}
}

// Start START MqServer
func (ms *MQServer) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ms.Address, ms.Port))
	if err != nil {
		fmt.Println("create tcp server error:%v", err)
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("tcp accept error:", err)
		}
		go ms.handleConn(accept)
	}
}

func (ms *MQServer) handleConn(con net.Conn) {
	for {
		ip := strings.Split(con.RemoteAddr().String(), ":")[0]
		logker.Warning("ip = %s", ip)
		n, _ := con.Read(buffers[:])
		logker.Warning("%v", string(buffers[:n]))
	}
}
