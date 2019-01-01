package logx

import (
	"fmt"
	"os"
)

//default logger
var Log *Logx

func DisableBuffer(disable bool) {
	Log.DisableBuffer(disable)
}

func Fatalln(paramters ...interface{}) {
	Log.output(calldepth, outputLevelFatal, logRed(fmt.Sprintln(paramters...)))
	Log.GracefullyExit()
	os.Exit(1)
}

func Fatalf(format string, paramters ...interface{}) {
	Log.output(calldepth, outputLevelFatal, logRed(fmt.Sprintf(format, paramters...)))
	Log.GracefullyExit()
	os.Exit(1)
}

func Errorln(paramters ...interface{}) {
	Log.output(calldepth, outputLevelError, logRed(fmt.Sprintln(paramters...)))
}

func Errorf(format string, paramters ...interface{}) {
	Log.output(calldepth, outputLevelError, logRed(fmt.Sprintf(format, paramters...)))
}

func Debugln(paramters ...interface{}) {
	Log.output(calldepth, outputLevelDebug, fmt.Sprintln(paramters...))
}

func Debugf(format string, paramters ...interface{}) {
	//@TODO benchmark convertion efficency
	Log.output(calldepth, outputLevelDebug, fmt.Sprintf(format, paramters...))
}

func Warnln(paramters ...interface{}) {
	Log.output(calldepth, outputLevelWarn, logYellow(fmt.Sprintln(paramters...)))
}

func Warnf(format string, paramters ...interface{}) {
	//@TODO benchmark convertion efficency
	Log.output(calldepth, outputLevelWarn, logYellow(fmt.Sprintf(format, paramters...)))
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

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func PanicErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
