#!/bin/bash

go run main.go \
  -mysql.host=127.0.0.1 \
  -mysql.port=3306 \
  -mysql.user=root \
  -mysql.password=123456 \
  -mysql.db=net_work_g \
  -mode=dev