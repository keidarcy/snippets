## Google Spreadsheet

### Use custom function

- editor side

```js
/**
 * function description.
 *
 * @param param1 param1 description.
 * @param param2 param2 description.
 * @return return value description.
 * @customfunction
 */
function GET_SUM(value, range) {
  Logger.log(value, range);
  return value + range[0][0] + range[0][1];
}
```

- spreadsheet side

```
=GET_SUM(A3, B3:C3)
```