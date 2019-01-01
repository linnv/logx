package main

import (
	"github.com/linnv/logx"
)

func main() {
	logx.Debugln(111)
	logx.Warnln(111)
	logx.Errorln(111)

	logx.EnableDevMode(false)
	logx.Debugln("no log output on level debug ")
	logx.Warnln("no log output")
	logx.Errorln("no log output")
}
