package utils

import (
	"net/smtp"

	"github.com/xiaofengshuyu/vpn-manager/manage/config"
)

var (
	auth = smtp.PlainAuth(
		"",
		config.AppConfig.SMTP.User,
		config.AppConfig.SMTP.Password,
		config.AppConfig.SMTP.Host,
	)
)

func SendSimpleEmail(sender string, receiver []string, subject string, body []byte) error {

	return nil
}
