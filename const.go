package logx

import (
	"errors"
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

var ErrTooLarge = errors.New("too large slice to allocate")
