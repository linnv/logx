package models

import "testing"

func TestGetEnvs(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// {"normal", "/Users/Jialin/golang"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvs(); got != tt.want {
				t.Errorf("GetEnvs() = %v, want %v", got, tt.want)
			}
		})
	}
}
