package models

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type Logx struct {
	underFile *os.File
	toFile    bool
	// maxBuffer       int //@TODO bytes,maximun size of buffer to output at least
	// buf             []byte
}

func (l *Logx) output(level byte, content string) {
	const calldepth = 2
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		//@TODO differ by os type
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	content = file + "," + strconv.Itoa(line) + ": " + content
	var bytes []byte
	if level == outputLevelDebug {
		bytes = []byte(prefixDebug + content)
		os.Stdout.Write(bytes)
	}
	if level == outputLevelWarn {
		bytes = []byte(prefixWarn + content)
		os.Stderr.Write(bytes)
	}
	if level == outputLevelError {
		bytes = []byte(prefixError + content)
		os.Stderr.Write(bytes)
	}
	if level == outputLevelFatal {
		bytes = []byte(prefixFatal + content)
	}
	if l.toFile {
		l.underFile.Write(bytes)
	}
}

//@TODO use configuration
//Debug
func (l *Logx) Debug(format string, paramters ...interface{}) {
	//@TODO benchmark convertion efficency
	l.output(outputLevelDebug, fmt.Sprintf(format, paramters...))
}

//Debugln
func (l *Logx) Debugln(format string, paramters ...interface{}) {
	l.output(outputLevelDebug, fmt.Sprintf(format+"\n", paramters...))
}

//Warning To File
func (l *Logx) Warn(format string, paramters ...interface{}) {
	l.output(outputLevelWarn, fmt.Sprintf(format, paramters...))
}

func (l *Logx) Warnln(format string, paramters ...interface{}) {
	l.output(outputLevelWarn, fmt.Sprintf(format+"\n", paramters...))
}

//Warning To File
func (l *Logx) Fatal(format string, paramters ...interface{}) {
	l.output(outputLevelFatal, fmt.Sprintf(format, paramters...))
	os.Exit(1)
}

func (l *Logx) Fatalln(format string, paramters ...interface{}) {
	l.output(outputLevelFatal, fmt.Sprintf(format+"\n", paramters...))
	os.Exit(1)
}

func (l *Logx) Error(format string, paramters ...interface{}) {
	l.output(outputLevelError, fmt.Sprintf(format, paramters...))
}

func (l *Logx) Errorln(format string, paramters ...interface{}) {
	l.output(outputLevelError, fmt.Sprintf(format+"\n", paramters...))
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
	if fd == nil {
		return &Logx{
			underFile: nil,
			toFile:    false,
		}
	}
	return &Logx{
		underFile: fd,
		toFile:    true,
	}
}

func NewLogx() *Logx {
	return newLogx(nil)
}
