package server

import (
	"math/rand"
	"time"
)

// Random Password
func randomString() []byte {
	rand.Seed(time.Now().UnixNano())
	// set string length
	bytes := make([]byte, 0, 26)
	for i := 0; i < cap(bytes); i++ {
		bytes = append(bytes, byte(rand.Intn(26)+65))
	}
	return bytes
}
