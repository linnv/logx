package models

import (
	"fmt"
	"os"
)

type Logx struct {
	underFile *os.File `json:"underFile"`
	// outputLevel     int    `json:"outputLevel"`
	outputDirection byte
}

//Debugln
func (l *Logx) Debugln(format string, paramters ...interface{}) {
	const prefix = "[debug]"
	println(prefix, fmt.Sprintf(format+"\n", paramters...))
}

//Debug
func (l *Logx) Debug(format string, paramters ...interface{}) {
	const prefix = "[debug]"
	println(prefix, fmt.Sprintf(format, paramters...))
}

//default logger
var Log = NewLogx()

//@TODO from confireguration
const useFile = true

func NewLogx() *Logx {
	return &Logx{
		underFile:       nil,
		outputDirection: outputLevelDebug,
	}
}
