package logx

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/linnv/bufferlog"
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
	Infoln("abcd")

	EnableDevMode(true)
	Debugln("abcd")
	SetLevel(OutputLevelInfo)
	time.Sleep(time.Millisecond)
	Infoln("info abcd")
	SetLevel(OutputLevelDebug)
	Debugln("debug abcd")
	Infoln("info abcd")
	Flush()
}

type discardCloser struct{}

func (discardCloser) Write(p []byte) (int, error) {
	return len(p), nil
}

func (discardCloser) Close() error {
	return nil
}

func DiscardCloser() io.WriteCloser {
	return discardCloser{}
}

// with funcName enable
// 1782848               665.9 ns/op
//
// with funcName disable
// 1843812               628.0 ns/op
func BenchmarkLog(b *testing.B) {
	exit := make(chan struct{})
	logWriter := bufferlog.NewBufferLog(3*1024, time.Second*2, exit, discardCloser{})
	logger := NewLogx(logWriter)
	for i := 0; i < b.N; i++ {
		logger.Debugln("juset demo")
	}
}
