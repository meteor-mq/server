// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package server

import (
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	// DefaultAllow default allow address
	DefaultAllow = make([]string, 0, 8)
	// DefaultPassword default password
	DefaultPassword = string(Krand(16, RuleKindAll))
)

const (
	// DefaultPort default port
	DefaultPort = "9598"
	// DefaultAddress  default address
	DefaultAddress = "127.0.0.1"
)

// Option is create FastMQ server parameter
type Option struct {
	// Bind IP
	Address string
	// Bind Port 0-65535
	Port string
	// Allow access IP
	AllowIP []string
	// Auth password
	Password string
}

// DefaultConfig genreate server option default cofing
func DefaultConfig() *Option {
	return &Option{
		Address:  "127.0.0.1",
		Port:     DefaultPort,
		AllowIP:  DefaultAllow,
		Password: DefaultPassword,
	}
}

// NewConfig customize server config option
// Example Code :
// op := mq.NewConfig()
// op.Password = "XXXXXXX"
func NewConfig() *Option {
	return &Option{}
}

// loadConfigurationFile
func loadConfig(path string) Option {
	cfg, err := ini.Load(path)
	if err != nil {
		Log.Error("loading config file error:%s", err)
		os.Exit(1)
	}
	return Option{
		Address: cfg.Section("option").Key("BindIP").String(),
		Port:    cfg.Section("option").Key("BindPort").String(),
		AllowIP: func() []string {
			all := cfg.Section("option").Key("AllowIP").String()
			return strings.Split(all, ",")
		}(),
		Password: cfg.Section("option").Key("Password").String(),
	}
}
