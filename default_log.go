package logx

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

//default logger
var Log *Logx

func Fatalln(parameters ...interface{}) {
	Log.Output(Calldepth, OutputLevelFatal, logRed(fmt.Sprintln(parameters...)))
	Log.GracefullyExit()
	os.Exit(1)
}

func Fatalf(format string, parameters ...interface{}) {
	Log.Output(Calldepth, OutputLevelFatal, logRed(fmt.Sprintf(format, parameters...)))
	Log.GracefullyExit()
	os.Exit(1)
}

func Fatalfln(format string, parameters ...interface{}) {
	Log.Output(Calldepth, OutputLevelFatal, logRed(fmt.Sprintf(format+"\n", parameters...)))
	Log.GracefullyExit()
	os.Exit(1)
}

func Errorln(parameters ...interface{}) {
	Log.Output(Calldepth, OutputLevelError, logRed(fmt.Sprintln(parameters...)))
}

func Errorf(format string, parameters ...interface{}) {
	Log.Output(Calldepth, OutputLevelError, logRed(fmt.Sprintf(format, parameters...)))
}

func Errorfln(format string, parameters ...interface{}) {
	Log.Output(Calldepth, OutputLevelError, logRed(fmt.Sprintf(format+"\n", parameters...)))
}

func Println(parameters ...interface{}) {
	if !devMode {
		return
	}
	Log.Output(Calldepth, OutputLevelDebug, logBlue(fmt.Sprintln(parameters...)))
}

func Printfln(format string, parameters ...interface{}) {
	if !devMode {
		return
	}
	//@TODO benchmark conversion efficiency
	Log.Output(Calldepth, OutputLevelDebug, logBlue(fmt.Sprintf(format+"\n", parameters...)))
}

func Debugln(parameters ...interface{}) {
	if !devMode {
		return
	}
	Log.Output(Calldepth, OutputLevelDebug, logBlue(fmt.Sprintln(parameters...)))
}

func Printf(format string, parameters ...interface{}) {
	if !devMode {
		return
	}
	Log.Output(Calldepth, OutputLevelDebug, logBlue(fmt.Sprintf(format, parameters...)))
}

func Debugfln(format string, parameters ...interface{}) {
	if !devMode {
		return
	}
	//@TODO benchmark conversion efficiency
	Log.Output(Calldepth, OutputLevelDebug, logBlue(fmt.Sprintf(format+"\n", parameters...)))
}

func Debugf(format string, parameters ...interface{}) {
	if !devMode {
		return
	}
	Log.Output(Calldepth, OutputLevelDebug, logBlue(fmt.Sprintf(format, parameters...)))
}

func Warnln(parameters ...interface{}) {
	Log.Output(Calldepth, OutputLevelWarn, logYellow(fmt.Sprintln(parameters...)))
}

func Warnfln(format string, parameters ...interface{}) {
	Log.Output(Calldepth, OutputLevelWarn, logYellow(fmt.Sprintf(format+"\n", parameters...)))
}

func Warnf(format string, parameters ...interface{}) {
	Log.Output(Calldepth, OutputLevelWarn, logYellow(fmt.Sprintf(format, parameters...)))
}

func Flush() error {
	return Log.Flush()
}

func EnableDevMode(enabled bool) {
	Log.EnableDevMode(enabled)
	devMode = Log.DevMode
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

func EnableDebug(w http.ResponseWriter, r *http.Request) {
	debug := r.FormValue("debug")
	debug = strings.ToLower(debug)
	if strings.Contains(debug, "on") || strings.Contains(debug, "true") {
		EnableDevMode(true)
	} else {
		EnableDevMode(false)
	}
	fmt.Fprintf(w, " log debug feature:%v", Log.DevMode)
}
