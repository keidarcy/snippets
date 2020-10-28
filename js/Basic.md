## axios

### Create global axios instance

```js
Axios.create({
  baseURL: BASE_URL,
  headers: {
    // get: {
    //   'Content-Type': 'application/json',
    //   'access-token': TOKEN
    // }
    'Content-Type': 'application/json',
    'access-token': TOKEN
  }
});
```

### Difference of `config` between get and post method

```js
const data = {
  sort: 'desc',
  count: 12
};

axios.post(URL, data);
axios.get(URL, { data });
```
