#!/bin/bash

go run main.go \
  -mysql.host=127.0.0.1 \
  -mysql.port=3306 \
  -mysql.user=demo \
  -mysql.password=demo \
  -mysql.db=demo \
  -smtp.user=your@example.com \
  -smtp.password=pwd \
  -smtp.host=smtp.server.com \
  -smtp.port=25 \
  -mode=dev