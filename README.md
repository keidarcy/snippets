# snippets

#### Generate zip file from string

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

#### Complile ts, zip, update lambda function

```ts
const ncc = require('@vercel/ncc');
import jszip from 'jszip';
import dotenv from 'dotenv';
import AWS from 'aws-sdk';
import ora from 'ora';
import filesize from 'filesize';

dotenv.config();

const config = {
  aws: {
    accessKeyId: process.env.AWS_ACCESS_KEY_ID,
    secretAccessKey: process.env.AWS_SECRET_ACCESS_KEY,
    region: process.env.AWS_DEFAULT_REGION
  }
};

const compile = async (): Promise<string> => {
  const res = await ncc(__dirname + '/../index.ts', { minify: true });
  return res.code;
};

const createZip = async (code: string): Promise<Buffer> => {
  const zip = new jszip();
  zip.file('index.js', code);
  return zip.generateAsync({ type: 'nodebuffer' });
};

const uploadToAws = async (zipFile: Buffer): Promise<number> => {
  const lambda = new AWS.Lambda(config.aws);
  const res = await lambda
    .updateFunctionCode({
      ZipFile: zipFile,
      FunctionName: 'ts-lambda-test'
    })
    .promise();
  return res.CodeSize as number;
};

const deploy = async () => {
  try {
    const spinner = ora();
    const code = await compile();
    spinner.start('Creating an archive');
    const zipFile = await createZip(code);
    spinner.succeed('Archive ready. size:' + filesize(code.toString().length));
    spinner.start('Uploading to aws');
    const res = await uploadToAws(zipFile);
    spinner.succeed('Upload done. size:' + filesize(res));
  } catch (err) {
    console.error(err);
  }
};

deploy();
```

#### TS decorator

```ts
function time(name: string) {
    return function (target, propertyKey: string, descriptor: PropertyDescriptor) {
        const fn = descriptor.value;
        descriptor.value = (...args) => {
            console.time(name);
            const v = fn(...args);
            console.timeEnd(name);
            return v;
        }
    }
}

class C{
    @time('C.method')
    method(name: string){
        console.log('method called', name)
        for (let i = 0; i < 1000000000; i++){}
    }
}

new C().method('koko')
```
#### Node cli parameter 

```
#! /usr/bin/env node

const args = process.argv.slice(2);
const mydir = args[0];
const myname = args[1].split('--name=')[1];
console.log({ mydir, myname });
```
input `./dist/index.js myqpp --name=fun`
output `{ mydir: 'myqpp', myname: 'fun' }`

