package models

import (
	"fmt"

	zh_locales "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	trans    ut.Translator
	validate *validator.Validate
)

func init() {
	zh := zh_locales.New()
	validate = validator.New()
	trans, _ = ut.New(zh, zh).GetTranslator("zh")
	zh_translations.RegisterDefaultTranslations(validate, trans)
	/**
	 */
}

// TranslateAll is translate validator information
func TranslateAll(data interface{}) string {
	err := validate.Struct(data)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		result := ""
		for k, v := range errs.Translate(trans) {
			result += fmt.Sprintf("%s:%s;", k, v)
		}
		return result
	}
	return ""
}
