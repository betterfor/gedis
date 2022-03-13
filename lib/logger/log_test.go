package logger

import "testing"

func TestExampleLogger(t *testing.T) {
	Debug("hello world")
	Info("hello world")
	Warn("hello world")
	Error("hello world")
	//Fatal("hello world")
}
