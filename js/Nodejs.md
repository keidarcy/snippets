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

## Native node to upload to aws lambda

```js
module.exports = {
  moduleName
};
```

```json
{
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1",
    "zip": "zip -r -j xxx.zip ./dist/index.js",
    "upload": "aws lambda update-function-code --function-name xxx --zip-file fileb://xxx.zip",
    "build": "ncc build -o ./dist -m",
    "deploy": "yarn build && yarn zip && yarn upload"
  }
}
```

### Create and update json file

```ts
import fs from 'fs';

const getMapping = async (id: string, myId: number) => {
  const myPath = __dirname + '/my.json';
  try {
    await fs.promises.access(myPath);
    const existedMapping = await fs.promises.readFile(myPath, 'utf-8');
    const mappingJson = JSON.parse(existedMapping);
    mappingJson[id] = myId;
    await fs.promises.writeFile(myPath, JSON.stringify(mappingJson), 'utf-8');
  } catch (error) {
    console.log(error);
    await fs.promises.writeFile(myPath, JSON.stringify({}), 'utf-8');
  }
};
getMapping('9', 789);
```

## npm peer dependency

- `peerDependency` changed in `npm7`

add `--legacy-peer-deps` to excute like `npm6`

example `npx --legacy-peer-deps sb init`
