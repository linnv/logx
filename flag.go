package logx

import (
	"flag"
	"sync"
)

var initFlag bool

var once sync.Once

//Init provides configure parameters of logx by args, you should call this functions after all of other flags have been defined
func init() {
	flag.String("logxfile", "", "absolut path of file,if empty no log will go into file")
	flag.Bool("defaultLogToFile", false, "flush to file in default mode eighter")
	once.Do(initDefaultLog)
}

func InitAndParsed() {
	initFlag = true
	initDefaultLog()
}

func initDefaultLog() {
	Log = NewLogx()
}

type ConfigByFlag struct {
	FilePath string
}

func GetFlags() ConfigByFlag {
	if !initFlag {
		return ConfigByFlag{}
	}

	if flag.Parsed() {
		var config ConfigByFlag
		fileVisitor := func(f *flag.Flag) {
			if f.Name == "logxfile" {
				config.FilePath = f.Value.String()
			}
		}
		flag.Visit(fileVisitor)
		return config

	}
	flag.Parse()
	return GetFlags()
}
