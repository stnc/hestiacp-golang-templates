#!/bin/bash

workingfolder="/home/0l/web/oo.com/public_html/"

cd $workingfolder

echo "server: oo.com
port: 0
service: oo.com" > /home/0l/web/oo.com/public_html/data.json

chmod +rwx main

systemctl enable ww.service

systemctl start ww

sudo systemctl daemon-reload

sudo systemctl restart ww

sudo systemctl restart nginx

exit 0