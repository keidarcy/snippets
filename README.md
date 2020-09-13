# snippets

#### Generate zip file from string

```ts
import jszip from 'jszip';

const createZip = async (code: string): Promise<void> => {
  const zip = new jszip();
  zip.file('index.js', code);
  const content = zip.generateAsync({ type: 'string' });
  console.log(content);
  zip
    .generateNodeStream({ type: 'nodebuffer', streamFiles: true })
    .pipe(fs.createWriteStream('./dist/out.zip'))
    .on('finish', function () {
      console.log('out.zip written.');
    });
};

```

#### TS decorator

```ts
function time(name: string) {
    return function (target, propertyKey: string, descriptor: PropertyDescriptor) {
        const fn = descriptor.value;
        descriptor.value = (...args) => {
            console.time(name);
            const v = fn(...args);
            console.timeEnd(name);
            return v;
        }
    }
}

class C{
    @time('C.method')
    method(name: string){
        console.log('method called', name)
        for (let i = 0; i < 1000000000; i++){}
    }
}

new C().method('koko')
```
#### Node cli parameter

```ts
#! /usr/bin/env node

const args = process.argv.slice(2);
const mydir = args[0];
const myname = args[1].split('--name=')[1];
console.log({ mydir, myname });
```
```
input `./dist/index.js myqpp --name=fun`
output `{ mydir: 'myqpp', myname: 'fun' }`
```

#### MySql
```
# bin
/usr/local/opt/mysql@5.7/bin/mysql
```

```
show global variables like "%datadir%";
```

#### Nginx

 - Location Directive
 
>Priority high => low

```
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

 - Reverse Proxy
```
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

 - Load Balancer
 
```
 http {
 
    upstream_group1 {
        server 192.168.0.12:80;
        server 192.168.0.12:81;
    }
    //upstream_group1 {
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
            proxy_pass http://group1/;   
        }
    }
}
```


### Git

 - Update forked repository to original repository latest

```
git remote add upstream https://github.com/original/repository.git

git fetch upstream

git rebase upstream/master

git push origin master --force
```

### Cors

```
[withCredentials=true] => [Access-Control-Allow-Origin] must be the address

[withCredentials=true] => [Access-Control-Allow-Credentials] must be configured

if request header is added, [Access-Control-Allow-Headers] must be allowed

```
