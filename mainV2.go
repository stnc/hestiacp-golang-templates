package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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

type Dir struct {
	mkdirFile string
}

func (server *Dir) setName(newName string) {
	server.mkdirFile = newName
}

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

	mkdir := Dir{}
	mkdir.setName(domainName) // changed
	fmt.Println(mkdir.mkdirFile)
	if err := mkdir.ensureDir(); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}
	mkdir.fileCreatedNginxHst(portNumber)
	mkdir.serviceCreatedUbuntu(portNumber, serviceName, domainName, userName)
	mkdir.serviceRun(portNumber, serviceName, domainName, userName)
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

func (server *Dir) ensureDir() error {
	var dirName string
	dirName = server.mkdirFile
	err := os.Mkdir(dirName, os.ModeDir)
	if err == nil {
		errs := os.Chmod(dirName, 0755)
		if errs != nil {
			fmt.Println("Error making file read-only:", errs)
			return errs
		}

		return nil
	}

	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dirName)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("path exists but is not a directory")
		}
		return nil
	}
	return err
}

func (server *Dir) fileCreatedNginxHst(portNumber int64) {

	//fileCreatedNginxHstV := `set $go_web_port  ` + strconv.FormatUint(uint64(portNumber), 10) + `;`
	//Write(fileCreatedNginxHstV, "nginx.hsts.conf")
}

func (server *Dir) serviceCreatedUbuntu(portNumber int64, serviceName, domainName, userName string) {
	//	serviceCreatedUbuntuV := `[Unit]
	//Description=` + serviceName + ` ` + strconv.FormatUint(uint64(portNumber), 10) + `
	//[Service]
	//Type=simple
	//Restart=always
	//RestartSec=5s
	//EnvironmentFile=/home/` + userName + `/web/` + domainName + `/public_html/.env
	//ExecStart=/home/` + userName + `/web/` + domainName + `/public_html/main
	//WorkingDirectory=/home/` + userName + `/web/` + domainName + `/public_html/
	//
	//[Install]
	//WantedBy=multi-user.target"
	//`

	//Write(serviceCreatedUbuntuV, serviceName+".service")
}

func (server *Dir) serviceRun(portNumber int64, serviceName, domainName, userName string) {
	serviceCreatedUbuntuV := `#!/bin/bash

workingfolder="/home/` + userName + `/web/` + domainName + `/public_html/"

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

	server.Write(serviceCreatedUbuntuV, "run.sh")
}

func (server *Dir) Write(value, filename string) {
	//https://linuxhint.com/create-file-golang/
	//path := filepath.Join("home", "ubuntu", "workspace", "newfile.txt")

	fmt.Println("chaa  " + server.mkdirFile)
	path := filepath.Join(server.mkdirFile, filename)
	fmt.Println(path)
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}

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
