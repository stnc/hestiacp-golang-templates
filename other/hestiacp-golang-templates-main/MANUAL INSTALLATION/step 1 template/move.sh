#!/bin/bash

mv golang.tpl /usr/local/hestia/data/templates/web/nginx/php-fpm

mv golang.stpl /usr/local/hestia/data/templates/web/nginx/php-fpm

sudo systemctl daemon-reload

exit 0
