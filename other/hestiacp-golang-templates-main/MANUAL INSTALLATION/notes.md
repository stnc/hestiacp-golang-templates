#GOOS=windows GOARCH=amd64 go build main.go

GOOS=linux GOARCH=386 go build main.go //32 bit
GOOS=linux GOARCH=amd64 go build main.go  //64bit

GOOS=linux GOARCH=arm GOARM=5 go build main.go

chmod +rwx main

nginx.hsts.conf
path file /home/admin/conf/web/panel.example.com
set $go_web_port "9091";


nginx hata bakma
tail -f /var/log/nginx/error.log