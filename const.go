package logx

import (
	"errors"
)

const Calldepth = 2

const (
	OutputLevelDebug = byte(iota)
	OutputLevelWarn
	OutputLevelError
	OutputLevelFatal
)

var prefix = [...][]byte{
	OutputLevelDebug: []byte("[debug]"),
	OutputLevelWarn:  []byte("[warn]"),
	OutputLevelError: []byte("[error]"),
	OutputLevelFatal: []byte("[fatal]"),
}

var ErrTooLarge = errors.New("too large slice to allocate")
