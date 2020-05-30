// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package server

import (
	"os"

	klog "github.com/higker/logker"
)

var Log klog.Logger

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

func init() {
	// Custom logging message template
	format := "{level} - 时间 {time}  - 位置 {position} - 消息 {message}"
	logger, err := klog.NewClog(klog.DEBUG, klog.Shanghai, format, klog.InitAsync(klog.Qs3w))
	if err != nil {
		panic("init logger failed !!")
		os.Exit(1)
	}
	Log = logger
}
