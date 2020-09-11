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

### multipass

```
yarn add multipass-js
```

```ts
import { Multipass } from "multipass-js"

const SHOPIFY_STORE_MULTIPASS_SECRET = 'xxx'; // GET from admin page setting => payment => enable Multipass loginultip
const multipass = new Multipass(SHOPIFY_STORE_MULTIPASS_SECRET);

// Create your customer data hash
const email = `bob@bob.com`
const customerData = {
    email: email,
    user: "your database user id",
    customer: "any custom data you want"
    // ...
};

const url = multipass
  .withCustomerData(customerData)
  .withDomain('xxxxx.myshopify.com/')
  .withRedirect('/products/adidas-smith')
  .url();

console.log(url);
// https://xxx.myshopify.com//account/login/multipass/[LONG_LONG_STRING]


// client may access shopify with `url`
// will give you URL like:  https://store.myshopify.com/account/login/multipass/<MULTIPASS-TOKEN>
// with optional redirection
```
