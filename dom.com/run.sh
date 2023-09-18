#!/bin/bash

workingfolder="/home/admin/web/dom.com/public_html/"

cd $workingfolder

echo "server: dom.com
port: 9898
service: dom.com" > /home/admin/web/dom.com/public_html/data.json

chmod +rwx main

systemctl enable ol.com.service

systemctl start ol.com

sudo systemctl daemon-reload

sudo systemctl restart ol.com

sudo systemctl restart nginx

exit 0