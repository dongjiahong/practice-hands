package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	zlog "zap-test/logWithContext"
)

// -------------------------------------------------------------------------------
// 通过中间件，我们可以方便的为每个request添加日志上下文关键字段，我这儿只添加
// 了traceId和一些请求信息字段，我们还可以根据应用场景添加其它自定义字段。
// -------------------------------------------------------------------------------

func traceLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 每个请求生成的请求traceId具有全局唯一性
		u1, _ := uuid.NewV4()
		traceId := u1.String()
		zlog.NewContext(ctx, zap.String("traceId", traceId))

		// 为日志添加请求的地址以及请求参数等信息
		zlog.NewContext(ctx, zap.String("request.method", ctx.Request.Method))
		headers, _ := json.Marshal(ctx.Request.Header)
		zlog.NewContext(ctx, zap.String("request.headers", string(headers)))
		zlog.NewContext(ctx, zap.String("request.url", ctx.Request.URL.String()))

		// 将请求参数json序列化后添加进日志上下文
		if ctx.Request.Form == nil {
			ctx.Request.ParseMultipartForm(32 << 20)
		}
		form, _ := json.Marshal(ctx.Request.Form)
		zlog.NewContext(ctx, zap.String("request.params", string(form)))

		ctx.Next()
	}
}
