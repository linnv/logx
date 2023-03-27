package logx

import "testing"

func TestErrorlnWithDepth(t *testing.T) {
	DebuglnWithDepth(0, "abcd")
	Flush()
}
