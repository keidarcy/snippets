# Useful shopify notes

## External Links

- [slatest](https://github.com/entozoon/slatest)(hot reload with compiler version themekit)
- [Liquid Objects](https://shopify.dev/docs/themes/liquid/reference/objects)

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
{{ request.page_type }} == {{ template }} which will be 404 | blog | cart ...
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
<script id="data" type="application/json">{{ product | json }}</script>
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
  <input id="name" type="text" name="attributes[name]" value="{{ cart.attributes["name"]
  }}">
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
{% if first_time_accessed %}
  // Conversion scripts you want to run only once
{% endif %}
```

## Product csv sample

```
Handle,Title,Body (HTML),Vendor,Type,Tags,Published,Option1 Name,Option1 Value,Option2 Name,Option2 Value,Option3 Name,Option3 Value,Variant SKU,Variant Grams,Variant Inventory Tracker,Variant Inventory Qty,Variant Inventory Policy,Variant Fulfillment Service,Variant Price,Variant Compare At Price,Variant Requires Shipping,Variant Taxable,Variant Barcode,Image Src,Image Position,Image Alt Text,Gift Card,SEO Title,SEO Description,Google Shopping / Google Product Category,Google Shopping / Gender,Google Shopping / Age Group,Google Shopping / MPN,Google Shopping / AdWords Grouping,Google Shopping / AdWords Labels,Google Shopping / Condition,Google Shopping / Custom Product,Google Shopping / Custom Label 0,Google Shopping / Custom Label 1,Google Shopping / Custom Label 2,Google Shopping / Custom Label 3,Google Shopping / Custom Label 4,Variant Image,Variant Weight Unit,Variant Tax Code,Cost per item
example-t-shirt,Example T-Shirt,,Acme,Shirts,mens t-shirt example,TRUE,Title,"Lithograph - Height: 9"" x Width: 12""",,,,,,3629,,,deny,manual,25,,TRUE,TRUE,,https://help.shopify.com/images/green-t-shirt.jpg,1,,FALSE,Our awesome T-shirt in 70 characters or less.,A great description of your products in 320 characters or less,Apparel & Accessories > Clothing,Unisex,Adult,7X8ABC910,T-shirts,"cotton, pre-shrunk",used,FALSE,,,,,,,g,,
example-t-shirt,,,,,,,,Small,,,,,example-shirt-s,200,,,deny,manual,19.99,24.99,TRUE,TRUE,,,,,,,,,,,,,,,,,,,,,,g,,
example-t-shirt,,,,,,,,Medium,,,,,example-shirt-m,200,shopify,,deny,manual,19.99,24.99,TRUE,TRUE,,,,,,,,,,,,,,,,,,,,,,g,,

```
