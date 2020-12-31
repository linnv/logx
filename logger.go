package logx

type Logger interface {
	Flush() error
	Output(Calldepth int, level byte, content string)
	GetCallDepth() int
	Debugf(format string, parameters ...interface{})
	Debugfln(format string, parameters ...interface{})
	Warnf(format string, parameters ...interface{})
	Warnfln(format string, parameters ...interface{})
	Errorf(format string, parameters ...interface{})
	Errorfln(format string, parameters ...interface{})
}
