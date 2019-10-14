package logx

import (
	"os"
	"testing"
)

func Test_checkDirAvailable(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"normal", args{"./version"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkDirAvailable(tt.args.filepath); (err != nil) != tt.wantErr {
				t.Errorf("checkDirAvailable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkFileAvailable(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr bool
	}{
		{"normal", args{"./version/doc.go"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkFileAvailable(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkFileAvailable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			_ = got
		})
	}
}
