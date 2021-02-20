require('dotenv').config();
const { GraphQLClient, gql } = require('graphql-request');

async function main() {
  const endpoint = process.env.SHOPIFY_URL;

  const graphqlClient = new GraphQLClient(endpoint, {
    headers: {
      'Content-type': 'application/json',
      'X-Shopify-Access-Token': process.env.SHOPIFY_API_PASSWORD
    }
  });

  const query = gql`
    query getProduct($id: ID!) {
      product(id: $id) {
        id
        title
      }
    }
  `;

  const variables = {
    id: 'gid://shopify/Product/XXX'
  };

  const data = await graphqlClient.request(query, variables);
  console.log(JSON.stringify(data, undefined, 2));
}

main().catch((error) => console.error(error));
