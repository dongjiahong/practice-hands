// Package logwithcontext log包
package logwithcontext

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const loggerKey = iota

var Logger *zap.Logger

func init() {
	level := zap.DebugLevel
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // json格式化日志
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),  // 打印到控制台和文件
		level,
	)

	// 开启文件和行号
	development := zap.Development()
	Logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel), // error 级别日志，打印堆栈
		development,
	)
}

func NewContext(ctx *gin.Context, fields ...zapcore.Field) {
	ctx.Set(strconv.Itoa(loggerKey), WithContext(ctx).With(fields...))
}

func WithContext(ctx *gin.Context) *zap.Logger {
	if ctx == nil {
		return Logger
	}
	l, _ := ctx.Get(strconv.Itoa(loggerKey))
	ctxLogger, ok := l.(*zap.Logger)
	if ok {
		return ctxLogger
	}
	return Logger
}
