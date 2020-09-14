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
