#!/bin/bash

go run main.go \
  -mysql.host=127.0.0.1 \
  -mysql.port=3306 \
  -mysql.user=demo \
  -mysql.password=demo \
  -mysql.db=demo \
  -mode=dev