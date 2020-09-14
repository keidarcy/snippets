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

```
{
  "id": "gid://shopify/ProductVariant/id"
}
```
