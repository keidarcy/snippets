### How theme urls map

```
 - /thisisntarealurl → 404.liquid
 - /blogs/{blog-name}/{article-id-handle} → article.liquid
 - /blogs/{blog-name} → blog.liquid
 - /cart → cart.liquid
 - /collections → list-collections.liquid
 - /collections/{collection-handle} → collection.liquid
 - /collections/{collection-handle}/{tag} → collection.liquid
 - / → index.liquid
 - /pages/{page-handle} → page.liquid
 - /products → list-collections.liquid
 - /products/{product-handle} → product.liquid
 - /search?q={search-term} → search.liquid
```

### Resources list with node

#### nodejs with `shopify-api-node` package for `private` app with Admin Rest
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

#### nodejs with no package for `private` app with Admin Rest
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

#### broswer js for `private` app with Admin Rest

```js
fetch('https://xxx.myshopify.com/admin/api/2020-07/shop.json', {method:'GET',
headers: {'Authorization': 'Basic ' + btoa('apiKey:password')}})
.then(response => response.json())
.then(json => console.log(json));
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
    //user: "your database user id",
    //customer: "any custom data you want",
    //identifier: "bob123",
    //remote_ip: "107.00.000.000",
    //return_to: "http://yourstore.com/some_specific_site",
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

### Pass liquid data to Vue instance

```
<script id="data" type="application/json">{{ product | json }}</script>
```
```
const data = JSON.parse(document.getElementById('data').innerHTML);
new Vue({
  extends: MyComponent,
  propsData: data,
})
```

### [process](https://docs.google.com/spreadsheets/d/1gqvWtquj9W-bKcxDwMojYvIszXoVKonjOx5wAHZyMtk/edit#gid=1390331955)
