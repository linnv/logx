package logx

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/linnv/bufferlog"
	"go.uber.org/zap"
)

// Logx a simple log
type Logx struct {
	writer bufferlog.BufferLogger //todo multi-writer

	level  atomic.Int32
	zaplog *zap.Logger
}

func (l *Logx) SetZapLogger(oneLogger *zap.Logger) {
	l.zaplog = oneLogger
}

func (l *Logx) SetWriter(w bufferlog.BufferLogger) {
	l.writer = w
}

func (l *Logx) Write(bs []byte) (err error) {
	_, err = l.writer.Write(bs)
	return
}

func (l *Logx) Output(Calldepth int, level int32, content string) {
	if l.zaplog != nil {
		content = strings.TrimSuffix(content, "\n")

		content = strings.Replace(content, "\n\x1b[0m", "\x1b[0m", 1)

		switch level {
		case OutputLevelDebug:
			l.zaplog.Debug(content)
		case OutputLevelInfo:
			l.zaplog.Info(content)
		case OutputLevelWarn:
			l.zaplog.Warn(content)
		case OutputLevelError:
			l.zaplog.Error(content)
		case OutputLevelFatal:
			l.zaplog.Fatal(content)
		}

		return
	}

	if level < l.level.Load() {
		return
	}

	pc, file, line, ok := runtime.Caller(Calldepth)
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

	callerFunc := runtime.FuncForPC(pc).Name()
	file = short

	excludeLen := len(callerFunc) + len(content) + len(file) + len(prefix[level]) + 22
	bs := make([]byte, 0, excludeLen)
	bs = append(bs, prefix[level]...)
	bs = append(bs, ' ')
	buf := &bs

	onetime := time.Now().Format("20060102 15:04:05.000")
	*buf = append(*buf, onetime...)

	bs = append(bs, ' ')
	bs = append(bs, file...)
	bs = append(bs, ' ')

	bs = append(bs, strconv.Itoa(line)...)
	bs = append(bs, ' ')
	bs = append(bs, callerFunc...)
	bs = append(bs, ':')
	bs = append(bs, content...)

	if n, err := l.writer.Write(bs); err != nil {
		errStr := "wrote " + strconv.Itoa(n) + " bytes want " + strconv.Itoa(len(bs)) + " bytes, err:" + err.Error()
		println(errStr)
	}
}

func (l *Logx) SetLevel(level int32) {
	l.level.Store(level)
}

func (l *Logx) EnableDevMode(enabled bool) {
	if enabled {
		l.level.Store(OutputLevelDebug)
		return
	}
}

func (l *Logx) Printf(format string, parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelDebug, fmt.Sprintf(format, parameters...))
}

func (l *Logx) Printfln(format string, parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelDebug, fmt.Sprintf(format+"\n", parameters...))
}

func (l *Logx) Println(parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelDebug, fmt.Sprintln(parameters...))
}

func (l *Logx) Infof(format string, parameters ...interface{}) {
	//@TODO benchmark conversion efficency
	l.Output(Calldepth, OutputLevelInfo, fmt.Sprintf(format, parameters...))
}

func (l *Logx) Infofln(format string, parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelInfo, fmt.Sprintf(format+"\n", parameters...))
}

func (l *Logx) Infoln(parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelInfo, fmt.Sprintln(parameters...))
}

func (l *Logx) Debugf(format string, parameters ...interface{}) {
	//@TODO benchmark conversion efficency
	l.Output(Calldepth, OutputLevelDebug, fmt.Sprintf(format, parameters...))
}

func (l *Logx) Debugfln(format string, parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelDebug, fmt.Sprintf(format+"\n", parameters...))
}

func (l *Logx) Debugln(parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelDebug, fmt.Sprintln(parameters...))
}

func (l *Logx) Warnf(format string, parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelWarn, fmt.Sprintf(format, parameters...))
}

func (l *Logx) Warnfln(format string, parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelWarn, fmt.Sprintf(format+"\n", parameters...))
}

func (l *Logx) Warnln(parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelWarn, fmt.Sprintln(parameters...))
}

func (l *Logx) Fatalf(format string, parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelFatal, fmt.Sprintf(format, parameters...))
	l.GracefullyExit()
	os.Exit(1)
}

func (l *Logx) Flush() error {
	return l.writer.Flush()
}

func (l *Logx) Fatalln(parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelFatal, fmt.Sprintln(parameters...))
	l.GracefullyExit()
	os.Exit(1)
}

func (l *Logx) Errorf(format string, parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelError, fmt.Sprintf(format, parameters...))
}

func (l *Logx) Errorfln(format string, parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelError, fmt.Sprintf(format+"\n", parameters...))
}

func (l *Logx) Errorln(parameters ...interface{}) {
	l.Output(Calldepth, OutputLevelError, fmt.Sprintln(parameters...))
}

// GracefullyExit implements flush log buffer to undferfile and close it
func (l *Logx) GracefullyExit() {
	if l.writer != nil {
		l.Flush()
		l.writer.Close()
	}
}

func NewLogx(w bufferlog.BufferLogger) *Logx {
	l := &Logx{
		writer: w,
	}
	l.level.Store(OutputLevelDebug)
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

func (l *Logx) GetCallDepth() int {
	return Calldepth
}
