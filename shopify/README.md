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

## Multiple currencies
> money filter depends on admin setting `/admin/settings/general`.
 - [liquid money filter](https://shopify.dev/docs/themes/liquid/reference/filters/money-filters)
 - [liquid tutorials](https://shopify.dev/tutorials/customize-theme-support-multiple-currencies)
 - [storefront api](https://shopify.dev/tutorials/support-multiple-currencies-with-storefront-api)

## How theme urls map

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

### multipass

 - [multipass-js](https://github.com/softmarshmallow/multipass-js)
 - [multipass-doc](https://shopify.dev/docs/admin-api/rest/reference/plus/multipass)



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
