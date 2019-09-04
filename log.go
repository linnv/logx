package logx

import (
	"fmt"
	"os"
	"path"
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

	l.writer.Write(bs)
}

func (l *Logx) EnableDevMode(enabled bool) {
	if enabled {
		l.DevMode = true
		return
	}
	l.DevMode = false
}

func (l *Logx) Debugf(format string, paramters ...interface{}) {
	//@TODO benchmark convertion efficency
	l.output(calldepth, outputLevelDebug, fmt.Sprintf(format, paramters...))
}

func (l *Logx) Debugln(paramters ...interface{}) {
	l.output(calldepth, outputLevelDebug, fmt.Sprintln(paramters...))
}

func (l *Logx) Warnf(paramters ...interface{}) {
	l.output(calldepth, outputLevelWarn, fmt.Sprintln(paramters...))
}

func (l *Logx) Warnln(paramters ...interface{}) {
	l.output(calldepth, outputLevelWarn, fmt.Sprintln(paramters...))
}

func (l *Logx) Fatalf(format string, paramters ...interface{}) {
	l.output(calldepth, outputLevelFatal, fmt.Sprintf(format, paramters...))
	l.GracefullyExit()
	os.Exit(1)
}

func (l *Logx) Flush() error {
	return l.writer.Flush()
}

func (l *Logx) Fatalln(paramters ...interface{}) {
	l.output(calldepth, outputLevelFatal, fmt.Sprintln(paramters...))
	l.GracefullyExit()
	os.Exit(1)
}

func (l *Logx) Errorf(format string, paramters ...interface{}) {
	l.output(calldepth, outputLevelError, fmt.Sprintf(format, paramters...))
}

func (l *Logx) Errorln(paramters ...interface{}) {
	l.output(calldepth, outputLevelError, fmt.Sprintln(paramters...))
}

//GracefullyExit implements flush log buffer to undferfile and close it
func (l *Logx) GracefullyExit() {
	if l.writer != nil {
		l.Flush()
		l.writer.Close()
	}
}

func checkFileAvailable(filepath string) (*os.File, error) {
	fd, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		if os.IsPermission(err) {
			err = os.Chmod(filepath, 0666)
			if err != nil {
				return nil, err
			}
			return checkFileAvailable(filepath)
		}
		return nil, err
	}
	return fd, nil
}

func checkDirAvailable(filepath string) error {
	_, err := os.Stat(filepath)
	if err == nil {
	} else if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(filepath), 0766)
		if err != nil {
			return err
		}
	} else if os.IsPermission(err) {
		err = os.Chmod(path.Dir(filepath), 0755)
		if err != nil {
			return err
		}
	}
	return nil
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
