# Jest

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

### test async data

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

### Setup and Teardown

```js
beforeEach(() => initDatabase());
afterEach(() => closeDatabase());

beforeAll(() => initDatabase());
afterAll(() => closeDatabase());

const initDatabase = () => console.log('Database Initialized...');
const closeDatabase = () => console.log('Database Closed...');
```

## React Testing Library

```tsx
import { useState } from 'react';
import axios from 'axios';

const Counter = () => {
  const [counter, setCounter] = useState(0);

  const fetchRandomNumber = () => {
    const fn = async () => {
      const { data } = await axios('https://jsonplaceholder.typicode.com/users/1');
      setCounter(data.address.geo.lng.replace(/[^\d]/g, ''));
    };
    fn();
  };

  return (
    <>
      <button data-testid="minus" onClick={() => setCounter(counter - 1)}>
        -
      </button>
      <input
        type="number"
        role="number-input"
        placeholder="random-number"
        value={counter}
        onChange={(e: React.ChangeEvent<HTMLInputElement>) => setCounter(+e.target.value)}
      />
      <button data-testid="plus" onClick={() => setCounter(counter + 1)}>
        +
      </button>

      <button onClick={fetchRandomNumber}>get random number</button>
    </>
  );
};

export default Counter;
```

## test change, click event and async data

```tsx
import Counter from '../pages/counter';
import { render, cleanup, fireEvent, waitFor } from '@testing-library/react';

describe('counter', () => {
  afterEach(cleanup);

  test('Counter test', () => {
    // const handleChange = jest.fn();
    const { getByText, container } = render(<Counter />);
    const minusButton = getByText('-');
    const plusButton = getByText('+');
    const input = container.getElementsByTagName('input')[0];

    fireEvent.change(input, { target: { value: 20 } });
    expect(+input.value).toBe(20);

    fireEvent.click(minusButton);
    expect(+input.value).toBe(19);

    fireEvent.click(plusButton);
    expect(+input.value).toBe(20);
  });

  test('counter with async', async () => {
    const { getByText, container } = render(<Counter />);
    const input = container.getElementsByTagName('input')[0];

    fireEvent.click(getByText('get random number'));
    await waitFor(() => expect(+input.value).toBeGreaterThan(200));
  });
});
```

### setup React testing Library in Nextjs

- `package.json`

```json
  "devDependencies": {
    "@testing-library/dom": "^7.29.4",
    "@testing-library/jest-dom": "^5.11.9",
    "@testing-library/react": "^11.2.3",
    "@types/react": "^17.0.0",
    "babel-jest": "^26.6.3",
    "jest": "^26.6.3",
    "typescript": "^4.1.3"
  },
  "jest": {
    "testPathIgnorePatterns": [
      "<rootDir>/.next/",
      "<rootDir>/node_modules/"
    ],
    "setupFilesAfterEnv": [
      "<rootDir>/setupTests.js"
    ],
    "transform": {
      "^.+\\.(js|jsx|ts|tsx)$": "<rootDir>/node_modules/babel-jest"
    },
    "moduleNameMapper": {
      "\\.(jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga)$": "<rootDir>/__mocks__/fileMock.js",
      "\\.(css|less)$": "<rootDir>/__mocks__/styleMock.js"
    }
  }
```

- `.babelrc`

```json
{
  "presets": ["next/babel"]
}
```

- `__mocks__/fileMock.js`

```js
module.exports = 'placeholder-file';
```

- `__mocks__/styleMock.js`

```js
module.exports = {};
```

- `./setupTests.js`

```js
import '@testing-library/jest-dom';
```
