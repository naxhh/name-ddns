#!/usr/bin/env bash

cat > /env << EOF
export NAME_DDNS_USER=$NAME_DDNS_USER
export NAME_DDNS_TOKEN=$NAME_DDNS_TOKEN
export NAME_DDNS_DOMAIN=$NAME_DDNS_DOMAIN
export NAME_DDNS_HOST=$NAME_DDNS_HOST
EOF

cat > /crontab << EOF
$NAME_DDNS_UPDATE_CRON . /env; /bin/name-ddns $@ >> /cron.log 2>&1
# An empty line is required at the end of this file for a valid crontab file.
EOF

crontab /crontab
rm -f /crontab

/bin/name-ddns --initial-run >> /cron.log 2>&1
cron

tail -f /cron.log
