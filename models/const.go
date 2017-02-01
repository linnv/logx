package models

import "os"

const calldepth = 2

const (
	outputLevelDebug = byte(iota)
	outputLevelWarn
	outputLevelError
	outputLevelFatal
)

var prefix = [...]string{
	outputLevelDebug: "[debug]",
	outputLevelWarn:  "[warn]",
	outputLevelError: "[error]",
	outputLevelFatal: "[fatal]",
}

var gopath string

func GetEnvs() string {
	gopath = os.Getenv("GOPATH")
	return gopath
}
