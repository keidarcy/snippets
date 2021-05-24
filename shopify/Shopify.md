# Useful shopify notes

- [Useful shopify notes](#useful-shopify-notes)
  - [External Links](#external-links)
  - [Private app auth for admin api and storefront api](#private-app-auth-for-admin-api-and-storefront-api)
  - [How theme urls map](#how-theme-urls-map)
    - [url - template](#url---template)
    - [hidden url](#hidden-url)
  - [Checkout process](#checkout-process)
  - [Multipass](#multipass)
  - [Multiple currencies](#multiple-currencies)
  - [Pass liquid data to Vue instance](#pass-liquid-data-to-vue-instance)
  - [Theme editor](#theme-editor)
  - [Cart attribute](#cart-attribute)
  - [Additional scripts order status page](#additional-scripts-order-status-page)
  - [Order status](#order-status)
    - [order status](#order-status-1)
    - [Payment status](#payment-status)
    - [Checkout](#checkout)
    - [Collections urls](#collections-urls)

## External Links

- [live chat](https://help.shopify.com/en/questions#/contact)
- [Liquid Objects](https://shopify.dev/docs/themes/liquid/reference/objects)
- [schema types](https://shopify.dev/docs/themes/settings)
- [slatest](https://github.com/entozoon/slatest)(hot reload with compiler version themekit)
- [webhook](https://shopify.dev/tutorials/manage-webhooks)
- [image lazyload](https://github.com/aFarkas/lazysizes)

## Private app auth for admin api and storefront api

- Admin API
  - graphql endpoint `/admin/api/2020-10/graphql.json`
  - access token -> `Admin API` -> `password`

```
curl --request POST \
  --url https://STORE.myshopify.com/admin/api/2020-10/graphql.json \
  --header 'content-type: application/json' \
  --header 'x-shopify-access-token: TOKEN' \
  --data '{"query":"query {\n  shop{\n    id\n    primaryDomain{\n      host\n      sslEnabled\n      url\n    }\n    description\n    paymentSettings{\n       supportedDigitalWallets\n    }\n  }\n}"}'
```

- Storefront API
  - graphql endponint `/api/2020-07/graphql.json`
  - access token -> `Storefront API` -> `Storefront access token`

```
curl --request POST \
  --url https://STORE.myshopify.com/api/2020-07/graphql.json \
  --header 'accept: application/json' \
  --header 'content-type: application/json' \
  --header 'x-shopify-storefront-access-token: TOKEN' \
  --data '{"query":"query {\n  shop{\n    primaryDomain{\n      host\n      sslEnabled\n      url\n    }\n    description\n    paymentSettings{\n      countryCode\n      acceptedCardBrands\n      enabledPresentmentCurrencies\n    }\n    moneyFormat\n  }\n}"}'
```

## How theme urls map

### url - template

```
{/{ request.page_type }/} == {/{ template }/} which will be 404 | blog | cart ...
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
 - /account/login → customers/login.liquid
 - /account → customers/account
 - /account/addresses  → customers/addresses.liquid
 - /account/register →  customers/register.liquid
```

### hidden url

```
# recommendation api endpoint per product
https://STORE.myshopify.com/recommendations/products.json?product_id=ID&limit=4
```

## Checkout process

- "checkout_token": TOKEN

```
https://STORE.myshopify.com/NUMBER/checkouts/TOKEN
https://STORE.myshopify.com/NUMBER/checkouts/TOKEN?previous_step=contact_information&step=shipping_method
https://STORE.myshopify.com/NUMBER/checkouts/TOKEN?previous_step=shipping_method&step=payment_method
https://STORE.myshopify.com/NUMBER/checkouts/TOKEN/thank_you
# order_status_url
https://STORE.myshopify.com/NUMBER/orders/token
```

## Multipass

- [multipass-js](https://github.com/softmarshmallow/multipass-js)
- [multipass-doc](https://shopify.dev/docs/admin-api/rest/reference/plus/multipass)

## Multiple currencies

> money filter depends on admin setting `/admin/settings/general`.

- [liquid money filter](https://shopify.dev/docs/themes/liquid/reference/filters/money-filters)
- [liquid tutorials](https://shopify.dev/tutorials/customize-theme-support-multiple-currencies)
- [storefront api](https://shopify.dev/tutorials/support-multiple-currencies-with-storefront-api)

## Pass liquid data to Vue instance

```
<script id="data" type="application/json">{/{ product | json }/}</script>
```

```
const data = JSON.parse(document.getElementById('data').innerHTML);
new Vue({
  extends: MyComponent,
  propsData: data,
})
```

## Theme editor

- the name of `/admin/themes/id/editor?picker=section` is `presets.name.en`

```

```

## Cart attribute

```html
<p class="cart-attribute__field">
  <label for="name">name</label>
  <input id="name" type="text" name="attributes[name]" value="{/{ cart.attributes["name"]
  }/}">
</p>
```

- `name` attribute will be add to shopify order

```json
{
  "order": {
    "id": ID,
    "note_attributes": [
      {
        "name": "name",
        "value": "value"
      }
    ]
  }
}
```

## Additional scripts order status page

- [customize-order-status](https://help.shopify.com/en/manual/orders/status-tracking/customize-order-status)

```
{/% if first_time_accessed %/}
  // Conversion scripts you want to run only once
{/% endif %/}
```

## Order status

[link](https://help.shopify.com/en/manual/orders/order-status)

### order status

- Open
- Archived
- Canceled

### Payment status

- Authorized
- Paid
- Partially refunded
- Partially paid
- Pending
- Refunded
- Unpaid
- Voided

###

- Fulfilled
- Unfulfilled
- Partially fulfilled
- Scheduled

### Checkout

- outside of shopify use storefront API to checkout

```graphql
mutation checkoutCreate($input: CheckoutCreateInput!) {
  checkoutCreate(input: $input) {
    checkout {
      id
      webUrl
    }
    checkoutUserErrors {
      code
      field
      message
    }
  }
}
```

`webUrl` is the checkout url

- shopify to checkout

```html
<form action="/cart/add" method="post">
  <input type="text" name="id" value="{product.variants[0].id}" />
  <input id="_key" type="text" name="properties[_key]" />
  <input type="number" name="quantity" value="1" min="1" />
  <button type="submit">カートにテスト追加</button>
</form>
```

### Collections urls

[sample](https://turbo-theme.myshopify.com/)

- /collections/all?sort_by=manual
- /collections/all?sort_by=best-selling
- /collections/all?sort_by=title-ascending
- /collections/all?sort_by=title-descending
- /collections/all?sort_by=price-ascending
- /collections/all?sort_by=price-descending
- /collections/all?sort_by=created-ascending
- /collections/all?sort_by=created-descending
- /collections/all/dress+purple+25-50?sort_by=
