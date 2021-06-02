package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"zap-test/log"
	zlog "zap-test/logWithContext"
)

func Log() {
	log.Logger.Info("this is info")
	log.Logger.Debug("this is Debug")
	log.Logger.Error("this is error")
}

type TestContextLog struct{}

func (t TestContextLog) Test(ctx *gin.Context) {
	name := ctx.Query("name")

	zlog.WithContext(ctx).Debug("测试日志", zap.String("name", name))
}

func main() {
	Log()
}
