package logx

import (
	"flag"
	"strings"
	"sync"
)

var initFlag bool

var once sync.Once

const (
	flagNameFile = "logxFile"
	flagNameDev  = "logxDev"
)

//Init provides configure parameters of logx by args, you should call this functions after all of other flags have been defined
func init() {
	once.Do(
		func() {
			// flag.String(flagNameFile, "", "absolut path of file,if empty no log will go into file")
			// flag.Bool(flagNameDev, true, "if true all log of debug level will be outputted or will be ignored,the default value is true")
			initDefaultLog()
		})

}

func InitAndParsed() {
	initFlag = true
	initDefaultLog()
}

func initDefaultLog() {
	Log = NewLogx()
}

//variables from cmd by flags, whis variabls will be assigned to log configurartion
//priority of configure from flags is less than flags
type ConfigByFlag struct {
	FilePath string
	DevMode  bool
}

//newConfigByFlag() will set default value as you want to fields
func newConfigByFlag() ConfigByFlag {
	return ConfigByFlag{
		DevMode: true,
	}
}

func GetFlags() ConfigByFlag {
	if !initFlag {
		return newConfigByFlag()
	}

	if flag.Parsed() {
		config := newConfigByFlag()
		fileVisitor := func(f *flag.Flag) {
			if f.Name == flagNameFile {
				config.FilePath = f.Value.String()
			}
		}
		modeVisitor := func(f *flag.Flag) {
			if f.Name == flagNameDev {
				if strings.ToUpper(f.Value.String()) == "FALSE" {
					config.DevMode = false
					return
				}
			}
		}
		flag.Visit(fileVisitor)
		flag.Visit(modeVisitor)
		return config

	}
	flag.Parse()
	return GetFlags()
}
