package models

import "os"

const (
	prefixDebug = "[debug]"
	prefixWarn  = "[warn]"
	prefixError = "[error]"
	prefixFatal = "[fatal]"
)

const (
	outputLevelDebug = byte(1 << iota)
	outputLevelWarn
	outputLevelError
	outputLevelFatal
)

var gopath string

func GetEnvs() string {
	gopath = os.Getenv("GOPATH")
	return gopath
}

// var (
// 	outputLevel = [...]string{
// 		outputLevelDebug: "DEBUG",
// 		outputLevelWarn:  "WARN",
// 		outputLevelError: "ERROR",
// 		"FATAL",
// 	}
// )

// func Const(n byte) string {
// 	return outputLevel[n]
// }
