package models

import (
	"fmt"
	"os"
)

type Logx struct {
	underFile *os.File `json:"underFile"`
	// outputLevel     int    `json:"outputLevel"`
	outputDirection byte
}

//@TODO configuration
//Debugln
func (l *Logx) Debugln(format string, paramters ...interface{}) {
	const prefix = "[debug]"
	// println(prefix, fmt.Sprintf(format+"\n", paramters...))
	bytes := []byte(prefix + fmt.Sprintf(format+"\n", paramters...))
	os.Stdout.Write(bytes)
	os.Stderr.Write(bytes)
}

//@TODO configuration
//Debug
func (l *Logx) Debug(format string, paramters ...interface{}) {
	const prefix = "[debug]"
	//@TODO benchmark convertion efficency
	bytes := []byte(prefix + fmt.Sprintf(format, paramters...))
	os.Stdout.Write(bytes)
	os.Stderr.Write(bytes)
}

//@TODO configuration
//Warning To File
func (l *Logx) Warn(format string, paramters ...interface{}) {
	const prefix = "[warn]"
	bytes := []byte(prefix + fmt.Sprintf(format+"\n", paramters...))
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
