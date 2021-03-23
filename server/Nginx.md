# Nginx

- [nginx.org docs](https://nginx.org/en/docs/)
- [nginx.com docs](https://docs.nginx.com/)

- [Nginx](#nginx)
  - [basic](#basic)
  - [Layer 7 proxy](#layer-7-proxy)
  - [Layer 4 proxy](#layer-4-proxy)
  - [Location Directive](#location-directive)
  - [Reverse Proxy](#reverse-proxy)
  - [Load Balancer](#load-balancer)
  - [Rewrite (redirect)](#rewrite-redirect)
      - [`example.com` => `www.example.com`](#examplecom--wwwexamplecom)
      - [`www.example.com/test1` => ``www.example.com/test2`](#wwwexamplecomtest1--wwwexamplecomtest2)
      - [`http` => `https`](#http--https)
      - [Static content](#static-content)
  - [SSL](#ssl)
    - [NO ssl](#no-ssl)
    - [With SSL](#with-ssl)
  - [Default](#default)

## basic

```
http {
    server {
        listen 8080;
        root /etc/nginx/html;

        location /images {
            root /etc/nginx/;
        }

        location ~ .png$ {
            return 403;
        }
    }

    server {
        listen 8888;

        location / {
            proxy_pass http://localhost:8080/;
        }
    }
}

events {}

```

## Layer 7 proxy

```
http {
    upstream allbackend {
        # ip_hash;
        server docker.for.mac.localhost:2222;
        server docker.for.mac.localhost:3333;
        server docker.for.mac.localhost:4444;
        server docker.for.mac.localhost:5555;
    }
    upstream app1backend {
        server docker.for.mac.localhost:2222;
        server docker.for.mac.localhost:3333;

    }

    upstream app2backend {
        server docker.for.mac.localhost:4444;
        server docker.for.mac.localhost:5555;
    }
    server {
        listen 80;

        location / {
            proxy_pass http://allbackend/;
        }

        location /app1 {
            proxy_pass http://app1backend/;
        }
        location /app2 {
            proxy_pass http://app2backend/;
        }

        location /admin {
            return 403;
        }
    }
}

events { }
```

## Layer 4 proxy

```
stream {
    upstream allbackend {
        server docker.for.mac.localhost:2222;
        server docker.for.mac.localhost:3333;
        server docker.for.mac.localhost:4444;
        server docker.for.mac.localhost:5555;
    }

    server {
        listen 80;
        proxy_pass allbackend;
    }
}

events { }
```

`docker run --rm -v `pwd`/nginx.conf:/etc/nginx/nginx.conf -v `pwd`/ssl:/etc/nginx/ssl -p 9999:443 --name hello_nginx -d nginx:alpine`

## Location Directive

> Priority high => low

```sh
http {
    sever {
        location = /a {
          echo "=/a";
        }

        location ^~ /a {
          echo "^~ /a";
        }

        location ~ /\w {
          echo "/\w";
        }

        location / {
          echo ="/";
        }
    }
}
```

## Reverse Proxy

```sh
http {
    server {
        listen    80;
        server_name localhost;
        default_type   text/html;

        location /a {
            proxy_pass http://192.168.0.12:80;
        }

        location /b/ {
            proxy_pass http://192.168.0.12:81/;
        }
    }
}
```

/a/** => http://192.168.0.12:80/a/**;
/b/** => http://192.168.0.12:81/**;

## Load Balancer

```
 http {

    upstream ANY_STRING {
        server 192.168.0.12:80;
        server 192.168.0.12:81;
    }
    //upstream ANY_OTHER_STRING {
    //  server 192.168.0.12:80 weight=10;
    //  server 192.168.0.12:81 weight=1;
    //}
    server {
        listen    80;
        server_name localhost;
        default_type   text/html;

        location / {
            echo "/";
        }

        location /a {
            proxy_pass http://ANY_STRING/;
        }
    }
}
```

## Rewrite (redirect)

`rewrite regex replacement [flag];`
_flag_ type _LAST_, _BREAK_, _PERMANENT_(status code 301), _REDIRECT_(status code 302)

- examples

#### `example.com` => `www.example.com`

```
server {
  listen 80;
  server_name example.com;

  rewrite ^(.*)$ http://www.example.com$1 permanent;
}
```

#### `www.example.com/test1` => ``www.example.com/test2`

```
server {
  listen 80;
  server_name www.example.com;

  rewrite ^/test1(.*)$ http://www.example.com/test2$1 permanent;
}
```

#### `http` => `https`

```
server {
  listen 80;
  server_name www.example.com;

  if ($http_x_forwarded_proto != https) {
    rewrite ^(.*)$ https://www.example.com$1 permanent;
  }
}
```

#### Static content

```
location ~ \.(png|jpeg|jpg|js|css|woff|ttf)$ {
    expires 1h;
}
```

or

```
location ~* ^/static/.+\.(png|whatever-else)$ {
    alias /var/www/some_static;
    expires 24h;
}
location / {
    # regular rules
}
```

## SSL

### NO ssl

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

### With SSL

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

## Default

```
# For more information on configuration, see:
#   * Official English Documentation: http://nginx.org/en/docs/
#   * Official Russian Documentation: http://nginx.org/ru/docs/

user nginx;
worker_processes auto;
error_log /var/log/nginx/error.log;
pid /run/nginx.pid;

# Load dynamic modules. See /usr/share/doc/nginx/README.dynamic.
include /usr/share/nginx/modules/*.conf;

events {
    worker_connections 1024;
}

http {
    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile            on;
    tcp_nopush          on;
    tcp_nodelay         on;
    keepalive_timeout   65;
    types_hash_max_size 4096;

    include             /etc/nginx/mime.types;
    default_type        application/octet-stream;

    # Load modular configuration files from the /etc/nginx/conf.d directory.
    # See http://nginx.org/en/docs/ngx_core_module.html#include
    # for more information.
    include /etc/nginx/conf.d/*.conf;

    server {
        listen       80;
        listen       [::]:80;
        server_name  _;
        root         /usr/share/nginx/html;

        # Load configuration files for the default server block.
        include /etc/nginx/default.d/*.conf;

        error_page 404 /404.html;
            location = /40x.html {
        }

        error_page 500 502 503 504 /50x.html;
            location = /50x.html {
        }
    }

# Settings for a TLS enabled server.
#
#    server {
#        listen       443 ssl http2;
#        listen       [::]:443 ssl http2;
#        server_name  _;
#        root         /usr/share/nginx/html;
#
#        ssl_certificate "/etc/pki/nginx/server.crt";
#        ssl_certificate_key "/etc/pki/nginx/private/server.key";
#        ssl_session_cache shared:SSL:1m;
#        ssl_session_timeout  10m;
#        ssl_ciphers PROFILE=SYSTEM;
#        ssl_prefer_server_ciphers on;
#
#        # Load configuration files for the default server block.
#        include /etc/nginx/default.d/*.conf;
#
#        error_page 404 /404.html;
#            location = /40x.html {
#        }
#
#        error_page 500 502 503 504 /50x.html;
#            location = /50x.html {
#        }
#    }

}
```
