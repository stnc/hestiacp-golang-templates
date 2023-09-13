sudo nano /lib/systemd/system/goweb.service

# service restart 
service nginx restart
service goweb restart


# nginx hata bakma
tail -f /var/log/nginx/error.log
