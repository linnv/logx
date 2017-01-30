package models

import "testing"

func TestConst(t *testing.T) {
	type args struct {
		n byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal", args{byte(0x01)}, outputDirection[outputDirectionDebug]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Const(tt.args.n); got != tt.want {
				t.Errorf("Const() = %v, want %v", got, tt.want)
			}
		})
	}
}
