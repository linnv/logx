package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/linnv/bufferlog"
	"github.com/linnv/logx"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	logx.Debugln(111)
	logx.Warnln(111)
	logx.Errorln(111)

	logx.EnableDevMode(false)
	logx.Debugln("no log output on level debug ")
	logx.Warnln("log output")
	logx.Errorln("log output")

	logx.EnableDevMode(true)
	logx.Debugln("log output on level debug with dev enabled")

	sigChan := make(chan os.Signal, 2)
	exit := make(chan struct{})
	fileBuffer := "./demoBuffer.log"
	under := &lumberjack.Logger{
		Filename:   fileBuffer,
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		LocalTime:  true,
		MaxAge:     28, // days
	}
	logWriter := bufferlog.NewBufferLog(3*1024, time.Second*2, exit, under)
	logger := logx.NewLogx(logWriter)
	logger.Debugln("juset demo")
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP)
	log.Print("use ctrl-c to exit: \n")
	<-sigChan
	close(exit)
	time.Sleep(time.Second * 3) //make sure logger has exited, or invoke Close() directly
}
