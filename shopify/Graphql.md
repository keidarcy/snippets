## Basic

- [QueryRoot](https://shopify.dev/docs/admin-api/graphql/reference/common-objects/queryroot/index)
- [Search Syntax](https://shopify.dev/concepts/about-apis/search-syntax)

### Operation Names and Variables

```graphql
query ProductTitleAndDescription($id: ID!) {
  product(id: $id){
    title
    description
  }
}

{
  "id": "gid://shopify/Product/xxx"
}

# Operation Names => ProductTitleAndDescription
# Variables => $id: ID!
```

#### graphql-request + graphql.macro

```graphql
query getProduct($id: ID!) {
  product(id: $id) {
    id
    title
  }
}
```

```js
require('dotenv').config();
const { loader } = require('graphql.macro');
const { GraphQLClient, gql } = require('graphql-request');

async function main() {
  const endpoint = process.env.SHOPIFY_URL;

  const graphqlClient = new GraphQLClient(endpoint, {
    headers: {
      'Content-type': 'application/json',
      'X-Shopify-Access-Token': process.env.SHOPIFY_API_PASSWORD
    }
  });

  // const query = gql`
  //   query getProduct($id: ID!) {
  //     product(id: $id) {
  //       id
  //       title
  //     }
  //   }
  // `;

  const variables = {
    id: 'gid://shopify/Product/XXX'
  };
  const query = loader('../graphqls/getOneProduct.gql');

  const data = await graphqlClient.request(query, variables);
  console.log(JSON.stringify(data, undefined, 2));
}

main().catch((error) => console.error(error));
```

### Aliases

- usage

```graphql
query ProductTitleAndDescription($id: ID!) {
  product(id: $id) {
    myRenamedAliasesTitle: title
    description
  }
}
```

- usage

```graphql
query  {
  product1:product(id: 'gid://shopify/Product/XXX') {
    title
    description
  }
  product2:product(id: 'gid://shopify/Product/XXX') {
    title
    description
  }
}
```

### Fragments

```graphql
query  {
  product1:product(id: 'gid://shopify/Product/XXX') {
    ...TitleAndDescription
  }
  product2:product(id: 'gid://shopify/Product/XXX') {
    ...TitleAndDescription
  }
}
fragment TitleAndDescription on Product {
  title
  description
  featuredImage {
    src
  }
}
```

- inline fragment

```graphql
mutation tagsAdd($id: ID!, $tags: [String!]!) {
  tagsAdd(id: $id, tags: $tags) {
    node {
      id
    }
    userErrors {
      field
      message
    }
  }
}

{
  "id": "Z2lkOi8vU2hvcGlmeS9FeGFtcGxlLzE=",
  "tags": [
    "placeholder"
  ]
}
```

_Customer_ and _Product_ object implments _node_ interface, so use inline fragment to query field in _Customer_ and _Product_ value.

```graphql
mutation tagsAdd($id: ID!, $tags: [String!]!) {
  tagsAdd(id: $id, tags: $tags) {
    node {
      id
      ... on Product {
        title
        tags
      }
      ... on Customer {
        email
      }
    }
    userErrors {
      field
      message
    }
  }
}
```

### Pagination

```graphql
query threeProducts {
  products(first: 3) {
    edges {
      cursor
      node {
        id
        title
      }
    }
    pageInfo {
      hasNextPage
      hasPreviousPage
    }
  }
}

query threeProducts {
  products(first: 3, after: "xxxx") {
    edges {
      cursor
      node {
        id
        title
      }
    }
    pageInfo {
      hasNextPage
      hasPreviousPage
    }
  }
}
```

### Query argument

- find products that tagged with 'a' AND not tagged with 'b'

```graphql
query taggedProducts {
  products(first: 3, query: '-tag: a AND tag: b') {
    edges {
      node {
        title
        description
        tags
      }
    }
  }
}

```

## Admin API

#### Get product list with requirement

```graphql
{
  products(first: 30, query: "", after: "{{ the cursor you want to use }}") {
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
query VariantWithMeta($id: ID!) {
  productVariant(id: $id) {
    metafields(first: 10) {
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

```graphql
query {
  shop {
    primaryDomain {
      host
      sslEnabled
      url
    }
    description
    paymentSettings {
      countryCode
      acceptedCardBrands
      enabledPresentmentCurrencies
    }
    moneyFormat
  }
}
```

#### Get price with different currencies(which are setted in admin page)

```graphql
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
