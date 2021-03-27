- [Quick server](#quick-server)
- [Development enviroment bar](#development-enviroment-bar)
- [Python watch dog](#python-watch-dog)
- [tar gz](#tar-gz)
  - [compress](#compress)
    - [tar.gz](#targz)
    - [tar](#tar)
  - [archive](#archive)
    - [tar.gz](#targz-1)
    - [tar.gz](#targz-2)
- [find out listening port](#find-out-listening-port)

## Quick server

```
#  python 3.x
python3 -m http.server

# php
php -S localhost:8080

# ruby
ruby -run -e httpd . -p 8080
```

## Development enviroment bar

```html
<div
  style="background-color:#ff0;color:#111;padding:10px;text-align:center;font-size:12px;font-weight:bold;"
>
  DEVELOPMENT
</div>
```

## Python watch dog

- [when-changed](https://github.com/joh/when-changed)

```
-r recursively
-v verbose output
-1 don't re-run when running
-s start immediately
-q run quietly
```

```bash
when-changed -r -v -1 . python3 init.py
```

## tar gz

### compress

#### tar.gz

```sh
tar -zcvf filename.tar.gz directoryname
```

#### tar

```sh
tar -cvf filename.tar directoryname
```

### archive

#### tar.gz

```sh
tar -zxvf filename.tar.gz
```

#### tar.gz

```sh
tar -xvf filename.tar
```


## find out listening port

```sh
sudo lsof -i -P -n | grep LISTEN
sudo netstat -tulpn | grep LISTEN
sudo lsof -i:22 ## see a specific port such as 22 ##
sudo nmap -sTU -O IP-address-Here
```