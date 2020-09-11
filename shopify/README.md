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

 - [multipass-js](https://github.com/softmarshmallow/multipass-js)
 - [multipass-doc](https://shopify.dev/docs/admin-api/rest/reference/plus/multipass)

```
yarn add multipass-js
```

```ts
import { Multipass } from "multipass-js"

const SHOPIFY_STORE_MULTIPASS_SECRET = 'xxx'; // GET from admin page setting => checkout => enable Multipass loginultip
const multipass = new Multipass(SHOPIFY_STORE_MULTIPASS_SECRET);

// Create your customer data hash
const customerData = {
    email: email,
    user: "your database user id",
    customer: "any custom data you want"
    // ...
    email: 'bob@bob.com',
    created_at: '2013-04-11T15:16:23-04:00',
    first_name: 'Bob',
    last_name: 'Bobsen',
    tag_string: 'canadian, premium',
    addresses: [
      {
        address1: '123 Oak St',
        city: 'Ottawa',
        country: 'Canada',
        first_name: 'Bob',
        last_name: 'Bobsen',
        phone: '555-1212',
        province: 'Ontario',
        zip: '123 ABC',
        province_code: 'ON',
        country_code: 'CA',
        default: true
      }
    ]
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
