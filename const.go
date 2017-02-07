package logx

import (
	"errors"
	"os"
)

const calldepth = 2

const (
	outputLevelDebug = byte(iota)
	outputLevelWarn
	outputLevelError
	outputLevelFatal
)

var prefix = [...][]byte{
	outputLevelDebug: []byte("[debug]"),
	outputLevelWarn:  []byte("[warn]"),
	outputLevelError: []byte("[error]"),
	outputLevelFatal: []byte("[fatal]"),
}

const maxDefaultBufferSize = 1 << 20 //1MB
var ErrTooLarge = errors.New("too large slice to allocate")

var gopath string

func GetEnvs() string {
	gopath = os.Getenv("GOPATH")
	return gopath
}
