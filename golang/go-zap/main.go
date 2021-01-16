package main

import (
	"zap-test/log"
)

func main() {
	log.Logger.Info("this is info")
	log.Logger.Debug("this is Debug")
	log.Logger.Error("this is error")
}
