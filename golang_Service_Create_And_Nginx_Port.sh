#!/bin/bash

user="$1"
domain="$2"
port="$3"
aSservice="$4"


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

if [ ! -f "/home/${user}/conf/web/${domain}/nginx.hsts.conf" ]; then

echo "set \$go_web_port \"${port}\"; " > /home/${user}/conf/web/${domain}/nginx.hsts.conf

fi

workingfolder="/home/$user/web/$domain/public_html/"

cd $workingfolder

echo "server: ${domain}
port: ${port}
service: ${aSservice}" > /home/$user/web/$domain/public_html/data.json

chmod +rwx main

# chmod 755 static

systemctl enable $aSservice.service

systemctl start $aSservice

sudo systemctl daemon-reload

sudo systemctl restart $aSservice

sudo systemctl restart nginx

exit 0
