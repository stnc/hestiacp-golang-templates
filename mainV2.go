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

type Config struct {
	mkdirFile string
}

func (configuration *Config) setName(fileName string) {
	configuration.mkdirFile = fileName
}

func prompt() {
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

	mkdir := Config{}
	mkdir.setName(domainName) // changed
	fmt.Println(mkdir.mkdirFile)
	if err := mkdir.ensureDir(); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}
	mkdir.fileCreatedNginxHst(portNumber)
	mkdir.serviceCreatedUbuntu(portNumber, serviceName, domainName, userName)
	mkdir.serviceRun(portNumber, serviceName, domainName, userName)
}

func ready() {
	mkdir := Config{}
	var portNumber int64
	var serviceName, domainName, userName string
	portNumber = 9090
	serviceName = "selweb"
	domainName = "selweb.com"
	userName = "admin"

	mkdir.setName(domainName) // changed

	fmt.Println(mkdir.mkdirFile)
	if err := mkdir.ensureDir(); err != nil {
		fmt.Println("Directory creation failed with error: " + err.Error())
		os.Exit(1)
	}

	mkdir.fileCreatedNginxHst(portNumber)
	mkdir.serviceCreatedUbuntu(portNumber, serviceName, domainName, userName)
	mkdir.serviceRun(portNumber, serviceName, domainName, userName)
	mkdir.template("ssl")
	mkdir.template("no")
	mkdir.templateSSH()
}
func main() {
	//prompt()
	ready()

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

func (configuration *Config) ensureDir() error {
	var dirName string
	dirName = configuration.mkdirFile
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

func (configuration *Config) fileCreatedNginxHst(portNumber int64) {
	fileCreatedNginxHstV := `set $go_web_port  ` + strconv.FormatUint(uint64(portNumber), 10) + `;`
	configuration.Write(fileCreatedNginxHstV, "nginx.hsts.conf")
}

func (configuration *Config) serviceCreatedUbuntu(portNumber int64, serviceName, domainName, userName string) {

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
	configuration.Write(serviceCreatedUbuntuV, serviceName+".service")

}
func (configuration *Config) templateSSH() {
	serviceCreatedUbuntuV := `#!/bin/bash

cp golang.tpl /usr/local/hestia/data/templates/web/nginx/php-fpm

cp golang.stpl /usr/local/hestia/data/templates/web/nginx/php-fpm

exit 0`

	configuration.Write(serviceCreatedUbuntuV, "template.sh")
}

func (configuration *Config) serviceRun(portNumber int64, serviceName, domainName, userName string) {

	serviceCreatedUbuntuV := `#!/bin/bash

cp /home/` + userName + `/conf/web/` + domainName + ` nginx.hsts.conf

cp  /lib/systemd/system  ` + serviceName + `.service

workingfolder="/home/` + userName + `/web/` + domainName + `/public_html/"

cd $workingfolder

echo "configuration: 
domain name: ` + domainName + `
port: ` + strconv.FormatUint(uint64(portNumber), 10) + `
service: ` + serviceName + `" > /home/` + userName + `/web/` + domainName + `/public_html/data.json

chmod +rwx main

systemctl enable ` + serviceName + `.service

systemctl start ` + serviceName + `

sudo systemctl daemon-reload

sudo systemctl restart ` + serviceName + `

sudo systemctl restart nginx

exit 0`

	configuration.Write(serviceCreatedUbuntuV, "run.sh")
}

func (configuration *Config) template(types string) {
	var serviceCreatedUbuntuV string
	var sslCertificate string
	var locationSSL string
	var listen string
	var fallback string
	var conf string
	var tplName string

	if types == "ssl" {
		listen = `listen      %ip%:%web_ssl_port% ssl http2;`
		sslCertificate = `ssl_certificate      %ssl_pem%;
   ssl_certificate_key  %ssl_key%;
   ssl_stapling on;
   ssl_stapling_verify on;
   #ssl_verify_client off;`

		locationSSL = `location / {
		proxy_set_header Host $http_host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Forwarded-Proto $scheme;
		proxy_set_header VERIFIED $ssl_client_verify;
		proxy_set_header DN $ssl_client_s_dn;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto: https;
		#proxy_ssl_verify off;
		proxy_pass  https://%ip%:$go_web_port;   
	}`
		fallback = `location @fallback {
        proxy_pass  https://%ip%:$go_web_port;
    }`
		conf = `include %home%/%user%/conf/web/%domain%/nginx.ssl.conf_*;`
		tplName = "stpl"
	} else {
		listen = `listen      %ip%:%web_port%;`

		sslCertificate = ``

		locationSSL = `location / {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header Host $http_host;
        proxy_pass  http://%ip%:$go_web_port;
    }`
		fallback = `location @fallback {
        proxy_pass  https://%ip%:$go_web_port;
    }`
		conf = `include %home%/%user%/conf/web/%domain%/nginx.conf_*;`
		tplName = "tpl"
	}

	serviceCreatedUbuntuV = `server {
   ` + listen + `
   server_name %domain_idn% %alias_idn%;
   root        %docroot%;

   access_log  /var/log/nginx/domains/%domain%.log combined;
   access_log  /var/log/nginx/domains/%domain%.bytes bytes;
   error_log   /var/log/nginx/domains/%domain%.error.log error;

   ` + sslCertificate + `
  
   include %home%/%user%/conf/web/%domain%/nginx.hsts.conf*;

   #include %home%/%user%/conf/web/%domain%/nginx.forcessl.conf*;


  ` + locationSSL + `

  ` + fallback + `

    location ~ /\.ht    {return 404;}
    location ~ /\.svn/  {return 404;}
    location ~ /\.git/  {return 404;}
    location ~ /\.hg/   {return 404;}
    location ~ /\.bzr/  {return 404;}


    location = /favicon.ico {
        log_not_found off;
        access_log off;
    }


    location = /robots.txt {
        allow all;
        log_not_found off;
        access_log off;
    }

    location ~ /\.(?!well-known\/|file) {
       deny all;
       return 404;
    }


    location /error/ {
        alias   %home%/%user%/web/%domain%/document_errors/;
    }

    location /vstats/ {
        alias   %home%/%user%/web/%domain%/stats/;
        include %home%/%user%/web/%domain%/stats/auth.conf*;
    }


    include /etc/nginx/conf.d/phpmyadmin.inc*;
    include /etc/nginx/conf.d/phppgadmin.inc*;
    ` + conf + `
}
`

	configuration.Write(serviceCreatedUbuntuV, "golang."+tplName+"")

}

func (configuration *Config) Write(value, filename string) {
	//https://linuxhint.com/create-file-golang/
	//path := filepath.Join("home", "ubuntu", "workspace", "newfile.txt")

	fmt.Println("mkdirFile  " + configuration.mkdirFile)
	path := filepath.Join("export", configuration.mkdirFile, filename)
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
