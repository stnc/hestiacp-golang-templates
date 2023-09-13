#!/bin/bash

workingfolder="/home/admin/web/der.com/public_html/"

cd $workingfolder

echo "server: der.com
port: 9090
service: der.com" > /home/admin/web/der.com/public_html/data.json

chmod +rwx main

systemctl enable go.com.service

systemctl start go.com

sudo systemctl daemon-reload

sudo systemctl restart go.com

sudo systemctl restart nginx

exit 0