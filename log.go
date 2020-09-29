package logx

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/linnv/bufferlog"
)

//Logx a simple log
type Logx struct {
	writer bufferlog.BufferLogger //todo multi-writer

	DevMode bool //if true, all debug level info will be ignored, default is true
}

func (l *Logx) SetWriter(w bufferlog.BufferLogger) {
	l.writer = w
}

func (l *Logx) Write(bs []byte) (err error) {
	_, err = l.writer.Write(bs)
	return
}

func (l *Logx) output(calldepth int, level byte, content string) {
	if level == outputLevelDebug && !l.DevMode {
		return
	}

	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if os.IsPathSeparator(file[i]) {
			short = file[i+1:]
			break
		}
	}
	file = short

	//30 for datatime, 5 for separetor
	excludeLen := len(content) + len(file) + len(prefix[level]) + 35
	bs := make([]byte, 0, excludeLen)
	bs = append(bs, prefix[level]...)
	bs = append(bs, ' ')
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

	bs = append(bs, ' ')
	bs = append(bs, file...)
	bs = append(bs, ' ')

	bs = append(bs, strconv.Itoa(line)...)
	bs = append(bs, ':')
	bs = append(bs, content...)

	if n, err := l.writer.Write(bs); err != nil {
		errStr := "wrote " + strconv.Itoa(n) + " bytes want " + strconv.Itoa(len(bs)) + " bytes, err:" + err.Error()
		print(errStr)
	}
}

func (l *Logx) EnableDevMode(enabled bool) {
	if enabled {
		l.DevMode = true
		return
	}
	l.DevMode = false
}

func (l *Logx) Printf(format string, parameters ...interface{}) {
	if !l.DevMode {
		return
	}

	l.output(calldepth, outputLevelDebug, fmt.Sprintf(format, parameters...))
}

func (l *Logx) Printfln(format string, parameters ...interface{}) {
	if !l.DevMode {
		return
	}

	l.output(calldepth, outputLevelDebug, fmt.Sprintf(format+"\n", parameters...))
}

func (l *Logx) Println(parameters ...interface{}) {
	if !l.DevMode {
		return
	}

	l.output(calldepth, outputLevelDebug, fmt.Sprintln(parameters...))
}

func (l *Logx) Debugf(format string, parameters ...interface{}) {
	if !l.DevMode {
		return
	}

	//@TODO benchmark conversion efficency
	l.output(calldepth, outputLevelDebug, fmt.Sprintf(format, parameters...))
}

func (l *Logx) Debugfln(format string, parameters ...interface{}) {
	if !l.DevMode {
		return
	}

	l.output(calldepth, outputLevelDebug, fmt.Sprintf(format+"\n", parameters...))
}

func (l *Logx) Debugln(parameters ...interface{}) {
	if !l.DevMode {
		return
	}

	l.output(calldepth, outputLevelDebug, fmt.Sprintln(parameters...))
}

func (l *Logx) Warnf(format string, parameters ...interface{}) {
	l.output(calldepth, outputLevelWarn, fmt.Sprintf(format, parameters...))
}

func (l *Logx) Warnfln(format string, parameters ...interface{}) {
	l.output(calldepth, outputLevelWarn, fmt.Sprintf(format+"\n", parameters...))
}

func (l *Logx) Warnln(parameters ...interface{}) {
	l.output(calldepth, outputLevelWarn, fmt.Sprintln(parameters...))
}

func (l *Logx) Fatalf(format string, parameters ...interface{}) {
	l.output(calldepth, outputLevelFatal, fmt.Sprintf(format, parameters...))
	l.GracefullyExit()
	os.Exit(1)
}

func (l *Logx) Flush() error {
	return l.writer.Flush()
}

func (l *Logx) Fatalln(parameters ...interface{}) {
	l.output(calldepth, outputLevelFatal, fmt.Sprintln(parameters...))
	l.GracefullyExit()
	os.Exit(1)
}

func (l *Logx) Errorf(format string, parameters ...interface{}) {
	l.output(calldepth, outputLevelError, fmt.Sprintf(format, parameters...))
}

func (l *Logx) Errorfln(format string, parameters ...interface{}) {
	l.output(calldepth, outputLevelError, fmt.Sprintf(format+"\n", parameters...))
}

func (l *Logx) Errorln(parameters ...interface{}) {
	l.output(calldepth, outputLevelError, fmt.Sprintln(parameters...))
}

//GracefullyExit implements flush log buffer to undferfile and close it
func (l *Logx) GracefullyExit() {
	if l.writer != nil {
		l.Flush()
		l.writer.Close()
	}
}

func NewLogx(w bufferlog.BufferLogger) *Logx {
	l := &Logx{
		writer:  w,
		DevMode: true,
	}
	return l
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
