package models

import (
	"os"
	"testing"
)

func TestLogx_Debugln(t *testing.T) {
	type fields struct {
		underFile       *os.File
		outputDirection byte
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
			fields{nil, outputDirectionDebug},
			args{"string %v %s %d", []interface{}{"somethine ", "just string", 11000}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logx{
				underFile:       tt.fields.underFile,
				outputDirection: tt.fields.outputDirection,
			}
			l.Debugln(tt.args.format, tt.args.paramters...)
			//@TODO fill the boring and unskilled testcase
			l.Debug(tt.args.format, tt.args.paramters...)
		})
	}
}

func TestLogx_Warn(t *testing.T) {
	type fields struct {
		underFile       *os.File
		outputDirection byte
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
			fields{nil, outputDirectionDebug},
			args{"string %v %s %d", []interface{}{"somethine ", "just string", 11000}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLogxFile()
			l.Warn(tt.args.format, tt.args.paramters...)
			l.GracefullyExit()
		})
	}
}
