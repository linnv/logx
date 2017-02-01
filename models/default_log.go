package models

//default logger
var Log = NewLogx()

func Errorln(format string, paramters ...interface{}) {
	Log.Errorxln(format, paramters...)
}

func Error(format string, paramters ...interface{}) {
	Log.Errorx(format, paramters...)
}

func Debugln(format string, paramters ...interface{}) {
	Log.Debugxln(format, paramters...)
}

func Debug(format string, paramters ...interface{}) {
	Log.Debugx(format, paramters...)
}
