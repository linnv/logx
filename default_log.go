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

func Debugln(parameters ...interface{}) {
	Log.output(calldepth, outputLevelDebug, logBlue(fmt.Sprintln(parameters...)))
}

func Debugfln(format string, parameters ...interface{}) {
	//@TODO benchmark conversion efficiency
	Log.output(calldepth, outputLevelDebug, logBlue(fmt.Sprintf(format+"\n", parameters...)))
}

func Debugf(format string, parameters ...interface{}) {
	Log.output(calldepth, outputLevelDebug, logBlue(fmt.Sprintf(format, parameters...)))
}

func Warnln(parameters ...interface{}) {
	Log.output(calldepth, outputLevelWarn, logYellow(fmt.Sprintln(parameters...)))
}

func Warnfln(format string, parameters ...interface{}) {
	Log.output(calldepth, outputLevelWarn, logYellow(fmt.Sprintf(format+"\n", parameters...)))
}

func Warnf(format string, parameters ...interface{}) {
	Log.output(calldepth, outputLevelWarn, logYellow(fmt.Sprintf(format, parameters...)))
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
