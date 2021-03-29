- [TCP/IP](#tcpip)
- [WebRTC](#webrtc)
- [WebSocket & WebTransport](#websocket--webtransport)
  - [server node demo](#server-node-demo)
  - [client browser websocket handshake](#client-browser-websocket-handshake)

## TCP/IP

![handshake](tcp-ip-handshake.png)

![tcp-tls](https://www.cloudflare.com/resources/images/slt3lc6tev37/5aYOr5erfyNBq20X5djTco/3c859532c91f25d961b2884bf521c1eb/tls-ssl-handshake.png)
![High Performance Browser Networking](https://hpbn.co/)

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

- Pros
  - Full-duplex(no polling)
  - HTTP compatible
  - Filewall friendly(standard)
- Cons
  - Proxying is tricky
  - L7 L/B challengind (timeouts)
  - Stateful, difficult to horizontally scale

### server node demo

```js
const http = require('http');
const WebSocketServer = require('websocket').server;

let connection = null;

const httpserver = http.createServer((req, res) => {
  console.log('We have received a request');
});

httpserver.listen(8080, () => console.log('my server runing on 8080 port'));

const websocket = new WebSocketServer({
  httpServer: httpserver
});
websocket.on('request', (request) => {
  connection = request.accept(null, request.origin);
  connection.on('open', () => console.log('Opened'));
  connection.on('close', () => console.log('closed'));
  connection.on('message', (message) => {
    console.log(`Received message ${message.utf8Data}`);
  });
  sendPerFiveSeconds();
});

const sendPerFiveSeconds = () => {
  connection.send(`message ${Math.random()}`);
  setTimeout(sendPerFiveSeconds, 5000);
};
```

### client browser websocket handshake

```js
let ws = new WebSocket('ws://localhost:8080');
ws.onmessage = (message) => console.log(`message from server ${message.data}`);
ws.send('hhello server this is client');
ws.close();
```
