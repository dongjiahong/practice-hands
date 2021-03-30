package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	v1 "webt/api/v1"
	_ "webt/docs"
	"webt/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.)

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // api注释
	apiv1 := r.Group("api/v1")
	apiv1.GET("/hello", v1.Hello) // echo hello

	return r
}
