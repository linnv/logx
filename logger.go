package logx

type Logger interface {
	Flush() error
	Output(Calldepth int, level byte, content string)
	GetCallDepth() int
	Debugf(format string, parameters ...interface{})
	Debugfln(format string, parameters ...interface{})
	Debugln(parameters ...interface{})
	Infof(format string, parameters ...interface{})
	Infofln(format string, parameters ...interface{})
	Infoln(parameters ...interface{})
	Warnf(format string, parameters ...interface{})
	Warnfln(format string, parameters ...interface{})
	Warnln(parameters ...interface{})
	Errorf(format string, parameters ...interface{})
	Errorfln(format string, parameters ...interface{})
	Errorln(parameters ...interface{})
}
