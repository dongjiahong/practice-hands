package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"

	"webt/pkg/e"
)

type Gin struct {
	C *gin.Context
}

// 定义一个全局的翻译器
var trans ut.Translator

func init() {
	trans, _ = ut.New(zh.New()).GetTranslator("zh")
	zhTranslations.RegisterDefaultTranslations(binding.Validator.Engine().(*validator.Validate), trans)
}

func validateError(err error) (ret error) {
	var retStr string
	if validationErrors, ok := err.(validator.ValidationErrors); !ok {
		return err
	} else {
		for _, e := range validationErrors {
			retStr += e.Translate(trans) + ";"
		}
	}
	return fmt.Errorf("%s", retStr)
}

func (g *Gin) Response(httpCode, errCode int, data interface{}, err error) {
	var errStr string
	if err == nil {
		errStr = ""
	} else {
		errStr = validateError(err).Error()
	}

	g.C.JSON(http.StatusOK, gin.H{
		"code":       httpCode,
		"msg":        e.GetMsg(errCode),
		"data":       data,
		"err_detail": errStr,
	})
	return
}
