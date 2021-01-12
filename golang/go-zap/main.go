package main

import (
	"zap-test/log"
	"zap-test/log2"
)

func main() {
	log.Logger.Info("info")
	log2.Logger.Info("this is info")
	log2.Logger.Debug("this is Debug")
	log2.Logger.Error("this is error")
}
