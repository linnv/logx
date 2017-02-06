package logx

import (
	"fmt"
	"sync"
)

//default logger
var Log *Logx

var once sync.Once

func initDefaultLog() {
	f := func() {
		_, mode := GetFlags()
		if mode {
			Log = NewLogxFile()
		} else {
			Log = NewLogx()
		}
		return
	}
	once.Do(f)
	return
}

func Errorln(format string, paramters ...interface{}) {
	Log.output(calldepth, outputLevelError, fmt.Sprintf(format+"\n", paramters...))
}

func Error(format string, paramters ...interface{}) {
	Log.output(calldepth, outputLevelError, fmt.Sprintf(format, paramters...))
}

func Debugln(format string, paramters ...interface{}) {
	Log.output(calldepth, outputLevelDebug, fmt.Sprintf(format+"\n", paramters...))
}

func Debug(format string, paramters ...interface{}) {
	//@TODO benchmark convertion efficency
	Log.output(calldepth, outputLevelDebug, fmt.Sprintf(format, paramters...))
	Log.LogConfigure()
}
