## No SSL

```
server {
  listen       80;
  server_name  domain name;

  location / {
    proxy_pass http:// ip address : port;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection 'upgrade';
    proxy_set_header Host $host;
    proxy_cache_bypass $http_upgrade;
  }
}
```

## With SSL
```
server {
        listen 80;
        server_name domain name;
        return 301 https://$host$request_uri;
}

server {
listen 443 ssl;

  server_name _
ssl_certificate /etc/letsencrypt/live/domain name/fullchain.pem; # managed by Certbot
ssl_certificate_key /etc/letsencrypt/live/domain name/privkey.pem; # managed by Certbot

    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
    ssl_prefer_server_ciphers on;
    ssl_ciphers 'EECDH+AESGCM:EDH+AESGCM:AES256+EECDH:AES256+EDH';

 location / {
        proxy_pass http://ip address : port;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```
