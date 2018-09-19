#!/bin/bash

chmod +x /var/run/manage
sleep 5
/var/run/manage  \
  -mysql.host=db \
  -mysql.port=3306 \
  -mysql.user=${MYSQL_USER} \
  -mysql.password=${MYSQL_PASSWORD} \
  -mysql.db=${MYSQL_DATABASE} \
  -mode=dev
