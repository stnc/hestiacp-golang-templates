package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
https://github.com/hestiacp/hestiacp-generator/blob/master/js.js
https://github.com/charmbracelet/bubbletea/tree/master/examples
https://www.udemy.com/course/introduction-to-testing-in-go-golang/learn/lecture/33526058#overview
*/

/*
run sudo sh golang_Service_Create_And_Nginx_Port.sh admin test.example.com 8080 goweb
ssh parameter description

user= hestia root user name

domain=website domain name

port=the port that golang will use

asservice=name of golang service

*/

func main() {
	var userName, domainName, serviceName string
	var portNumber int64

	fmt.Print("Please enter root user name: ")
	fmt.Scanf("%v", &userName)

	fmt.Print("Please enter domain name: ")
	fmt.Scanf("%v", &domainName)

	fmt.Print("Please enter port number that golang will use for 'example 8080': ")
	fmt.Scanf("%v", &portNumber)

	fmt.Print("Please enter that golang service name for 'example goweb': ")
	fmt.Scanf("%v\n", &serviceName)

	fmt.Println("Would you like to continue with the installation? [Y/N]")

	fileCreatedNginxHst(portNumber)
	serviceCreatedUbuntu(portNumber, serviceName, domainName, userName)
	serviceRun(portNumber, serviceName, domainName, userName)
	/*
	   	//step 1
	   	//copy golang.stpl ve golang.tpl path = /usr/local/hestia/data/templates/web/nginx/php-fpm

	      #!/bin/bash

	      cp golang.tpl /usr/local/hestia/data/templates/web/nginx/php-fpm

	      cp golang.stpl /usr/local/hestia/data/templates/web/nginx/php-fpm

	      sudo systemctl daemon-reload

	      exit 0
	*/
}
func fileCreatedNginxHst(portNumber int64) {
	fileCreatedNginxHstV := `set $go_web_port  ` + strconv.FormatUint(uint64(portNumber), 10) + `";`
	Write(fileCreatedNginxHstV, "nginx.hsts.conf")
}

func serviceCreatedUbuntu(portNumber int64, serviceName, domainName, userName string) {
	serviceCreatedUbuntuV := `[Unit]
Description=` + serviceName + ` ` + strconv.FormatUint(uint64(portNumber), 10) + `
[Service]
Type=simple
Restart=always
RestartSec=5s
EnvironmentFile=/home/` + userName + `/web/` + domainName + `/public_html/.env
ExecStart=/home/` + userName + `/web/` + domainName + `/public_html/main
WorkingDirectory=/home/` + userName + `/web/` + domainName + `/public_html/

[Install]
WantedBy=multi-user.target" 
`
	Write(serviceCreatedUbuntuV, serviceName+".service")
}

func serviceRun(portNumber int64, serviceName, domainName, userName string) {
	serviceCreatedUbuntuV := `workingfolder="/home/` + userName + `/web/` + domainName + `/public_html/"

cd $workingfolder

echo "server: ` + domainName + `
port: ` + strconv.FormatUint(uint64(portNumber), 10) + `
service: ` + domainName + `" > /home/` + userName + `/web/` + domainName + `/public_html/data.json

chmod +rwx main

systemctl enable ` + serviceName + `.service

systemctl start ` + serviceName + `

sudo systemctl daemon-reload

sudo systemctl restart ` + serviceName + `

sudo systemctl restart nginx

exit 0`
	Write(serviceCreatedUbuntuV, "run.sh")
}

func Write(value, filename string) {

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(value)

	l, err := f.WriteString(value)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	currentTime := time.Now()
	fmt.Println(l, currentTime, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
