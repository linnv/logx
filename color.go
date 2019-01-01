package logx

import "fmt"

const (
	colorRed    = uint8(31)
	colorGreen  = uint8(32)
	colorYellow = uint8(33)
	colorBlue   = uint8(36)
)

func logBlue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorBlue, s)
}

func logRed(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorRed, s)
}

func logYellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorYellow, s)
}

func logGreen(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorGreen, s)
}
