const express = require('express');
const { MongoClient } = require('mongodb');
const url = process.env.MONGO_CONNECTION_STRING || 'mongodb://localhost:27017';
const dbName = 'dockerApp';
const collectionName = 'count';

const app = express();
const port = 3000;

async function start() {
  const client = await MongoClient.connect(url);
  const db = client.db(dbName);
  const collection = db.collection(collectionName);

  app.get('/', (req, res) => res.send('Hello World!'));

  app.get('/get', async (req, res) => {
    const count = await collection.count();
    console.log('count', count);
    res.json({ success: true, count });
  });

  app.get('/add', async (req, res) => {
    console.log('update');
    const response = await collection.insertOne({});
    res.json({ inserted: response.insertedCount });
  });
}
start();

app.listen(port, () => console.log(`Example app listening on port ${port}!`));
