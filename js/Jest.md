## Basic

```js
describe('number test', () => {
  it('1 is true', () => {
    expect(1).toBeTruthy();
  });
  test('2 is true', () => {
    expect(2).toBeTruthy();
  });
});
```

## test async data

```js
test('async axios json', () => {
  expect.assertions(1);
  return functions.fetchUser().then((data) => {
    expect(data.name).toBe('Leanne Graham');
  });
});

test('async axios json', async () => {
  expect.assertions(1);
  const data = await functions.fetchUser();
  expect(data.name).toBe('Leanne Graham');
});
```
