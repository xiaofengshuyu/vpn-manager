#!/bin/bash

set -e

echo ${RSYNC_PASSWORD} > /etc/rsyncd.password
chmod 600 /etc/rsyncd.password

mkdir -p /root/password

cat <<EOF > /run/rsync.sh
date
rsync -avz ${RSYNC_USERNAME}@${RSYNC_HOST}::${RSYNC_MODULE} --password-file=/etc/rsyncd.password /root/password/
if [ -f "/root/password/passwd" ];then
  cp /root/password/passwd /etc/ipsec.d/passwd
fi

if [ -f "/root/password/chap-secrets" ];then
  cp /root/password/chap-secrets /etc/ppp/chap-secrets
fi
EOF

sed -i '/session    required     pam_loginuid.so/c\#session    required   pam_loginuid.so' /etc/pam.d/cron

# write crontab
cron

if [ ${RSYNC_ENABLE} == "true" ] 
then
  cat <<EOF >> /run/cron.rsync
* * * * * echo '1234' >> /tmp/haha.log
* * * * * bash /run/rsync.sh >> /var/log/rsync_recorder.log
EOF
crontab /run/cron.rsync
fi


# start server
/opt/src/run.sh