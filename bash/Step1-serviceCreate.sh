#!/bin/bash
user="$1"
domain="$2"
#port="$3"
aSservice="$3"


if [ ! -f "/lib/systemd/system/$aSservice.service" ]; then

echo "[Unit]
Description=$aSservice $domain go port

[Service]
Type=simple
Restart=always
RestartSec=5s
EnvironmentFile=/home/$user/web/$domain/public_html/.env
ExecStart=/home/$user/web/$domain/public_html/main
WorkingDirectory=/home/$user/web/$domain/public_html/

[Install]
WantedBy=multi-user.target" > /lib/systemd/system/$aSservice.service
fi

systemctl enable $aSservice.service

systemctl start $aSservice

sudo systemctl daemon-reload

sudo systemctl restart $aSservice

sudo systemctl restart nginx

exit 0
