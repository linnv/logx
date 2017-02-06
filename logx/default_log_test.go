package model

import "testing"

func TestErrorln(t *testing.T) {
	type args struct {
		format    string
		paramters []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"normal", args{"string %v %s %d %.2f", []interface{}{"something ", "pubic default logger just warning string", 11000, 1.29447383}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorln(tt.args.format, tt.args.paramters...)
		})
	}
}
