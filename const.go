package logx

import (
	"errors"
)

const Calldepth = 2

const (
	OutputLevelDebug = byte(iota)
	OutputLevelInfo
	OutputLevelWarn
	OutputLevelError
	OutputLevelFatal
)

var prefix = [...][]byte{
	OutputLevelDebug: []byte("[debug]"),
	OutputLevelInfo:  []byte("[info]"),
	OutputLevelWarn:  []byte("[warn]"),
	OutputLevelError: []byte("[error]"),
	OutputLevelFatal: []byte("[fatal]"),
}

var ErrTooLarge = errors.New("too large slice to allocate")
