package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 测试
// @version 0.0.1
// @description 测试
// @BasePath /api/v1/
func main() {
	r := gin.New()

	// 创建路由组
	v1 := r.Group("/api/v1")
	v1.GET("/record/:userId", record)

	// 文档界面方位URL
	// http://127.0.0.1:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

// @获取制定ID记录
// @Description get record by ID
// @Accept json
// @Produce json
// @Param some_id path int true "userId"
// @Success 200 {string} string "ok"
// @Router /record/{some_id} [get]
func record(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
