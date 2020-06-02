package server

import (
	"math/rand"
	"testing"
	"time"
)

func TestKrand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// go test -v -run=TestKrand
	t.Log("num:   " + string(Krand(16, RuleKindNumber)))
	t.Log("lower: " + string(Krand(16, RuleKindLower)))
	t.Log("upper: " + string(Krand(16, RuleKindUpper)))
	t.Log("all:   " + string(Krand(16, RuleKindAll)))
	t.Log(rand.Intn(3))
	t.Log(byte(49))
	// type args struct {
	// 	size int
	// 	kind int
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want []byte
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := Krand(tt.args.size, tt.args.kind); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("Krand() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
