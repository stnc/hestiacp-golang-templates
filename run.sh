workingfolder="/home/admin/web/ert.com/public_html/"

cd $workingfolder

echo "server: ert.com
port: 9090
service: ert.com" > /home/admin/web/ert.com/public_html/data.json

chmod +rwx main

systemctl enable goweb.service

systemctl start goweb

sudo systemctl daemon-reload

sudo systemctl restart goweb

sudo systemctl restart nginx

exit 0