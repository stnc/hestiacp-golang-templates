server {
   listen      %ip%:%web_port%;
   server_name %domain_idn% %alias_idn%;
  
   include %home%/%user%/conf/web/%domain%/nginx.hsts.conf*;

   #include %home%/%user%/conf/web/%domain%/nginx.forcessl.conf*;

   location / {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header Host $http_host;
        proxy_pass  http://%ip%:$go_web_port;
    }



    location @fallback {
        proxy_pass  https://%ip%:$go_web_port;
    }


    location ~ /\.(?!well-known\/|file) {
       deny all;
       return 404;
    }



    location ~ /\.ht    {return 404;}
    location ~ /\.svn/  {return 404;}
    location ~ /\.git/  {return 404;}
    location ~ /\.hg/   {return 404;}
    location ~ /\.bzr/  {return 404;}

    location = /robots.txt {
        allow all;
        log_not_found off;
        access_log off;
    }


    location /error/ {
        alias   %home%/%user%/web/%domain%/document_errors/;
    }

    location /vstats/ {
        alias   %home%/%user%/web/%domain%/stats/;
        include %home%/%user%/web/%domain%/stats/auth.conf*;
    }

    proxy_hide_header Upgrade;

    include /etc/nginx/conf.d/phpmyadmin.inc*;
    include /etc/nginx/conf.d/phppgadmin.inc*;
    include %home%/%user%/conf/web/%domain%/nginx.conf_*;
}
