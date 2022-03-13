package main

import (
	"github.com/betterfor/gedis/lib/logger"
	"time"
)

func main() {
	logger.Debug("hello world")
	logger.Info("hello world")
	logger.Warn("hello world")
	logger.Error("hello world")
	logger.Fatal("hello world")

	time.Sleep(time.Second)
}
