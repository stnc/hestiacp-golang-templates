golang.stpl ve golang.tpl dosyasini /usr/local/hestia/data/templates/web/nginx/php-fpm buraya gonder 


Activate the template Web Template NGINX (select golang)

Activate the desired Backend Template PHP-FPM. It is recommended to set the backend template to no-php.



arm64 build 
GOOS=linux GOARCH=arm GOARM=5 go build main.go  

GOOS=linux GOARCH=386 go build main.go //32 bit 

sonra hestiapanel web arayuzu ile gonder 

chmod +rwx main

nginx hata bakma 
tail -f /var/log/nginx/error.log


sudo nano /lib/systemd/system/gouama.service

# -------

[Unit]
Description=gokys

[Service]
Type=simple
Restart=always
RestartSec=5s
EnvironmentFile=/home/admin/web/panel.example.com/public_html/.env
ExecStart=/home/admin/web/panel.example.com/public_html/main
WorkingDirectory=/home/admin/web/panel.example.com/public_html/

[Install]
WantedBy=multi-user.target