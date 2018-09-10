package common

import (
	"github.com/sirupsen/logrus"
	"github.com/xiaofengshuyu/vpn-manager/manage/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	// Logger is a log for system run
	Logger *logrus.Logger
	// AccessLogger is a log for access
	AccessLogger *logrus.Logger
)

func initLog() {

	Logger = logrus.New()
	AccessLogger = logrus.New()

	if config.AppConfig.Mode == config.PROD {
		Logger.SetOutput(&lumberjack.Logger{
			Filename: "/tmp/log/run.log",
			MaxAge:   90,
			Compress: true,
		})
		Logger.SetFormatter(&logrus.JSONFormatter{})

		AccessLogger.SetOutput(&lumberjack.Logger{
			Filename: "/tmp/log/access.log",
			MaxAge:   90,
			Compress: true,
		})
		AccessLogger.SetFormatter(&logrus.JSONFormatter{})
	}
}
