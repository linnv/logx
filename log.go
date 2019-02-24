package logx

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

//Logx  contains under field for log entity
type Logx struct {
	underFile     *os.File
	toFile        bool
	DevMode       bool //if true, all debug level info will be ignored
	maxBuffer     int  //bytes,maximun size of buffer for one sync
	disableBuffer bool
	currentIndex  int
	buf           []byte
}

func (l *Logx) resetbuf() {
	l.currentIndex = 0
}

func (l *Logx) availableCount() int {
	return l.maxBuffer - l.currentIndex
}

func (l *Logx) DisableBuffer(disable bool) {
	if disable {
		l.disableBuffer = true
		return
	}
	l.disableBuffer = false
}

func (l *Logx) Sync() {
	l.underFile.Write(l.buf[:l.currentIndex])
	l.underFile.Sync()
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

	//@TODO optimize
	if l.toFile {
		if l.disableBuffer {
			l.underFile.Write(bs)
		}

		bytesLen := len(bs)
		if bytesLen > l.availableCount() {
			l.Sync()
			l.resetbuf()
			if bytesLen >= l.maxBuffer {
				bytesLen = l.maxBuffer
			}
		}

		for i, bi := l.currentIndex, 0; bi < bytesLen; i, bi = i+1, bi+1 {
			l.buf[i] = bs[bi]
		}
		l.currentIndex += bytesLen
	}
	if level == outputLevelDebug {
		os.Stdout.Write(bs)
		return
	}
	//other level of log output to stderr
	os.Stderr.Write(bs)
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
	if l.underFile != nil {
		l.Sync()
		l.underFile.Close()
	}
}

func (l *Logx) LogConfigure() {
	println("to file:", l.toFile)
	println("dev mode:", l.DevMode)
	if l.underFile != nil {
		println("under file:", l.underFile.Name())
		println("DisableBuffer:", l.disableBuffer)
	} else {
		println("file no provied")
	}
	println("max buffer:", l.maxBuffer)
	println("current index:", l.currentIndex)
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

func newLogxFile() (newLog *Logx) {
	flags := GetFlags()
	l := &Logx{
		DevMode: flags.DevMode,
	}
	filepath := flags.FilePath
	if jsonConfig != nil {
		filepath = jsonConfig.FilePath
		l.DevMode = jsonConfig.DevMode
		l.disableBuffer = jsonConfig.DisableBuffer
		l.maxBuffer = jsonConfig.MaxbufferInt
	}
	if len(filepath) < 1 || !os.IsPathSeparator(filepath[0]) {
		return l
	}

	if err := checkDirAvailable(filepath); err != nil {
		panic(err.Error())
	}
	if fd, err := checkFileAvailable(filepath); err != nil {
		panic(err.Error())
	} else {
		newLogx(fd, l)
	}
	return l
}

func newLogx(fd *os.File, l *Logx) {
	if fd == nil || l == nil {
		return
	}

	l.underFile = fd
	l.toFile = true
	newSliceByte := func(n int) []byte {
		defer func() {
			if recover() != nil {
				panic(ErrTooLarge)
			}
		}()
		return make([]byte, n)
	}
	l.buf = newSliceByte(l.maxBuffer)
	return
}

func NewLogx() *Logx {
	return newLogxFile()
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
