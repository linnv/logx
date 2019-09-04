package main

import (
	"github.com/linnv/bufferlog"
	"github.com/linnv/logx"
)

func main() {
	logger := logx.NewLogx(bufferlog.Buffer)
	logger.Debugln("ok, flush output")
	logger.Debugln("ok, delay output,no log  bellow maybe")
	logger.Flush()
	logger.Debugln("ok, delay output")
}
