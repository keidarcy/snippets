- [Quick server](#quick-server)
- [Development enviroment bar](#development-enviroment-bar)
- [Python watch dog](#python-watch-dog)

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