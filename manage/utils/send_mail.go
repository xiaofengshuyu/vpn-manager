package utils

import (
	"fmt"
	"net/smtp"
	"strings"

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

var emailTmpl = `From: VPN <%s>
To: Receiver <%s>
Subject: %s

%s
`

// SendSimpleEmail is send an simple email
func SendSimpleEmail(receiver []string, subject string, body string) error {
	data := fmt.Sprintf(emailTmpl, config.AppConfig.SMTP.User, strings.Join(receiver, ","), subject, body)
	return smtp.SendMail(
		fmt.Sprintf("%s:%d", config.AppConfig.SMTP.Host, config.AppConfig.SMTP.Port),
		auth,
		config.AppConfig.SMTP.User,
		receiver,
		[]byte(data),
	)
}
