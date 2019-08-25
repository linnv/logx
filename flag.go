package logx

import (
	"flag"
	"sync"

	"github.com/linnv/bufferlog"
)

var once sync.Once

var devMode bool

const (
	flagNameDev = "logxDev"
)

//Init provides configure parameters of logx by args, you should call this functions after all of other flags have been defined
func init() {
	devMode = *flag.Bool(flagNameDev, true, "if true all log of debug level will be outputted or will be ignored,the default value is true")
	once.Do(initDefaultLog)
}

func initDefaultLog() {
	Log = NewLogx(bufferlog.Buffer)
	Log.EnableDevMode(devMode)
}
