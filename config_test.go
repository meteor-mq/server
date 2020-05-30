// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package server

import (
	"reflect"
	"testing"
)

// TestDefaultConfig  go test -v -run=TestDefaultConfig/error
func TestDefaultConfig(t *testing.T) {

	tests := []struct {
		name string
		want *Option
	}{
		//cannot use "c1" (type string) as type struct { name string; want *Option } in slice literal
		{
			name: "success",
			want: DefaultConfig(),
		},
		{
			name: "error",
			want: errorConfig(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultConfig(); !reflect.DeepEqual(got, tt.want) {
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

func Test_randomString(t *testing.T) {
	t.Log(string(randomString()))
	t.Log(randomString())
	//  go test -v -run=Test_randomString
	// === RUN   Test_randomString
	//     Test_randomString: config_test.go:48: OSHWZSWPRDLYRRJD
	//     Test_randomString: config_test.go:49: [67 83 66 77 73 88 67 86 76 68 67 85 77 68 81 78]
	// --- PASS: Test_randomString (0.00s)
	// PASS
	// ok      github.com/fast-mq/server       0.004s
}

func Test_loadConfig(t *testing.T) {
	t.Log(loadConfig("/Users/ding/Documents/GO_CODE_DEV/src/server/config.ini"))
}
