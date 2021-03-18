package logx

import (
	"fmt"
	"testing"
	"time"
)

func TestErrorln(t *testing.T) {
	onetime := time.Now().Format("20060102 15:04:05.000")
	fmt.Printf("onetime: %s\n", onetime)
	type args struct {
		format     string
		parameters []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"normal", args{"demo string %v %s %d %.2f\n", []interface{}{"something ", "pubic default logger just warning string\n", 11000, 1.29447383}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorf(tt.args.format, tt.args.parameters...)
		})
	}

	EnableDevMode(false)
	Debugln("abcd")
	Debugln("abcd")

	EnableDevMode(true)
	Debugln("abcd")
	time.Sleep(time.Millisecond)
	Debugln("abcd")
	Flush()
}
