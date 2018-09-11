package models

import (
	"fmt"

	zh_locales "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
	"github.com/xiaofengshuyu/vpn-manager/manage/config"
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
	var ok bool
	trans, ok = ut.New(zh, zh).GetTranslator("zh")
	if ok {
		zh_translations.RegisterDefaultTranslations(validate, trans)
	} else {
		common.Logger.Error("translator zh not found")
	}

	// auto migration
	if config.AppConfig.Mode == config.DEV {
		common.DB.LogMode(true)
		common.DB.Set("gorm:table_options", "engine=InnoDB").
			Set("gorm:table_options", "charset=utf8").
			AutoMigrate(
				&CommonUser{},
			)
	}

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
