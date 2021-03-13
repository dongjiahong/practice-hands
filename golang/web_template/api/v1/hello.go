package v1

import (
	"fmt"

	"webt/models"
	"webt/pkg/app"
)

// @Summary Hello
// @Tag Example
// @Param name query string true "你的名字"
// @Router /api/v1/hello [GET]
func Hello(c *gin.Context) {
	appG := app.Gin{C: c}

	var hello models.HelloSwag
}
