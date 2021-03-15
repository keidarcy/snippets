# Typescript basic

[handbook](https://microsoft.github.io/TypeScript-New-Handbook/everything)

## Decorator

```ts
function time(name: string) {
  return function (target, propertyKey: string, descriptor: PropertyDescriptor) {
    const fn = descriptor.value;
    descriptor.value = (...args) => {
      console.time(name);
      const v = fn(...args);
      console.timeEnd(name);
      return v;
    };
  };
}

class C {
  @time('C.method')
  method(name: string) {
    console.log('method called', name);
    for (let i = 0; i < 1000000000; i++) {}
  }
}

new C().method('koko');
```

## keyof

```ts
interface Person {
  name: string;
  age: number;
  location: string;
}

type K1 = keyof Person; // "name" | "age" | "location"
type K2 = keyof Person[]; // "length" | "push" | "pop" | "concat" | ...
type K3 = keyof { [x: string]: Person }; // string
```

## Mapped Types

```ts
interface Person {
  name: string;
  age: number;
  location: string;
}

type Partial<T> = {
  [P in keyof T]?: T[P];
};

type PartialPerson = Partial<Person>;
```

## For those packages that have no type definition

- putting an empty declaration for it in a ` .d.ts` file in your project like below.

```ts
declare module 'some-untyped-module';
```

## Add variable to globle object

```ts
declare global {
  interface Window {
    Hello: {
      moneyFormatter: (price: string) => string;
    };
  }
}
```

## never keyword

- example

```ts
const d = function neverExcute() {
  throw 'nothing';
};

// const d: () => never
```

- why it is useful

```ts
interface Square {
  kind: 'square';
}

interface Rectangle {
  kind: 'rectandle';
}

interface Circle {
  kind: 'circle';
}

type Shape = Square | Rectangle | Circle;

function something(s: Shape) {
  if (s.kind === 'square') {
  } else if (s.kind === 'circle') {
  } else if (s.kind === 'rectandle') {
  }
  const _bugTester: never = s;
}
```

## Confitional Type

```ts
type IsNumber<T> = T extends number ? 'number' : 'other';

type WithNumber = IsNumber<number>;
type WithOther = IsNumber<boolean[]>;
```

generic type is `number` => literal 'number' as type.
not `number` => literal 'other' as type.

```ts
type NumberOrString<T extends number | string> = T extends number ? number : string;

function createLabel<T extends number | string>(numberOrString: T): NumberOrString<T> {
  throw 'unimplemented';
}

const createSecondLabel = <T extends number | string>(
  numberOrString: T
): NumberOrString<T> => {
  throw 'unimplemented';
};

const a = createLabel('d');
const b = createLabel(8.9);
const c = createLabel(+Math.random().toFixed(2) * 2 > 1 ? 'STRING' : 8);
```

## `infer` operater

```ts
type IsArray<T> = T extends Array<infer O> ? O : T;

type ItIsArray = IsArray<string[]>; // string
type ItIsNotArray = IsArray<boolean>; // boolean
```

## Indexed Access Types

```ts
const MyArray = [
  { name: 'Alice', age: 15 },
  { name: 'Bob', age: 23 },
  { name: 'Eve', age: 38 }
] as const;

type Ages = typeof MyArray[number]['age'];
```

`as const` makes const value readonly literal values

## `--strictPropertyInitialization`

```ts
class BadGreeter {
  name: string;
}

class GoodGreeter {
  name: string;

  constructor() {
    this.name = 'hello';
  }
}

class OKGreeter {
  // Not initialized, but no error
  name!: string;
}
```

## Unions and

- Unions with common fileds

```ts
interface Bird {
  fly(): void;
  layEggs(): void;
}

interface Fish {
  swim(): void;
  layEggs(): void;
}

declare function getSmallPet(): Fish | Bird;

let pet = getSmallPet();
pet.layEggs();

// Only available in one of the two possible types
pet.swim(); // ERROR
```

- Intersection Types

```ts
interface ErrorHandling {
  success: boolean;
  error?: { message: string };
}

interface ArtworksData {
  artworks: { title: string }[];
}

interface ArtistsData {
  artists: { name: string }[];
}

// These interfaces are composed to have
// consistent error handling, and their own data.

type ArtworksResponse = ArtworksData & ErrorHandling;
type ArtistsResponse = ArtistsData & ErrorHandling;
```

- re export

```ts
// index.d.ts
export { default as CSSTransition } from './CSSTransition';
```

## nodemon for ts

- nodemon.json
- execute with only `nodemon`

```json
{
  "watch": ["server"],
  "ext": "ts",
  "exec": "ts-node index.ts"
}
```