#!/bin/bash

# start sync server

echo "$RSYNC_USERNAME:$RSYNC_PASSWORD" > /etc/rsyncd.secrets
chmod 0400 /etc/rsyncd.secrets

mkdir -p /data
chmod a+rw /data

[ -f /etc/rsyncd.conf ] || cat <<EOF > /etc/rsyncd.conf
pid file = /var/run/rsyncd.pid
log file = /dev/stdout
timeout = 300
max connections = 100
[data]
    uid = root
    gid = root
    hosts deny = *
    read only = true
    hosts allow = ${HOST_ALLOW}
    path = /data
    comment = data directory
    auth users = $RSYNC_USERNAME
    secrets file = /etc/rsyncd.secrets
EOF

exec /usr/bin/rsync --no-detach --daemon --config /etc/rsyncd.conf "$@" &

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
  -rsync.path=/data \
  -mode=${RUN_MODE}
