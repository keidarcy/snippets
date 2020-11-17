# Nodejs basic

## Cli parameters

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

## Generate zip file from string

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

## Node cli clean screen

- `Ctrl + L` clean screen, similar to `process.stdout.write("\u001b[2J\u001b[0;0H");`

## Get file name without extension

```js
arr.split('.').slice(0, -1).join('.');
```
