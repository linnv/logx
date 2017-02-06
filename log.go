package logx

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

type Logx struct {
	underFile *os.File
	toFile    bool
	// maxBuffer       int //@TODO bytes,maximun size of buffer to output at least
	// buf             []byte
}

func (l *Logx) output(calldepth int, level byte, content string) {
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		//@TODO differ by os type
		if os.IsPathSeparator(file[i]) {
			short = file[i+1:]
			break
		}
	}
	file = short

	bs := make([]byte, 0, 30)
	buf := &bs
	t := time.Now()
	year, month, day := t.Date()
	itoa(buf, year, 4)
	*buf = append(*buf, os.PathSeparator)
	itoa(buf, int(month), 2)
	*buf = append(*buf, os.PathSeparator)
	itoa(buf, day, 2)
	*buf = append(*buf, ' ')

	hour, min, sec := t.Clock()
	itoa(buf, hour, 2)
	*buf = append(*buf, ':')
	itoa(buf, min, 2)
	*buf = append(*buf, ':')
	itoa(buf, sec, 2)

	//@TODO optimize
	content = prefix[level] + string(*buf) + " " + file + " " + strconv.Itoa(line) + ": " + content
	bytes := []byte(content)
	if l.toFile {
		l.underFile.Write(bytes)
	}
	if level == outputLevelDebug {
		os.Stdout.Write(bytes)
		return
	}
	//other level of output to stderr
	os.Stderr.Write(bytes)
}

//@TODO use configuration
func (l *Logx) Debug(format string, paramters ...interface{}) {
	//@TODO benchmark convertion efficency
	l.output(calldepth, outputLevelDebug, fmt.Sprintf(format, paramters...))
}

func (l *Logx) Debugln(format string, paramters ...interface{}) {
	l.output(calldepth, outputLevelDebug, fmt.Sprintf(format+"\n", paramters...))
}

//default log is wrapped by one more function,so calldepth plus one
func (l *Logx) Debugx(format string, paramters ...interface{}) {
	//@TODO benchmark convertion efficency
	l.output(calldepth+1, outputLevelDebug, fmt.Sprintf(format, paramters...))
}

func (l *Logx) Debugxln(format string, paramters ...interface{}) {
	l.output(calldepth+1, outputLevelDebug, fmt.Sprintf(format+"\n", paramters...))
}

func (l *Logx) Warn(format string, paramters ...interface{}) {
	l.output(calldepth, outputLevelWarn, fmt.Sprintf(format, paramters...))
}

func (l *Logx) Warnln(format string, paramters ...interface{}) {
	l.output(calldepth, outputLevelWarn, fmt.Sprintf(format+"\n", paramters...))
}

//Warning To File
func (l *Logx) Fatal(format string, paramters ...interface{}) {
	l.output(calldepth, outputLevelFatal, fmt.Sprintf(format, paramters...))
	l.GracefullyExit()
	os.Exit(1)
}

func (l *Logx) Fatalln(format string, paramters ...interface{}) {
	l.output(calldepth, outputLevelFatal, fmt.Sprintf(format+"\n", paramters...))
	l.GracefullyExit()
	os.Exit(1)
}

func (l *Logx) Errorx(format string, paramters ...interface{}) {
	l.output(calldepth+1, outputLevelError, fmt.Sprintf(format, paramters...))
}

func (l *Logx) Errorxln(format string, paramters ...interface{}) {
	l.output(calldepth+1, outputLevelError, fmt.Sprintf(format+"\n", paramters...))
}
func (l *Logx) Error(format string, paramters ...interface{}) {
	l.output(calldepth, outputLevelError, fmt.Sprintf(format, paramters...))
}

func (l *Logx) Errorln(format string, paramters ...interface{}) {
	l.output(calldepth, outputLevelError, fmt.Sprintf(format+"\n", paramters...))
}

func (l *Logx) GracefullyExit() {
	if l.underFile != nil {
		l.underFile.Close()
	}
}

func (l *Logx) LogConfigure() {
	//@TODO show entity inof if Logx
}

func NewLogxFile() *Logx {
	filepath := GetFlags()
	if len(filepath) < 1 {
		return newLogx(nil)
	}
	if !os.IsPathSeparator(filepath[0]) {
		return newLogx(nil)
	}
	_, err := os.Stat(filepath)
	if err == nil {
	} else if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(filepath), 0766)
		if err != nil {
			panic(err.Error())
		}
	} else if os.IsPermission(err) {
		err = os.Chmod(path.Dir(filepath), 0755)
		if err != nil {
			panic(err.Error())
		}
	}

newFile:
	fd, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		if os.IsPermission(err) {
			err = os.Chmod(filepath, 0666)
			if err != nil {
				panic(err.Error())
			}
			goto newFile
		}
		panic("unknow error")
	}

	return newLogx(fd)
}

func newLogx(fd *os.File) (l *Logx) {
	l = new(Logx)
	if fd == nil {
		return
	}
	l.underFile = fd
	l.toFile = true
	return
}

func NewLogx() *Logx {
	return newLogx(nil)
}

// Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}
