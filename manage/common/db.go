package common

import (
	"github.com/jinzhu/gorm"
	"github.com/xiaofengshuyu/vpn-manager/manage/config"
	"github.com/xiaofengshuyu/vpn-manager/manage/models"
)

var (
	// DB is global db client
	DB *gorm.DB
)

func initDB() {
	var err error
	uri := config.AppConfig.MYSQL.BuildClientURI()
	Logger.Info(uri)
	DB, err = gorm.Open("mysql", uri)
	if err != nil {
		Logger.Error(err)
	}

	// auto migration
	if config.AppConfig.Mode == config.DEV {
		DB.LogMode(true)
		DB.Set("gorm:table_options", "engine=InnoDB").
			Set("gorm:table_options", "charset=utf8").
			AutoMigrate(
				&models.CommonUser{},
			)
	}
}
