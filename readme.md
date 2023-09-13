# HestiaCP Golang templates

## Careful! This is still in development and will probably break your server.

Golang templates for [HestiaCP](https://www.hestiacp.com/).

This project was originally based on the work done by [realjumy](https://github.com/realjumy/hestiacp-python-templates) 

https://gist.github.com/soheilhy/8b94347ff8336d971ad0

## Disclaimer

1. This code comes without warranty of any kind. Please refer to `README.md` for more details about this and the license to which this software is bounded. 
2. All this is still in experimental stage.
3. These templates will install application **in debug mode and without database connection**. Is therefore your responsibility to complete the configuration process and make the app safe.

## Requirements

- HestiaCP
- Golang 1.17

I presume it can be adapted to VestaCP after small modifications.

## Tested with

- [X] HestiaCP 1.6.7
- [X] Ubuntu 22.04 LTS
- [X] Golang 1.17


If you have tested it with a different version or different distro, feel free to contact me to provide feedback.

## Instructions for Ubuntu:

Recommend HESTIA Install



```bash
wget https://raw.githubusercontent.com/hestiacp/hestiacp/release/install/hst-install.sh


sudo bash hst-install.sh --apache no --phpfpm no --multiphp no --vsftpd no --proftpd yes --named yes --mysql yes --postgresql yes --exim yes --dovecot yes --sieve no --clamav yes --spamassassin yes --iptables yes --fail2ban yes --quota no --api yes --interactive yes --with-debs no  --port 2083 --hostname panel.kurbandefteri.com --email selmantunc@yandex.com --password yourpassword123456789 --lang en  --force
```
- NGINX templates (file golang.tpl, golang.stpl) goes into `/usr/local/hestia/data/templates/web/nginx/php-fpm` or run `sudo sh template-move.sh` 

4. Activate the template Web Template NGINX (select golang)

5. Activate the desired Backend Template PHP-FPM. It is recommended to set the backend template to `no-php`.

![Screen](https://raw.githubusercontent.com/stnc/hestiacp-golang-templates/main/hestia1.png)

6. Complete the setup process of the terminal. This includes setting up the database, adding the users, disabling the debug/setting environment to production, modifying the allowed hosts, and so on.

## SSL HTTPS SUPPORT

![Screen](https://raw.githubusercontent.com/stnc/hestiacp-golang-templates/main/hestiaSsl.png)

sslExample.go 

```go

package main

import (
	"log"
	"net/http"
)


func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an ssl https example server.\n"))
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":8080", "/home/admin/conf/web/us.example.com/ssl/test.example.com.crt", "/home/admin/conf/web/us.example.com/ssl/test.example.com.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


```
# BONUS (golang service and nginx port maker)

run `sudo sh golang_Service_Create_And_Nginx_Port.sh admin test.example.com 8080 goweb` 

## ssh parameter description

user= hestia root user name

domain=website domain name

port=the port that golang will use

asservice=name of golang service

```bash 

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

systemctl enable $aSservice.service

systemctl start $aSservice

sudo systemctl daemon-reload

sudo systemctl restart $aSservice

sudo systemctl restart nginx

exit 0


```
