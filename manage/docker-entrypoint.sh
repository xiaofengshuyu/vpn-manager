#!/bin/bash

chmod +x /var/run/manage
sleep 5
echo ${ADMIN_USER}
echo ${ADMIN_PASSWORD}
echo ${MYSQL_HOST}
echo ${MYSQL_PORT}
echo ${APPSTORE_BUNDLE}

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
