package models

const (
	outputLevelDebug = byte(0x01)
	outputLevelWarn  = byte(0x02)

	outputLevelFatal = byte(0x03)
)

const (
	outputDirectionDebug = byte(0x01)
	outputDirectionWarn  = byte(0x02)
)

var (
	outputDirection = [...]string{"STDOUT", "STDERR", "FILE"}
	outputLevel     = [...]string{"DEBUG", "WARN", "FATAL"}
)
