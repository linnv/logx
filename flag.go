// Package  provides ...
package logx

import "flag"

var initFlag bool

//InitFlags provides configure parameters of logx by args, you should call this functions after all of other flags have been defined
func InitFlags() {
	flag.String("logxfile", "", "absolut path of file,if empty no log will go into file")
	initFlag = true
}

//@TODO maybe not only one configure parameter through flag
func GetFlags() string {
	if !initFlag {
		return ""
	}
	if flag.Parsed() {
		var file string
		visitor := func(f *flag.Flag) {
			if f.Name == "logxfile" {
				file = f.Value.String()
			}
		}
		flag.Visit(visitor)
		return file

	}
	flag.Parse()
	return GetFlags()
}
