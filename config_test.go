package logx

import (
	"reflect"
	"testing"
)

func TestLoadConfJson(t *testing.T) {
	const testLog = "/Users/Jialin/golang/src/github.com/linnv/test.log"
	const bs = `{
		  "DisableBuffer": false,
		  "Maxbuffer": "3MB",
		  "ToDifferentFile": true,
		  "DevMode": true,
		  "FilePath":"/Users/Jialin/golang/src/github.com/linnv/test.log"
		}`
	r := &LogxConfig{
		Maxbuffer:       "3MB",
		MaxbufferInt:    2 * (1 << 20), //maxbuffer
		DisableBuffer:   false,
		ToDifferentFile: true,
		DevMode:         true,
		FilePath:        testLog,
	}
	const bsTwo = `{
		  "DisableBuffer": false,
		  "Maxbuffer": "1MB",
		  "ToDifferentFile": true,
		  "FilePath": "/Users/Jialin/golang/src/github.com/linnv/test.log"
		}`
	rTwo := &LogxConfig{
		Maxbuffer:       "1MB",
		MaxbufferInt:    1 << 20, //maxbuffer
		DisableBuffer:   false,
		ToDifferentFile: true,
		DevMode:         true,
		FilePath:        testLog,
	}
	type args struct {
		conf []byte
	}
	tests := []struct {
		name string
		args args
		want *LogxConfig
	}{
		{"normal", args{[]byte(bs)}, r},
		{"normal", args{[]byte(bsTwo)}, rTwo},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadConfJson(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unitParse(t *testing.T) {
	type args struct {
		one string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// {"normal", args{"1byte"}, 1},
	// {"normal", args{"2kb"}, 2 * (1 << 10)},
	// {"normal", args{"1mb"}, 1 << 20},
	// {"normal", args{"1"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unitParse(tt.args.one); got != tt.want {
				t.Errorf("unitParse() = %v, want %v", got, tt.want)
			}
		})
	}
}
