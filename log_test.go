package logx

import (
	"os"
	"testing"

	"github.com/linnv/bufferlog"
	"go.uber.org/zap"
)

func TestLogx_Debugln(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	type fields struct {
		underFile *os.File
		toFile    bool
	}
	type args struct {
		format     string
		parameters []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"normal",
			fields{nil, false},
			args{"string %v %s %d\n", []interface{}{"something ", "just string", 11000}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLogx(bufferlog.Buffer)
			l.SetZapLogger(logger)
			l.Debugf(tt.args.format, tt.args.parameters...)
			l.GracefullyExit()
		})
	}
}

func TestLogx_Warn(t *testing.T) {
	type fields struct {
		underFile *os.File
		toFile    bool
	}
	type args struct {
		format     string
		parameters []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"normal",
			fields{nil, true},
			args{"demo string %v %s %d %.2f\n", []interface{}{"something ", "just warning string", 11000, 1.29447383}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLogx(bufferlog.Buffer)
			l.Errorf(tt.args.format, tt.args.parameters...)
			l.GracefullyExit()
		})
	}
}
