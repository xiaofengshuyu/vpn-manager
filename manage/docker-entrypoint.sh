#!/bin/bash

# build
export GO111MODULE=on
cd /go/src/github.com/github.com/xiaofengshuyu/vpn-manager/manage/
go mod download
go build -o /var/run/manage main.go


sleep 5

/var/run/manage  \
  -auth.user=${ADMIN_USER} \
  -auth.password=${ADMIN_PASSWORD} \
  -mysql.host=${MYSQL_HOST} \
  -mysql.port=${MYSQL_PORT} \
  -mysql.user=${MYSQL_USER} \
  -mysql.password=${MYSQL_PASSWORD} \
  -mysql.db=${MYSQL_DATABASE} \
  -smtp.user=${SMTP_USER} \
  -smtp.password=${SMTP_PASSWORD} \
  -smtp.host=${SMTP_HOST} \
  -smtp.port=${SMTP_PORT} \
  -apple.bundle=${APPSTORE_BUNDLE} \
  -mode=${RUN_MODE}
