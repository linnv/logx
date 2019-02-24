package logx

import (
	"os"
	"testing"
)

func TestLogx_Debugln(t *testing.T) {
	type fields struct {
		underFile *os.File
		toFile    bool
	}
	type args struct {
		format    string
		paramters []interface{}
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
			l := &Logx{
				underFile: tt.fields.underFile,
				toFile:    tt.fields.toFile,
			}
			l.Debugf(tt.args.format, tt.args.paramters...)
			//@TODO fill the boring and unskilled testcase
			// l.Debug(tt.args.format, tt.args.paramters...)
		})
	}
}

func TestLogx_Warn(t *testing.T) {
	type fields struct {
		underFile *os.File
		toFile    bool
	}
	type args struct {
		format    string
		paramters []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"normal",
			fields{nil, true},
			args{"string %v %s %d %.2f\n", []interface{}{"something ", "just warning string", 11000, 1.29447383}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLogx()
			// l.Warn(tt.args.format, tt.args.paramters...)
			// l.Fatalln(tt.args.format, tt.args.paramters...)
			l.Errorf(tt.args.format, tt.args.paramters...)
			l.GracefullyExit()
		})
	}
}
