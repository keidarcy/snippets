# snippets

## js/ts

#### Nodejs

 - `Ctrl + L` clean screen, similar to `process.stdout.write("\u001b[2J\u001b[0;0H");`

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
### Node 

#### Cli Parameters

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
### Download file

```js
const filename = '1.txt'
const a = document.createElement('a')
const blob = new Blob(['hello world!'])
a.download = filename
a.href = URL.createObjectURL(blob)
a.click()
URL.revokeObjectURL(blob)
```

### Cors

```
[withCredentials=true] => [Access-Control-Allow-Origin] must be the address

[withCredentials=true] => [Access-Control-Allow-Credentials] must be configured

if request header is added, [Access-Control-Allow-Headers] must be allowed

```

## MySql
```
# bin
/usr/local/opt/mysql@5.7/bin/mysql
```

```
show global variables like "%datadir%";
```

## Nginx

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


## Git

 - Update forked repository to original repository latest

```
git remote add upstream https://github.com/original/repository.git

git fetch upstream

git rebase upstream/master

git push origin master --force
```

 - Clean current branch
 
```
nah='git reset --hard;git clean -df;'
```


## Regex

 - [regexr](https://regexr.com/)

### Basic
 
 - `.` => Matches any character except line breaks.
 - `^` => Begining of the string.
 - `$` => End of the string.
 - `[^]` => Negataed set.
 - `[]{3, }` => Match between 3 to unlimited characters.
 - `\d <=> [0-9]`, `\D <=> [^0-9]`, `\w <=> [a-zA-Z0-9_]` // (alphanumeric & underscore).
 - `\s <=> [\r\n\t\v ]` => Matches any whitespace character (spaces, tabs, line breaks).
 - `{0,1} <=> ?` => Match between 0 and 1 times.
 - `{0, } <=> *` => Match between 0 to unlimited times.
 - `{1, } <=> +` => Match between 1 to unlimited times.

### Group

 - use group
 ```
 `xx@gmail.com`.match(/^([a-zA-Z0-9]\w*)@gmail\.com$/)[1] // 'xx'
 ```
 
 - rename group
 ```
 `xx@gmail.com`.match(/^(?<first>[a-zA-Z0-9]\w*)@gmail\.com$/).groups.first // 'xx'
 ```
 
 - reuse group in regex
 ```
 ^(\d\d)\1$    // match 1212
 ```
 
 - reuse group with renamed
 
 ```
 ^(?<first>\d+)\k<first>$
 ```
 
 - `(?=` Positive lookahead
 
 ```js
 'foobar, foopoo'.replace(/foo(?=bar)/g, 'replaced') // "replacedbar, foopoo"
 ```
 - `(?!` Negative lookahead
 
 ```js
 'foobar, foopoo'.replace(/foo(?!bar)/g, 'replaced') // "foobar, replacedpoo"
 ```
 
 - `(?<=` Positive lookbehiend
 
 ```js
'foobar, foopoo'.replace(/(?<=foo)bar/g, 'replaced') // "fooreplaced, foopoo"
 ```
 
  - `(?<!` Negative lookbehiend
 
 ```js
'foobar, foopoo'.replace(/(?<=foo)bar/g, 'replaced') // "foobar, fooreplaced"
 ```
 
 
### Examples

 - match email address
`^[a-zA-Z0-9]\w*@gmail\.com$` => `xx@gmail.com`

 - match float string and get fixed 
`'rgba(100,150,200,.72312332)'.replace(/(\.\d{2})[0-9]*/,"$1"); // rgba(100,150,200,.72)`

 - match two separate part
 `'/api/xxxx/edit?uu=xxxx&id=1&type=&page='.replace(/\/api|edit.*/g, '') // /xxxx/`


## Shell

#### `node (eval):1: command not found: _node` zsh problem
 - update omz `upgrade_oh_my_zsh`
 - delete all caches `rm ~/.zcompdump*`

#### ssh in pi from mac without password

```
# from mac
scp pi_rsa.pub pi@0.0.0.000:/home/pi/.ssh

# in pi
cat pi_rsa.pub >> authorized_keys
```

## [Xpath](http://xpather.com/)
