package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2" // 分割日志
)

var Logger *zap.SugaredLogger

func init() {
	var coreArr []zapcore.Core

	// 获取编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "function",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 按级别显示不同颜色，不需要的话用zapcore.CapitalLevelEncoder
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // 指定时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,   // 一般zapcore.SecondsDurationEncoder, 执行消耗的时间转化为浮点型的秒
		EncodeCaller:   zapcore.ShortCallerEncoder,       // 一般zapcore.ShortCallerEncoder,以包/文件：行号 格式化调用堆栈
	}
	//encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 指定时间格式
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 按级别显示不同颜色，不需要的话用zapcore.CapitalLevelEncoder
	////encoderConfig.EncodeCaller = zapcore.FullCallerEncoder       // 显示完整路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig) // NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式

	// 日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // info和debug级别，debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	// info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./info.log", // 日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    200,          // 文件限制，单位MB
		MaxAge:     7,            // 日志保留的天数
		MaxBackups: 10,           // 最大保留日志文件数量
		Compress:   false,        // 是否压缩日志
	})
	infoFileCore := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)),
		lowPriority) // 第三个及之后的参数为写入文件的日志级别，ErrorLevel模式只记录error级别的日志

	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./error.log", // 日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    200,           // 文件限制，单位MB
		MaxAge:     7,             // 日志保留的天数
		MaxBackups: 10,            // 最大保留日志文件数量
		Compress:   false,         // 是否压缩日志
	})
	errorFileCore := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)),
		highPriority) // 第三个及之后的参数为写入文件的日志级别，ErrorLevel模式只记录error级别的日志

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	Logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()).Sugar() // zap.AddCaller()为显示文件名和行号，可省略
	//Logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller(), zap.AddStacktrace(highPriority)).Sugar() // zap.AddCaller()为显示文件名和行号，可省略
}
