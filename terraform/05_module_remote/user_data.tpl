#!/bin/bash
yum update -y
yum install -y httpd
systemctl start httpd
systemctl enable httpd
echo "Hello World from ${name} again" >/var/www/html/index.html
