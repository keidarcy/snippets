# Browser related javascript

- [chrome devtools](https://developers.google.com/web/tools/chrome-devtools/console/utilities)
- [web.dev](https://web.devG/)
- [how chrome work](https://developers.google.com/web/updates/2018/09/inside-browser-part1)
- [webpage tester](https://www.webpagetest.org/)
- [crux](https://web.dev/chrome-ux-report-data-studio-dashboard/)
- [url viewer](https://view.hugo-decoded.be/)

- [Browser related javascript](#browser-related-javascript)
  - [Event and CustomEvent](#event-and-customevent)
  - [Go to top](#go-to-top)
  - [Webpack setting for ts with dom](#webpack-setting-for-ts-with-dom)
  - [Upload image](#upload-image)
    - [Download file](#download-file)
  - [Cors](#cors)
    - [`fetch` example](#fetch-example)
    - [Express example](#express-example)
  - [Alphinejs fetch and render loop](#alphinejs-fetch-and-render-loop)
  - [DOM storage](#dom-storage)
    - [localstorage](#localstorage)
    - [sessionstorage](#sessionstorage)
    - [indexedDB](#indexeddb)
    - [Cookies](#cookies)
      - [get and set](#get-and-set)
      - [js-cookie](#js-cookie)
      - [get cookies object with vanilla js](#get-cookies-object-with-vanilla-js)
    - [Click others(usually toggle overlay)](#click-othersusually-toggle-overlay)
    - [Share to facebook and twitter](#share-to-facebook-and-twitter)
    - [Hiragana kanakana converter](#hiragana-kanakana-converter)
  - [axios](#axios)
    - [Create global axios instance](#create-global-axios-instance)
    - [Difference of `config` between get and post method](#difference-of-config-between-get-and-post-method)
  - [Get file name without extension](#get-file-name-without-extension)
  - [Date](#date)
  - [XMLHttpRequest](#xmlhttprequest)
  - [Beacon API](#beacon-api)
  - [JSONP](#jsonp)
  - [Leave alert](#leave-alert)
  - [Composition start|end|update event](#composition-startendupdate-event)
  - [script tag scope](#script-tag-scope)
  - [Web Vitals](#web-vitals)
    - [CLS](#cls)
      - [What causes](#what-causes)
      - [Layout shift score](#layout-shift-score)
      - [solutions](#solutions)
    - [LCP](#lcp)
      - [critical css](#critical-css)
      - [Server push(http2)](#server-pushhttp2)
    - [FID](#fid)
      - [what causes](#what-causes-1)

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

## axios

### Create global axios instance

```js
Axios.create({
  baseURL: BASE_URL,
  headers: {
    // get: {
    //   'Content-Type': 'application/json',
    //   'access-token': TOKEN
    // }
    'Content-Type': 'application/json',
    'access-token': TOKEN
  }
});
```

### Difference of `config` between get and post method

```js
const data = {
  sort: 'desc',
  count: 12
};

axios.post(URL, data);
axios.get(URL, { data });
```

## Get file name without extension

```js
arr.split('.').slice(0, -1).join('.');
```

## Date

```js
Date.now();
//1605615577351
new Date()(
  //Tue Nov 17 2020 21:19:42 GMT+0900 (Japan Standard Time)
  new Date()
).toString();
//"Tue Nov 17 2020 21:20:30 GMT+0900 (Japan Standard Time)"
Date.parse('Tue Nov 17 2020 21:20:30 GMT+0900 (Japan Standard Time)');
//1605615630000
new Date().getFullYear();
//2020
// GMT is an abbreviation for Greenwich Mean Time
```


## XMLHttpRequest

```js
const getBtn = document.getElementById('get-btn');
const postBtn = document.getElementById('post-btn');

const sendHttpRequest = (method, url, data) => {
  const promise = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open(method, url);

    xhr.responseType = 'json';
    if (data) {
      xhr.setRequestHeader('Content-Type', 'application.json');
    }
    xhr.onload = () => {
      if (xhr.status > 400) {
        reject('nonnon');
      }
      const data = xhr.response;
      resolve(data);
    };
    xhr.onerror = () => {
      reject('wrong');
    };
    xhr.send(JSON.stringify(data));
  });
  return promise;
};

getData = async () => {
  const data = await sendHttpRequest('GET', 'https://reqres.in/api/users');
  console.log({ data });
};

postData = async () => {
  try {
    const data = await sendHttpRequest('POST', 'https://reqres.in/ap/users', {
      email: 'test@test.com',
      password: 'testtest'
    });
    console.log({ data });
  } catch (ex) {
    console.error(ex);
  }
};

getBtn.addEventListener('click', getData);
postBtn.addEventListener('click', postData);
```

## Beacon API

- what it is

Beacon requests use the HTTP POST method and requests typically do not require a response, just return boolean value.

- why use it and difference between XMLHttpRequest

Ajax request using the XMLHTTPREQUEST object or even the fetch api can do way much more than the beacon request. But the ajax and fetch requests are waiting for result from the server which make them unresponsive in events such as window unload or page unload or even animations..

The beacon api only does a POST request and it immediately returns true if the request get piped for execution and false when it didnt. Plus it doesn’t wait for result from the server which makes it great for logging during animations or during page unload.. you send the request and you don’t wait for result.

- basic usage

```js
navigator.sendBeacon(`URL?timestay=${time}`);
```

- example

```js
document.addEventListener('unload', () => {
  navigator.sendBeacon(`URL?timestay=${time}`);
});
```

## JSONP

```html
<script>
  const cb = (res) => {
    console.log({ res });
  };
</script>
<script src="https://jsonplaceholder.typicode.com/todos?jsoncallback=cb"></script>
<script>
  console.log({ cb });
</script>
```

## Leave alert

- basic implementation

```js
window.addEventListener('beforeunload', ev => {
  if (pendingOps.size) {
    ev.retrunValue = 'strill waiting';
  }
})
```

- example with http request

```js
const pendingOps = new Set();

function addToPendingWork(promise) {
  const promise = fetch('/new-content');
  pendingOps.add(promise);
  spinner.hidden = false;

  const cleanup = () => {
    pendingOps.delete(promise);
    spinner.hidden = pendingOps.size === 0;
  };

  promise.then(cleanup).catch(cleanup);
}

window.addEventListener('beforeunload', ev => {
  if (pendingOps.size) {
    ev.retrunValue = 'strill waiting';
  }
})
```

- example with beacon api

```js
const data = JSON.stringify({ action: 'close', when: +new Data() });
window.addEventListener('beforeunload', (ev) => {
  navigator.sendBeacon('/analytics', data);
});
```

## Composition start|end|update event

- [mdn](https://developer.mozilla.org/en-US/docs/Web/API/Element/compositionstart_event)

```js
const inputElement = document.querySelector('input[type="text"]');

inputElement.addEventListener('compositionstart', (event) => {
  console.log(`generated characters were: ${event.data}`);
});


inputElement.addEventListener('compositionend', (event) => {
  console.log(`generated characters were: ${event.data}`);
});

inputElement.addEventListener('compositionupdate', (event) => {
  console.log(`generated characters were: ${event.data}`);
});
```
## script tag scope

```html
<script>
  var one = true;
  const two = true;
</script>
<script type="module">
  var three = true
  const four = true
</script>
<script type="module">
  console.log(self.one)
  console.log(two)
  // console.log(three)
  // console.log(four)
</script>
```

## Web Vitals

- LCP(Largest Content Paint) - Loading
- FID(First Input Delay) - Interactivity
- CLS(Cumulative Layout Shift) - Visual Stability
- FCP(First Contentful Paint)
- TTFB(Time to First Byte)

![LCP](https://webdev.imgix.net/vitals/lcp_ux.svg)
![FID](https://webdev.imgix.net/vitals/fid_ux.svg)
![CLS](https://webdev.imgix.net/vitals/cls_ux.svg)


### CLS

#### What causes

1. image without dimensions
2. ads, embeds, iframess without dimensions
3. dynamically injected content
4. web fonts causing FOIT/FOUT

#### Layout shift score

```
layout shift score = impact fraction * distance fraction
```

#### solutions

1. Always include width and height size attributes on images and videos.

```html
<!-- default css setting -->
<style>
img {
  aspect-ratio: attr(width) / attr(height)
}
</style>
<img src="src" width="640" height="360" alt="alt" />
```

2. Reserve enough space for dynamic content, like ads or promos. Avoid inserting new content above existing content, unless in response to a user interaction.

```html
<style>
.container {
  display: block;
  width: 720px;
  height: 90px;
  background: #ccc;
  overflow: hidden;
}
</style>

<div class="container">
  <iframe scr="...">
</div>
```

### LCP

#### critical css

```js
for (let i = 0; i < CssFiles.media.length; i++) {
  const link = document.createElement('link');
  link.href = CssFiles.media[i];
  link.type = 'text/css';
  link.ref = 'stylesheet';
  link.media = 'print';
  link.onload = link.media = 'all';
  document.getElementsByTagName('head')[0].appendChild(link);
}
```

#### Server push(http2)

### FID

#### what causes

1. Long taks
2. Long javascript execution time
3. Large javascript bundles
4. render-blocking javascript