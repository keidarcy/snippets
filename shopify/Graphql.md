## Admin API
#### Get product list with requirement

```
{
  products(first: 30, query: "",  after:"{{ the cursor you want to use }}") {
    pageInfo {
      hasNextPage
      hasPreviousPage
    }
    edges {
      cursor
      node {
        id
        title
        images(first: 1) {
          edges {
            node {
              originalSrc
            }
          }
        }
      }
    }
  }
}

```


#### Get variant with metafields

```graphql
query VariantWithMeta($id: ID!){
      productVariant(id: $id) {
          metafields(first:10){
              edges {
                  node {
                      id
                      key
                      value
                  }
              }
          }
      }
  }

```

```graphql
{
  "id": "gid://shopify/ProductVariant/id"
}
```

#### Update product metafields

```graphql

mutation($input: ProductInput!) {
    productUpdate(input: $input) {
        product {
            id
        }
        userErrors {
            field
            message
        }
    }
}

{
    "id": "gid://shopify/Product/id",
    "metafields" => [
        {
            "id": "gid://shopify/Metafield/id",
            "key": "string",
            "value": "string",
            "valueType" "STRING",
            "namespace": "string"
        },
        {
            "id": "gid://shopify/Metafield/id",
            "key": "string",
            "value": "string",
            "valueType" "STRING",
            "namespace": "string"
        }
    ]
}
```

## Storefront Api

#### Shop information

```
query {
  shop{
    primaryDomain{
      host
      sslEnabled
      url
    }
    description
    paymentSettings{
      countryCode
      acceptedCardBrands
      enabledPresentmentCurrencies
    }
    moneyFormat
  }
}
```

#### Get price with different currencies(which are setted in admin page)
 
```
{
  productByHandle(handle: "adidas-classic-backpack") {
    id
    title
    variants(first: 1) {
      edges {
        node {
          presentmentPrices(first: 1, presentmentCurrencies: [USD, CNY]) {
            edges {
              node {
                compareAtPrice {
                  amount
                  currencyCode
                }
                price {
                  amount
                  currencyCode
                }
              }
            }
          }
          unitPrice {
            amount
          }
          requiresShipping
          availableForSale
          id
          title
          priceV2 {
            currencyCode
            amount
          }
        }
      }
    }
  }
}
```
