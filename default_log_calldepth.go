package logx

import (
	"fmt"
	"os"
)

func FatallnWithDepth(addCallDepth int, parameters ...interface{}) {
	Log.Output(Calldepth+addCallDepth, OutputLevelFatal, logRed(fmt.Sprintln(parameters...)))
	Log.GracefullyExit()
	os.Exit(1)
}

func FatalfWithDepth(format string, addCallDepth int, parameters ...interface{}) {
	Log.Output(Calldepth+addCallDepth, OutputLevelFatal, logRed(fmt.Sprintf(format, parameters...)))
	Log.GracefullyExit()
	os.Exit(1)
}

func FatalflnWithDepth(format string, addCallDepth int, parameters ...interface{}) {
	Log.Output(Calldepth+addCallDepth, OutputLevelFatal, logRed(fmt.Sprintf(format+"\n", parameters...)))
	Log.GracefullyExit()
	os.Exit(1)
}

func ErrorlnWithDepth(addCallDepth int, parameters ...interface{}) {
	Log.Output(Calldepth+addCallDepth, OutputLevelError, logRed(fmt.Sprintln(parameters...)))
}

func ErrorfWithDepth(format string, addCallDepth int, parameters ...interface{}) {
	Log.Output(Calldepth+addCallDepth, OutputLevelError, logRed(fmt.Sprintf(format, parameters...)))
}

func ErrorflnWithDepth(format string, addCallDepth int, parameters ...interface{}) {
	Log.Output(Calldepth+addCallDepth, OutputLevelError, logRed(fmt.Sprintf(format+"\n", parameters...)))
}

func PrintlnWithDepth(addCallDepth int, parameters ...interface{}) {
	if !devMode {
		return
	}
	Log.Output(Calldepth+addCallDepth, OutputLevelDebug, logBlue(fmt.Sprintln(parameters...)))
}

func PrintflnWithDepth(format string, addCallDepth int, parameters ...interface{}) {
	if !devMode {
		return
	}
	//@TODO benchmark conversion efficiency
	Log.Output(Calldepth+addCallDepth, OutputLevelDebug, logBlue(fmt.Sprintf(format+"\n", parameters...)))
}

func DebuglnWithDepth(addCallDepth int, parameters ...interface{}) {
	if !devMode {
		return
	}
	Log.Output(Calldepth+addCallDepth, OutputLevelDebug, logBlue(fmt.Sprintln(parameters...)))
}

func PrintfWithDepth(format string, addCallDepth int, parameters ...interface{}) {
	if !devMode {
		return
	}
	Log.Output(Calldepth+addCallDepth, OutputLevelDebug, logBlue(fmt.Sprintf(format, parameters...)))
}

func DebugflnWithDepth(format string, addCallDepth int, parameters ...interface{}) {
	if !devMode {
		return
	}
	//@TODO benchmark conversion efficiency
	Log.Output(Calldepth+addCallDepth, OutputLevelDebug, logBlue(fmt.Sprintf(format+"\n", parameters...)))
}

func DebugfWithDepth(format string, addCallDepth int, parameters ...interface{}) {
	if !devMode {
		return
	}
	Log.Output(Calldepth+addCallDepth, OutputLevelDebug, logBlue(fmt.Sprintf(format, parameters...)))
}

func WarnlnWithDepth(addCallDepth int, parameters ...interface{}) {
	Log.Output(Calldepth+addCallDepth, OutputLevelWarn, logYellow(fmt.Sprintln(parameters...)))
}

func WarnflnWithDepth(format string, addCallDepth int, parameters ...interface{}) {
	Log.Output(Calldepth+addCallDepth, OutputLevelWarn, logYellow(fmt.Sprintf(format+"\n", parameters...)))
}

func WarnfWithDepth(format string, addCallDepth int, parameters ...interface{}) {
	Log.Output(Calldepth+addCallDepth, OutputLevelWarn, logYellow(fmt.Sprintf(format, parameters...)))
}
