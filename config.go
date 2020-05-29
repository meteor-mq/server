// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package server

var (
	// DefaultAllow default allow address
	DefaultAllow = make([]string, 0, 8)
	// DefaultPassword default password
	DefaultPassword = string(randomString())
)

const (
	// DefaultPort default port
	DefaultPort = 9598
	// DefaultAddress  default address
	DefaultAddress = "127.0.0.1"
)

// Option is create FastMQ server parameter
type Option struct {
	// Bind IP
	Address string
	// Bind Port 0-65535
	Port uint16
	// Allow access IP
	AllowIP []string
	// Auth password
	Password string
}

// NewDefaultConfig genreate server option default cofing
func NewDefaultConfig() *Option {
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
