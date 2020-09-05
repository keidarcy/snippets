### Resources list with node

#### nodejs with `shopify-api-node` package for `private` app
```js
import Shopify from 'shopify-api-node';
import fetch from 'node-fetch';

const getProducts = async () => {
  const shopify = new Shopify({
    shopName: 'xxx.myshopify.com',
    apiKey: 'aaa',
    password: 'sss'
  });

  const orders = await shopify.product.list({ limit: 5 });
  console.log(orders[0]);
};

getProducts();
```

#### nodejs with no package for `private` app
```js
import fetch from 'node-fetch'

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

#### broswer js for `private` app

```js
fetch('https://xxx.myshopify.com/admin/api/2020-07/shop.json', {method:'GET',
headers: {'Authorization': 'Basic ' + btoa('apiKey:password')}})
.then(response => response.json())
.then(json => console.log(json));
```