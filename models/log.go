package models

import (
	"fmt"
	"os"
)

type Logx struct {
	underFile       *os.File
	maxBuffer       int //@TODO bytes,maximun size of buffer to output at least
	buf             []byte
	outputDirection byte
}

//@TODO use configuration
//Debugln
func (l *Logx) Debugln(format string, paramters ...interface{}) {
	// println(prefix, fmt.Sprintf(format+"\n", paramters...))
	bytes := []byte(prefixDebug + fmt.Sprintf(format+"\n", paramters...))
	os.Stdout.Write(bytes)
	os.Stderr.Write(bytes)
}

//@TODO use configuration
//Debug
func (l *Logx) Debug(format string, paramters ...interface{}) {
	//@TODO benchmark convertion efficency
	bytes := []byte(prefixDebug + fmt.Sprintf(format, paramters...))
	os.Stdout.Write(bytes)
	os.Stderr.Write(bytes)
}

//@TODO use configuration
//Warning To File
func (l *Logx) Warn(format string, paramters ...interface{}) {
	bytes := []byte(prefixWarn + fmt.Sprintf(format+"\n", paramters...))
	l.underFile.Write(bytes)
	os.Stdout.Write(bytes)
	os.Stderr.Write(bytes)
}

func (l *Logx) GracefullyExit() {
	l.underFile.Close()
}

//default logger
var Log = NewLogx()

func NewLogxFile() *Logx {
	filepath := "/Users/Jialin/myGit/OpenDemo/golang/main/logx/models/t.log"
	fd, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		//@TODO
		panic(err.Error())
	}

	return newLogx(fd)
}

func newLogx(fd *os.File) *Logx {
	return &Logx{
		underFile:       fd,
		outputDirection: outputLevelDebug,
	}
}

func NewLogx() *Logx {
	return newLogx(nil)
}
