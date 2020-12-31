package logx

import (
	"testing"
)

func TestNewPrefixLogx(t *testing.T) {
	defer Log.Flush()

	aLog := NewPrefixLogx("a->", Log)
	aLog.Debugfln("log")
	bLog := NewPrefixLogx("b->", Log)
	bLog.Debugfln("log")
	oneFunc := func() {
		cLog := NewPrefixLogx("c->", aLog)
		cLog.Debugfln("log")
		innerLog(cLog)
	}

	oneFunc()
	innerLog(aLog)

	//output
	// [debug] 2020/12/31 11:54:16 prefix_log_test.go 11:a->log
	// [debug] 2020/12/31 11:54:16 prefix_log_test.go 13:b->log
	// [debug] 2020/12/31 11:54:16 prefix_log_test.go 16:a->c->log
	// [debug] 2020/12/31 11:54:16 prefix_log_test.go 33:a->c->inner->ok
	// [debug] 2020/12/31 11:54:16 prefix_log_test.go 33:a->inner->ok
}

func innerLog(oneLog Logger) {
	innerLog := NewPrefixLogx("inner->", oneLog)
	innerLog.Debugfln("ok")
}
