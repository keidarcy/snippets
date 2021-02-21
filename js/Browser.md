# Browser related javascript

- [chrome devtools](https://developers.google.com/web/tools/chrome-devtools/console/utilities)

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
  - [WebRTC](#webrtc)
  - [WebSocket & WebTransport](#websocket--webtransport)
    - [XMLHttpRequest](#xmlhttprequest)

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

## WebRTC

1. connect two browsers(A & B)
2. A will create an offer(sdp) and set it as local description
3. B will get the offer and set it as remote description
4. B creates an answer sets it as its local description and signal the answer(sdp) to A
5. A sets the answer as its remote description
6. Connection established, exchange data channel

- A

```js
const lc = new RTCPeerConnection(); //locationconnection
const dc = lc.createDataChannel('channel');
dc.onmessage = (e) => console.log('Got a message' + e.data);
dc.onopen = (e) => console.log('connection opened!');
lc.onicecandidate = (e) =>
  console.log('new ice candidate! reprinting SDP' + JSON.stringify(lc.localDescription));
lc.createOffer()
  .then((o) => lc.setLocalDescription(o))
  .then((a) => console.log('set successfully'));
// new ice candidate! reprinting SDP{"type":"offer","sdp":......
// copy JSON object as offer to B
```

- B

```js
let dc;
const offer = ''; //{"type":"offer","sdp":......}"
const rc = new RTCPeerConnection();
rc.onicecandidate = (e) =>
  console.log('new ice candidate! reprinting SDP' + JSON.stringify(rc.localDescription));
rc.ondatachannel = (e) => {
  dc = e.channel;
  dc.onmessage = (e) => console.log('New message from A' + e.data);
  dc.onopen = (e) => console.log('connection opened!!');
};
rc.setRemoteDescription(offer).then((a) => console.log('offset set'));
rc.createAnswer()
  .then((a) => rc.setLocalDescription(a))
  .then((a) => console.log('answer created'));
```

- A

```js
const answer = ''; //
lc.setRemoteDescription(answer);
dc.send('yoyo from A');
// Got a messageyoyo from B
```

- B

```js
// New message from Ayoyo A
dc.send('yoyo from B');
```

- `media`, `carmera` will use `lc.addTract`

## WebSocket & WebTransport

- server node

```js
const WebSocket = require('ws');
const server = new WebSocket.Server({ port: '8000' });

server.on('connection', (socket) => {
  socket.on('message', (message) => {
    socket.send(message);
  });
});
```

- client browser

```js
const socket = new WebSocket('ws://localhost:8000');
socket.onmesssage = ({ data }) => {
  console.log('Message from server', data);
};
document.querySelector('button').onclick = () => {
  socket.send('hello');
};
```

### XMLHttpRequest

```js
const getBtn = document.getElementById('get-btn');
const postBtn = document.getElementById('post-btn');

const sendHttpRequest = (method, url, data)=>{
  const promise = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open(method, url);

    xhr.responseType = 'json'
    if (data){
      xhr.setRequestHeader('Content-Type', 'application.json')
    }
    xhr.onload = () => {
      if(xhr.status > 400){
        reject('nonnon')
      }
      const data = xhr.response;
      resolve(data)
    }
    xhr.onerror = () => {
      reject('wrong')
    }
    xhr.send(JSON.stringify(data));
  })
  return promise;
}

getData = async () => {
  const data = await sendHttpRequest('GET', 'https://reqres.in/api/users')
  console.log({data})
}

postData = async () => {
  try{
  const data = await sendHttpRequest('POST', 'https://reqres.in/ap/users', {email: 'test@test.com', password: 'testtest'})
  console.log({data})
  }catch(ex){
    console.error(ex)
  }
}

getBtn.addEventListener('click', getData)
postBtn.addEventListener('click', postData)
```