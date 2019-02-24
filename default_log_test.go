package logx

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
		{"normal", args{"demo string %v %s %d %.2f\n", []interface{}{"something ", "pubic default logger just warning string\n", 11000, 1.29447383}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Errorf(tt.args.format, tt.args.paramters...)
		})
	}
}
