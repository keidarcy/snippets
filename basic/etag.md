## Etag

### Use it in express

- [best explain](https://stackoverflow.com/a/67929691)
- pros
  - fast response
  - save bandwidth
- cons
  - hard to work with load balancer(extra etag header...) but can be solved by server(apache) configuration to make etag aware
  - hulu use it track user

```js
const express = require('express')();
express.get('/todo/1', (req, res) => {
  res.send({ id: 1, title: 'Get milk' });
});
```

```js
const etag = crypto.createHash('sha1').update(entity, 'utf8').digest('base64').substring(0, 27)
// 'MFvBQ/sYEFurbR/3xEt4Rikg7d4'
```

- generate based on [etag](https://github.com/jshttp/etag)

```sh
curl -I http://127.0.0.1:3000/todo/1

HTTP/1.1 200 OK
X-Powered-By: Express
Content-Type: application/json; charset=utf-8
Content-Length: 27
ETag: W/"1b-MFvBQ/sYEFurbR/3xEt4Rikg7d4" # W means weak
Date: Wed, 23 Feb 2022 09:34:06 GMT
Connection: keep-alive
Keep-Alive: timeout=5
```

- validate use [refresh](https://github.com/jshttp/fresh)

```sh
curl -H "if-none-match: W/\"1b-MFvBQ/sYEFurbR/3xEt4Rikg7d4\"" -I http://127.0.0.1:3000/todo/1

HTTP/1.1 304 Not Modified
X-Powered-By: Express
ETag: W/"1b-MFvBQ/sYEFurbR/3xEt4Rikg7d4"
Date: Wed, 23 Feb 2022 09:38:06 GMT
Connection: keep-alive
Keep-Alive: timeout=5
```
