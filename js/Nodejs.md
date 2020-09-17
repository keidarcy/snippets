
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
