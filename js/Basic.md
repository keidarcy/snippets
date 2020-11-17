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

## Get file name without extension

```js
arr.split('.').slice(0, -1).join('.');
```

## Date

```js
Date.now();
//1605615577351
new Date()(
  //Tue Nov 17 2020 21:19:42 GMT+0900 (Japan Standard Time)
  new Date()
).toString();
//"Tue Nov 17 2020 21:20:30 GMT+0900 (Japan Standard Time)"
Date.parse('Tue Nov 17 2020 21:20:30 GMT+0900 (Japan Standard Time)');
//1605615630000
new Date().getFullYear();
//2020
// GMT is an abbreviation for Greenwich Mean Time
```
