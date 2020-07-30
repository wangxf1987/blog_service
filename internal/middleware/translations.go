package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hans_HK"
	ut "github.com/go-playground/universal-translator"
	validatorv10 "github.com/go-playground/validator/v10"
	translations_en "github.com/go-playground/validator/v10/translations/en"
	translations_zh "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hans_HK.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validatorv10.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = translations_zh.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = translations_en.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = translations_zh.RegisterDefaultTranslations(v, trans)
			}
			c.Set("trans", trans)
		}
	}
}
