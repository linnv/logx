package logx

import "fmt"

//default logger
var Log *Logx

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
}

func EnableDevMode(enabled bool) {
	Log.EnableDevMode(enabled)
}

func LogConfigure() {
	Log.LogConfigure()
}

func Sync() {
	Log.Sync()
}
