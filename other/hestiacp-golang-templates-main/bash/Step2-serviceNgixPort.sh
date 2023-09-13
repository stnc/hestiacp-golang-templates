#!/bin/bash
# Adding php wrapper
user="$1"
domain="$2"
port="$3"



if [ ! -f "/home/${user}/conf/web/${domain}/nginx.hsts.conf" ]; then

echo "set \$go_web_port \"${port}\"; " > /home/${user}/conf/web/${domain}/nginx.hsts.conf

fi

exit 0
