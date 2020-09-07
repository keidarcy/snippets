### AWS S3 + CloudWatch + Lambda + APIGatewat with ts

#### Complile ts, zip, update lambda function => scripts/deply.ts

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

#### Implement S3 resources and send metrics to cloudwatch

```ts
import {
  Context,
  APIGatewayProxyHandler,
  APIGatewayEvent,
  APIGatewayProxyResult
} from 'aws-lambda';
import AWS from 'aws-sdk';

const s3 = new AWS.S3();
const cloudWatch = new AWS.CloudWatch();

const BUCKET_NAME = 'image-uploader-s3-app';

const getFiles = async () => {
  const files: Array<{ name: string }> = [];
  const data = await s3
    .listObjectsV2({
      Bucket: BUCKET_NAME
    })
    .promise();
  if (data && data.Contents) {
    for (const item of data.Contents) {
      files.push({ name: item.Key as string });
    }
  }
  return { files };
};

const deleteFiles = async (name: string) => {
  const res = await s3
    .deleteObject({
      Key: name,
      Bucket: BUCKET_NAME
    })
    .promise();
  return res;
};

const updateFiles = async (name: string, contentType: string, content: string) => {
  await s3
    .upload({
      Key: name,
      Bucket: BUCKET_NAME,
      ContentType: contentType,
      Body: content
    })
    .promise();
  return {};
};

export const handler: APIGatewayProxyHandler = async (
  event: APIGatewayEvent,
  context: Context
): Promise<APIGatewayProxyResult> => {
  const metrics: AWS.CloudWatch.MetricData = [];
  const res: APIGatewayProxyResult = {
    headers: {
      'Access-Control-Allow-Origin': '*',
      'Content-Type': 'application/json'
    },
    statusCode: 200,
    body: ''
  };
  let body = {};

  try {
    const method = event.httpMethod;
    const path = event.path;
    const req = JSON.parse(event.body as string);
    if (path === '/files') {
      if (method === 'GET') {
        // GET /files => {files: [{name:string}]}
        body = await getFiles();
        metrics.push({ MetricName: 'get-files', Value: 1, Timestamp: new Date() });
      } else if (method === 'POST') {
        // POST /files {name: string, contentType: string, content: string } => {}
        body = await updateFiles(req.name, req.contentType, req.content);
        metrics.push({ MetricName: 'post-files', Value: 1, Timestamp: new Date() });
      } else if (method === 'DELETE') {
        // DELETE /files {name: string} => {}
        body = await deleteFiles(req.name);
        metrics.push({ MetricName: 'delete-files', Value: 1, Timestamp: new Date() });
      } else if (method === 'PUT') {
        body = { method: 'THIs is put' };
        metrics.push({ MetricName: 'put-files', Value: 1, Timestamp: new Date() });
      }
    }
    try {
      await cloudWatch
        .putMetricData({
          Namespace: BUCKET_NAME,
          MetricData: metrics
        })
        .promise();
    } catch (error) {}
    res.body = JSON.stringify(body);
  } catch (err) {
    console.log(err);
    return err;
  }
  return res;
};
```

#### Dependency, Execution
```bash
yarn add -D @types/aws-lambda @types/aws-sdk @types/dotenv @types/filesize @types/node @types/ora @vercle/ncc aws-sdk dotenv filesize jszip ora ts-node typescript
```

```bash
node --max-old-space-size=4096 -- node_modules/.bin/ts-node -P tsconfig.json scripts/deploy.ts
```
```json
{
  "scripts": {
    "deploy": "ts-node scripts/deploy.ts",
  },
  "devDependencies": {
    "@types/aws-lambda": "^8.10.61",
    "@types/aws-sdk": "^2.7.0",
    "@types/dotenv": "^8.2.0",
    "@types/express": "^4.17.8",
    "@types/filesize": "^5.0.0",
    "@types/node": "^14.6.4",
    "@types/ora": "^3.2.0",
    "@vercel/ncc": "^0.24.0",
    "aws-sdk": "^2.746.0",
    "dotenv": "^8.2.0",
    "express": "^4.17.1",
    "filesize": "^6.1.0",
    "jszip": "^3.5.0",
    "ora": "^5.0.0",
    "ts-node": "^9.0.0",
    "typescript": "^4.0.2"
  }
}

```