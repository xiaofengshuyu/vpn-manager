package common

import (
	"github.com/jinzhu/gorm"
	"github.com/xiaofengshuyu/vpn-manager/manage/config"
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
}
