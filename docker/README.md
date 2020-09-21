 - [Play with docker](https://labs.play-with-docker.com/)
 

## Basic concept

[figma](https://www.figma.com/file/GAMKg6zWYqYId04ICOHOPq/funny?node-id=2%3A3)

 - pull
 ```
 docker pull nginx
 ```
 
  - remove image
 ```
 docker rmi {{image name}}
 ```
 
 - run
 ```
 docker run -d -p 8000:80 nginx
 docker run -d -p 8000:80 --name mynginx -v `pwd`:/usr/share/nginx/html nginx:1.13
 ```
 
  - remove container
 ```
 docker rm -f {{container name}}
 ```

 - commit
 ```
 docker commit {{running container name}} {{new image name}}
 ```
 
 - save
 ```
 docker save {{image name}} > {{tar file name}}
 ```
 
 - load
 ```
 docker load < {{tar file name}}
 ```
 
 ## Docker network
 
 ```
 docker run -dit --link mynginx:mynginx alpine
 ```
 
 > mynginx becomes domain
  
 ```
 cat /etc/hosts
 
 #172.17.0.2      mynginx 9b003a399629
 #172.17.0.3      fa05ac17834a

 ```

 - Simple docker-compose
 
```yml
version: "3"
services:
  nginx:
    image: nginx:alpine
    ports:
    - 80:80
    volumes:
    - /root/html:/usr/share/nginx/html
    - /root/conf/nginx.conf:/etc/nginx/nginx.conf
  php:
    image: devilbox/php-fpm:5.2-work-0.89
    volumes:
    - /root/html:/var/www/html
  mysql:
    image: mysql:5.7
    environment:
    - MYSQL_ROOT_PASSWORD=123456
```

## Dockerfile

 - basic
```
FROM alpine
WORKDIR /app
COPY src/ /app
RUN echo 321 >> 1.txt
CMD tail -f 1.txt
```

- similar
```
COPY vs ADD (ADD for url resource)
ENTRYPOINT vs CMD
```
- intermediate
```
EXPOSE
VOLUMN /a/b
ENV
ARG
LABEL
```
- example

```
FROM alpine
LABEL key="value" key1="value1"
WORKDIR /app
COPY src/ /app
RUN echo 321 >> 1.txt
ONBUILD ENV C=100
CMD echo $c
```

- Unusual

```
STOPSIGNAL # INT KILL .etc
HEALTHCHECK
SHELL
USER
```
 
## Reset All the Local Resources

```
docker system prune -a
```
 - [basic help gist](https://gist.github.com/bradtraversy/89fad226dc058a41b596d586022a9bd3)
 - [wp example gist](https://gist.github.com/bradtraversy/faa8de544c62eef3f31de406982f1d42)


 
