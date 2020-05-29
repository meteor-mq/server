// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package server

var (
	// DefaultAllow default allow address
	DefaultAllow = make([]string, 0, 8)
	// DefaultPassword default password
	DefaultPassword = func() string {
		return "123456"
	}()
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
