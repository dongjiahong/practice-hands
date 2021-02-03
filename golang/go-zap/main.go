package main

import (
	"zap-test/log"
)

func Log() {
	log.Logger.Info("this is info")
	log.Logger.Debug("this is Debug")
	log.Logger.Error("this is error")
}

func main() {
	Log()
}
