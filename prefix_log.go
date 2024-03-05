package logx

import "fmt"

type PrefixLog struct {
	Prefix string
	Log    Logger
}

func (pl *PrefixLog) Infof(format string, parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelInfo, fmt.Sprintf(format, parameters...))
}

func (pl *PrefixLog) Infofln(format string, parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelInfo, fmt.Sprintf(format+"\n", parameters...))
}

func (pl *PrefixLog) Infoln(parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelInfo, fmt.Sprintln(parameters...))
}
func (pl *PrefixLog) Warnf(format string, parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelWarn, fmt.Sprintf(format, parameters...))
}

func (pl *PrefixLog) Warnfln(format string, parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelWarn, fmt.Sprintf(format+"\n", parameters...))
}

func (pl *PrefixLog) Errorf(format string, parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelError, fmt.Sprintf(format, parameters...))
}

func (pl *PrefixLog) Errorfln(format string, parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelError, fmt.Sprintf(format+"\n", parameters...))
}

func (pl *PrefixLog) Debugf(format string, parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelDebug, fmt.Sprintf(format, parameters...))
}

func (pl *PrefixLog) Debugfln(format string, parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelDebug, fmt.Sprintf(format+"\n", parameters...))
}

func (pl *PrefixLog) Debugln(parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelDebug, fmt.Sprintln(parameters...))
}

func (pl *PrefixLog) Warnln(parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelWarn, fmt.Sprintln(parameters...))
}

func (pl *PrefixLog) Errorln(parameters ...interface{}) {
	if pl == nil {
		return
	}
	pl.Output(pl.GetCallDepth(), OutputLevelError, fmt.Sprintln(parameters...))
}

func (pl *PrefixLog) Flush() error {
	if pl == nil {
		return nil
	}
	return pl.Log.Flush()
}

func (pl *PrefixLog) GetCallDepth() int {
	if pl == nil {
		return Calldepth
	}
	return pl.Log.GetCallDepth() + 1
}

func (pl *PrefixLog) Output(calldepth int, level int32, content string) {
	if pl == nil {
		return
	}
	pl.Log.Output(calldepth, level, pl.Prefix+content)
}

func NewPrefixLogx(prefix string, rawLogger Logger) Logger {
	return &PrefixLog{
		Prefix: prefix,
		Log:    rawLogger,
	}
}
