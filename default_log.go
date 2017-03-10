package logx

import "fmt"

//default logger
var Log *Logx

func DisableBuffer(disable bool) {
	Log.DisableBuffer(disable)
}

func Fatalln(paramters ...interface{}) {
	Log.output(calldepth, outputLevelFatal, fmt.Sprintln(paramters...))
}

func Fatalf(format string, paramters ...interface{}) {
	Log.output(calldepth, outputLevelFatal, fmt.Sprintf(format, paramters...))
}

func Errorln(paramters ...interface{}) {
	Log.output(calldepth, outputLevelError, fmt.Sprintln(paramters...))
}

func Errorf(format string, paramters ...interface{}) {
	Log.output(calldepth, outputLevelError, fmt.Sprintf(format, paramters...))
}

func Debugln(paramters ...interface{}) {
	Log.output(calldepth, outputLevelDebug, fmt.Sprintln(paramters...))
}

func Debugf(format string, paramters ...interface{}) {
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
