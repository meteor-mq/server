// Copyright (c) 2020  Shuo Ding
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// License that can be found in the LICENSE file.

package queue

// Queue interface
type Queue interface {
	Writer([]byte) error
	Read([]byte) error
}
