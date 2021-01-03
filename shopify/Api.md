### Resources list with node

#### nodejs with `shopify-api-node` package for `private` app with Admin Rest

- create new smart collection
- [smart collection all rules and columns](https://shopify.dev/docs/admin-api/rest/reference/products/smartcollection)

```js
const Shopify = require('shopify-api-node');
const dotenv = require('dotenv');

dotenv.config();

const shopify = new Shopify({
  shopName: process.env.shopName,
  apiKey: process.env.apiKey,
  password: process.env.password
});

const createSmartCollection = async (title, column, relation) => {
  const rules = [
    {
      column,
      relation,
      condition: title
    }
  ];
  try {
    await shopify.smartCollection.create({ title, body_html: '', rules });
    console.log(`---${title}---SUCCESS------`);
  } catch (error) {
    console.log('------ERROR------');
    console.log(error.message);
  }
};
```

- get products

```js
import Shopify from 'shopify-api-node';
import fetch from 'node-fetch';

const getProducts = async () => {
  const shopify = new Shopify({
    shopName: 'xxx.myshopify.com',
    apiKey: 'aaa',
    password: 'sss'
  });

  const products = await shopify.product.list({ limit: 5 });
  console.log(products[0]);
};

getProducts();
```

#### nodejs with no package for `private` app with Admin Rest

```js
import fetch from 'node-fetch';

const url = 'https://xx.myshopify.com/admin/api/2020-07/shop.json';
const username = 'aaa';
const password = 'bbb';

fetch(url, {
  method: 'GET',
  headers: {
    Authorization: 'Basic ' + Buffer.from(username + ':' + password).toString('base64')
  }
})
  .then((response) => response.json())
  .then((json) => console.log(json));
```

#### broswer js for `private` app with Admin Rest

```js
fetch('https://xxx.myshopify.com/admin/api/2020-07/shop.json', {
  method: 'GET',
  headers: { Authorization: 'Basic ' + btoa('apiKey:password') }
})
  .then((response) => response.json())
  .then((json) => console.log(json));
```

#### Curl`private` app with Stronfront Api(Graphql)

```bash
curl --request POST \
  --url https://xxx.myshopify.com/api/2019-07/graphql \
  --header 'accept: application/json' \
  --header 'content-type: application/json' \
  --header 'x-shopify-storefront-access-token: xxxx' \
  --cookie __cfduid=xxxxxx \
  --data '{"query":"{\n  shop {\n    name\n    primaryDomain {\n      url\n      host\n        \n    }\n  }\n}"}'
# Drag this into insomnia!
```

#### Curl `prvate` app with Admin API(rest)

```bash
curl --request GET \
  --url https://xxxxxx.myshopify.com//admin/api/2020-10/smart_collections.json \
  --header 'X-Shopify-Access-Token: shppa_xxxx'

```
