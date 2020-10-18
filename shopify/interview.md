[sandbox](https://codesandbox.io/s/festive-water-g9sbj?file=/src/index.js)

###### 1. What's the output?

```javascript
function sayHi() {
  console.log(name);
  console.log(age);
  var name = 'name';
  let age = 21;
}

sayHi();
```

- A: `name` and `undefined`
- B: `name` and `ReferenceError`
- C: `ReferenceError` and `21`
- D: `undefined` and `ReferenceError`

<details><summary><b>Answer</b></summary>
<p>

#### Answer: D

</p>
</details>

---

###### 2. How to create range function in js?

```python
# python 3.x
list(range(5))
# [0, 1, 2, 3, 4]
```

```javascript
range(5);
```

<details><summary><b>Answer</b></summary>
<p>

#### Answer:

```js
const range = (num) => [...Array(num).keys()];
```

</p>
</details>

---

###### 3. How to get a value from cookie?

```javascript
document.cookie;
// "_ga=GA1.2.1676417641.1593312683; tz=Asia%2FTokyo; _octo=GH1.1.945928125.1603004802"
// get tz value
```

<details><summary><b>Answer</b></summary>
<p>

#### Answer:

```js
let cookies = document.cookie
  .split(';')
  .map((cookie) => cookie.split('='))
  .reduce(
    (accumulator, [key, value]) => ({
      ...accumulator,
      [key.trim()]: decodeURIComponent(value)
    }),
    {}
  );
```

</p>
</details>

---

###### 4. What's the output?

```javascript
const promise1 = Promise.resolve('First')
const promise2 = Promise.resolve('Second')
const promise3 = Promise.reject('Third')
const promise4 = Promise.resolve('Fourth')

const runPromises = async () => {
	const res1 = await Promise.all([promise1, promise2])
	const res2  = await Promise.all([promise3, promise4])
	return [res1, res2]
}

runPromises()
	.then(res => console.log(res))
	.catch(err => console.log(err))o
```

- A: `[['First', 'Second'], ['Fourth']]`
- B: `[['First', 'Second'], ['Third', 'Fourth']]`
- C: `[['First', 'Second']]`
- D: `Third`

<details><summary><b>Answer</b></summary>
<p>

#### Answer: D

</p>
</details>

---

---

###### 6. GraphQL to fetch resouces

```graphql
query GetProductList {
  products(first: 10) {
    edges {
      node {
        id
        title
        tags
        featuredImage {
          originalSrc
        }
      }
    }
  }
}
```

- get `test` tagged product
- change sku

<details><summary><b>Answer</b></summary>
<p>

#### Answer:

```js
query GetProductList {
  products(first: 10, query: "tag:test") {
    edges {
      node {
        id
        title
        tags
      }
    }
  }
}


mutation skuUpdate($input: ProductVariantInput!) {
  productVariantUpdate(input: $input) {
    productVariant{
      id
      sku
    }
    userErrors {
      field
      message
    }
  }
}

{
	"input": {
		"id": "gid://shopify/ProductVariant/id",
		"sku": "test-sku"
	}
}
```

</p>
</details>

---
