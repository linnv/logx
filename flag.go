package logx

import (
	"flag"
	"strings"
)

var initFlag bool

//Init provides configure parameters of logx by args, you should call this functions after all of other flags have been defined
func Init() {
	flag.String("logxfile", "", "absolut path of file,if empty no log will go into file")
	flag.Bool("defaultLogToFile", false, "flush to file in default mode eighter")
	initFlag = true
	initDefaultLog()
}

func GetFlags() (string, bool) {
	if !initFlag {
		return "", false
	}
	if flag.Parsed() {
		var file string
		var toFileInDebugMode bool
		fileVisitor := func(f *flag.Flag) {
			if f.Name == "logxfile" {
				file = f.Value.String()
			}
		}
		modeVisitor := func(f *flag.Flag) {
			if f.Name == "defaultLogToFile" {
				s := f.Value.String()
				if strings.ToUpper(s) == "TRUE" {
					toFileInDebugMode = true
				}
			}
		}
		flag.Visit(fileVisitor)
		flag.Visit(modeVisitor)
		return file, toFileInDebugMode

	}
	flag.Parse()
	return GetFlags()
}
