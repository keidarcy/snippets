# Browser related javascript

## Event and CustomEvent

```ts
$0.dispatchEvent(new Event('newEvent'))
$0.dispatchEvent(new CustomEvent('newEvent', { detail: 'my detail' })

```

## Go to top

```ts
export const toTop = () => {
  document.body.scrollTop = 0; // For Safari
  document.documentElement.scrollTop = 0;
};
```

## Webpack setting for ts with dom

```js
// webpack.config.js
module.exports = {
  mode: 'development',
  entry: './src/index.ts',
  output: {
    path: `${__dirname}/dist`,
    filename: 'bundle.js'
  },
  devtool: 'eval-source-map', // for source map
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: 'ts-loader'
      }
    ]
  },
  resolve: {
    extensions: ['.ts', '.js']
  }
};
```

## Upload image

```ts
const fd = new FormData();
fd.append('image', file, file.name);
const res = await axios.post(apiPath, fd);
```

### Download file

```js
const filename = '1.txt';
const a = document.createElement('a');
const blob = new Blob(['hello world!']);
a.download = filename;
a.href = URL.createObjectURL(blob);
a.click();
URL.revokeObjectURL(blob);
```

## Cors

`example with express and react` - [repo](https://github.com/keidarcy/cors-node)

### `fetch` example

- from different origin
  [Access-Control-Allow-Origin] must be the address

- with headers
  if request header is added, [Access-Control-Allow-Headers] must be allowed

- with credentials
  [withCredentials=true] or `fetch('https://example.com', {credentials: 'include'})`=> [Access-Control-Allow-Credentials] must be configured

- simple and complex request
  simple `fetch`, add `Access-Control-Allow-Origin`,
  complex `fetch` like with header or `DELETE` method, config `options` first then add `Access-Control-Allow-Headers`.

### Express example

say express run on http:localhost:3001, and react runs on http://localhost:3000.

1. Add header setting to allow.

```ts
router.get('', async (req: Request, res: Response) => {
  res.header('Access-Control-Allow-Origin', 'http://localhost:3000');
  res.header('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE');
  res.header('Access-Control-Allow-Headers', 'Content-Type, Authorization, access_token');
  const products = await Product.find({});
  res.json(products);
});
```

2. Add `proxy` tp package.json of frontend project.

```
"proxy": "http://127.0.0.1:3000",
```

3. Add `cors` middleware to handler.

`yarn add cors`

```ts
const express = require('express');
const cors = require('cors');
const app = express();

app.use(cors());
```

OR

```ts
const express = require('express');
const cors = require('cors');
const app = express();

app.get('/user/:userId', cors(), function (req, res, next) {
  res.json({ result: 'Possible to access from any where' });
});
```

OR add `cors` configuration

- option:`origin、methods、allowedHeaders、exposedHeaders、credentials、maxAge、preflightContinue、optionsSuccessStatus`.

```ts
const express = require('express');
const cors = require('cors');
const app = express();

const corsOptions = {
  origin: 'http://example.com',
  optionsSuccessStatus: 200
};

app.get('/user/:userId', cors(corsOptions), function (req, res, next) {
  res.json({ msg: 'Possible to access from http://example.com' });
});
```

## Alphinejs fetch and render loop

```html
<div x-data="gistsData()" x-init="init()">
  <ul>
    <template x-for="gist in gists" :key="gist.id">
      <li>
        <h1 x-text="gist.public"></h1>
        <a x-bind:href="gist.html_url" class="font-bold">url</a>
        <br />
        <small x-text="gist.description"></small>
      </li>
    </template>
  </ul>
</div>
<script>
  function gistsData() {
    return {
      title: 'Latest Gists',
      gists: [],
      init() {
        fetch('https://api.github.com/users/keidarcy/gists')
          .then((response) => response.json())
          .then((response) => {
            this.gists = response;
          });
      }
    };
  }
</script>
```

## DOM storage

- cookies 4kb
- indexedDB 250MB per domain
- localstorage | sessionstorage 5mb per domain

### localstorage

- hold when browser close
- cross in same origin

### sessionstorage

- lose when browser close
- not cross even same origin

### indexedDB

- [Dexie](https://github.com/dfahlander/Dexie.js/)

```js
const db = new Dexie('mydb');
db.version(1).store({ person: '++id, name, age' });

db.person.add({ name: 'john', age: 20 });
db.person.add({ name: 'mary', age: 23, email: 'd@e.com' });

db.person.put({ name: 'john', age: 28 });

await db.person.get(1);

await db.person.where('age').above(30).toArray(); // age > 30

db.close();
```

### Cookies

- in request header everytime

#### get and set

```js
document.cookie;

document.cookie = 'key=value';
```

#### js-cookie

- [js-cookie](https://github.com/js-cookie/js-cookie)

```js
import Cookies from 'js-cookie';

Cookies.set('foo', 'bar');
Cookies.set('name', 'value', { expires: 7, path: '' });
Cookies.get('name'); // => 'value'
Cookies.get('nothing'); // => undefined
```

#### get cookies object with vanilla js

```js
let cookies = document.cookie
  .split(';')
  .map((cookie) => cookie.split('='))
  .reduce(
    (accumulator, [key, value]) => ({
      ...accumulator,
      [key.trim()]: decodeURIComponent(value)
    }),
    {}
  );
```

### Click others(usually toggle overlay)

```js
  $('body').on('click', function (event) {
    if (
      !$.contains(document.querySelector('.js-sidebar'), $(event.target)[0]) &&
      sidebarShow
    ) {
      sidebar.slideLeft(500);
      sidebarShow = false;

```

### Share to facebook and twitter

```js
window.open('https://twitter.com/share?url=' + encodeURIComponent(document.URL));
window.open('http://www.facebook.com/share.php?u=' + encodeURIComponent(location.href));
```

### Hiragana kanakana converter

```ts
const kanaToHira = (str: string) =>
  Array.from(str)
    .map((letter) =>
      letter.replace(/[\u30a1-\u30f6]/g, (match) =>
        String.fromCharCode(match.charCodeAt(0) - 0x60)
      )
    )
    .join('');

const hiraToKana = (str: string) =>
  Array.from(str)
    .map((letter) =>
      letter.replace(/[\u3041-\u3096]/g, (match) =>
        String.fromCharCode(match.charCodeAt(0) + 0x60)
      )
    )
    .join('');
```
