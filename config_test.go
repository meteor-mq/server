// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package server

import (
	"reflect"
	"testing"
)

// TestNewDefaultConfig  go test -v -run=TestNewDefaultConfig/error
func TestNewDefaultConfig(t *testing.T) {

	tests := []struct {
		name string
		want *Option
	}{
		//cannot use "c1" (type string) as type struct { name string; want *Option } in slice literal
		{
			name: "success",
			want: NewDefaultConfig(),
		},
		{
			name: "error",
			want: errorConfig(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
func errorConfig() *Option {
	return &Option{
		Address:  "127.0.0.1",
		Port:     DefaultPort,
		AllowIP:  DefaultAllow,
		Password: "0000000",
	}
}
