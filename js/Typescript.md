# Typescript basic

- [handbook](https://microsoft.github.io/TypeScript-New-Handbook/everything)
- [playground](https://www.typescriptlang.org/play)
- [type challenges](https://github.com/type-challenges/type-challenges)
- [ts blogs](https://devblogs.microsoft.com/typescript/)
- [Microsoft office tutorial](https://docs.microsoft.com/en-us/learn/paths/build-javascript-applications-typescript/)
- [arrow function in class](https://stackoverflow.com/questions/50325925/using-es6-arrow-functions-inside-class)

----
- [Typescript basic](#typescript-basic)
  - [BFE.dev ts puzzles](#bfedev-ts-puzzles)
      - [1. implement Partial<T>](#1-implement-partialt)
      - [2. implement Required<T>](#2-implement-requiredt)
      - [3. implement Readonly<T>](#3-implement-readonlyt)
      - [4. implement Record<K, V>](#4-implement-recordk-v)
      - [5. implement Pick<T, K>](#5-implement-pickt-k)
      - [6. implement Omit<T, K>](#6-implement-omitt-k)
      - [7. implement Exclude<T, E>](#7-implement-excludet-e)
      - [8. implement Extract<T, U>](#8-implement-extractt-u)
      - [9. implement NonNullable<T>](#9-implement-nonnullablet)
      - [10. implement Parameters<T>](#10-implement-parameterst)
      - [11. implement ConstructorParameters<T>](#11-implement-constructorparameterst)
      - [12. implement ReturnType<T>](#12-implement-returntypet)
      - [13. implement InstanceType<T>](#13-implement-instancetypet)
      - [14. implement ThisParameterType<T>](#14-implement-thisparametertypet)
      - [15. implement OmitThisParameter<T>](#15-implement-omitthisparametert)
      - [16. implement FirstChar<T>](#16-implement-firstchart)
      - [17. implement LastChar<T>](#17-implement-lastchart)
      - [18. implement TupleToUnion<T>](#18-implement-tupletouniont)
      - [19. implement FirstItem<T>](#19-implement-firstitemt)
      - [20: implement IsNever<T>](#20-implement-isnevert)
      - [21. implement LastItem<T>](#21-implement-lastitemt)
      - [22. implement StringToTuple<T>](#22-implement-stringtotuplet)
      - [23. implement LengthOfTuple<T>](#23-implement-lengthoftuplet)
      - [24. implement LengthOfString<T>](#24-implement-lengthofstringt)
      - [25. implement UnwrapPromise<T>](#25-implement-unwrappromiset)
      - [26. implement ReverseTuple<T>](#26-implement-reversetuplet)
      - [27. implement Flat<T>](#27-implement-flatt)
      - [29. implement Shift<T>](#29-implement-shiftt)
      - [31. implement Push<T, I>](#31-implement-pusht-i)
      - [33. implement TupleToString<T>](#33-implement-tupletostringt)
      - [42. implement Equal<A, B>](#42-implement-equala-b)
  - [Types](#types)
    - [keywords](#keywords)
      - [keyof](#keyof)
      - [never](#never)
      - [`infer`](#infer)
    - [concepts](#concepts)
      - [Mapped Types](#mapped-types)
      - [Conditional Type](#conditional-type)
      - [Indexed Access Types](#indexed-access-types)
      - [`as const`](#as-const)
      - [`--strictPropertyInitialization`](#--strictpropertyinitialization)
      - [Unions](#unions)
      - [Intersection Types](#intersection-types)
  - [others](#others)
    - [Decorator](#decorator)
    - [For those packages that have no type definition](#for-those-packages-that-have-no-type-definition)
    - [Add variable to globle object](#add-variable-to-globle-object)
    - [re export](#re-export)
    - [nodemon for ts](#nodemon-for-ts)
  - [tsconfig.json](#tsconfigjson)

----

## BFE.dev ts puzzles

#### 1. implement Partial<T>

```ts
type MyPartial<T> = {
  [K in keyof T]?: T[K]
}
```

#### 2. implement Required<T>

```ts
type MyRequired<T> = {
  [K in keyof T]-?: T[K]
}
```

#### 3. implement Readonly<T>

```ts
type MyRequired<T> = {
  readonly [K in keyof T]: T[K]
}
```

#### 4. implement Record<K, V>

```ts
type MyRecord<K extends keyof any, V> = {
  [P in K]: V
}

```

#### 5. implement Pick<T, K>

```ts
type MyPick<T, K extends keyof T> = {
  [P in K]: T[P]
}

```

#### 6. implement Omit<T, K>

```ts
type MyExclude<T, U> = T extends U ? never : T;
type MyOmit<T, K extends keyof T> = {
  [P in MyExclude<keyof T, K>]: P extends K ?  never :T[P]
}

```

#### 7. implement Exclude<T, E>

```ts
type MyExclude<T, E> = T extends E ? never : T;
```

#### 8. implement Extract<T, U>

```ts
type MyExtract<T, U> = T extends U ? T : never;
```

#### 9. implement NonNullable<T>

```ts
type MyNonNullable<T> = T extends null | undefined ? never : T;
```

#### 10. implement Parameters<T>

```ts
type MyParameters<Fn extends Function> = Fn extends (...args: infer U) => any ? U : never;
type MyParameters<T extends (...args: any[]) => any> = T extends (...args: infer P) => any ? P : never;
```

#### 11. implement ConstructorParameters<T>

```ts
type MyConstructorParameters<T extends new (...args: any) => any> =  T extends new (...args: infer P) => any ? P : never;
```

#### 12. implement ReturnType<T>

```ts
type MyReturnType<T extends Function> = T extends (...args: any) => infer R ? R : never;
```

#### 13. implement InstanceType<T>

```ts
type MyInstanceType<T extends new (...args: any) => any> = T extends new (...args: any) => infer R ? R : never;
```

#### 14. implement ThisParameterType<T>

```ts
type MyThisParameterType<T> = T extends (this: infer R, ...args: any) => any ? R : unknown;
```
#### 15. implement OmitThisParameter<T>

```ts
type MyOmitThisParameter<T> = T extends (this: any, ...args: infer U) => infer R ? (...args: U) => R : never;
```

#### 16. implement FirstChar<T>
```ts
type FirstChar<T extends string> =  T extends `${infer T}${any}` ? T : never;
```

#### 17. implement LastChar<T>

```ts
type LastChar<T extends string> =
  T extends `${infer F}${infer R}` ?
   R extends '' ? F : LastChar<R> : never;

```

#### 18. implement TupleToUnion<T>

```ts
type TupleToUnion<T extends any[]> = T[number]
```

#### 19. implement FirstItem<T>

```ts
type FirstItem<T extends any[]> = T[0] extends undefined ? never : T[0];
```

#### 20: implement IsNever<T>

```ts
type IsNever<T> = [T] extends [never] ? true: false;
```

#### 21. implement LastItem<T>
- solution one

```ts
type LastItem<T extends any[]> = T extends [...a: any, b: infer L] ? L : never;
```

- solution two

```ts
type LastItem<T extends any[]> = T extends [infer F, ...infer L] ? L extends [] ? F :  LastItem<L> : never;
```

#### 22. implement StringToTuple<T>

```ts
type StringToTuple<T extends string> =  T extends `${infer F}${infer L}` ? [F, ...StringToTuple<L>] : [];
```

#### 23. implement LengthOfTuple<T>

```ts
type LengthOfTuple<T extends any[]> = T["length"]

```

#### 24. implement LengthOfString<T>

```ts
type StringToTuple<T extends string> = T extends `${infer P}${infer U}` ? [P, ...StringToTuple<U>] : [];

type LengthOfString<T extends string> = StringToTuple<T>['length'];
```


#### 25. implement UnwrapPromise<T>

```ts
type UnwrapPromise<T> = T extends Promise<infer U> ? U : Error;
```

#### 26. implement ReverseTuple<T>

```ts
type ReverseTuple<T extends any[]> = T extends [...infer F, infer L] ? [L, ...ReverseTuple<F>] : [];
```


#### 27. implement Flat<T>

```ts
type Flat<T extends any[]> =
  T extends [infer P, ...infer U] ?
    P extends any[] ?
    [...Flat<P>, ...Flat<U>] :
  [P, ...Flat<U>] : [];

```

#### 29. implement Shift<T>

```ts
type Shift<T extends any[]> = T extends [infer T, ...infer U] ? U : [];
```

#### 31. implement Push<T, I>

```ts
type Push<T extends any[], I> = [...T, I];
```

#### 33. implement TupleToString<T>

```ts
type TupleToString<T extends string[]> =
  T extends [infer P, ...infer U] ?
    U extends string[] ?
    `${P & string}${TupleToString<U>}`:
    P:
  '';
```

#### 42. implement Equal<A, B>

```ts

type Equal<A, B> =
                (<T>() => T extends A ? 1: 0) extends
                (<T>() => T extends B ? 1 : 0)
                ? true
                : false;

```


## Types

### keywords

#### keyof

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

#### never

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

#### `infer`

```ts
type IsArray<T> = T extends Array<infer O> ? O : T;

type ItIsArray = IsArray<string[]>; // string
type ItIsNotArray = IsArray<boolean>; // boolean
```

### concepts

#### Mapped Types

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

#### Conditional Type

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

#### Indexed Access Types

```ts
const MyArray = [
  { name: 'Alice', age: 15 },
  { name: 'Bob', age: 23 },
  { name: 'Eve', age: 38 }
] as const;

type Ages = typeof MyArray[number]['age'];
```

#### `as const`

makes const value readonly literal values

#### `--strictPropertyInitialization`

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

#### Unions

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

#### Intersection Types

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

## others

### Decorator

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
### For those packages that have no type definition

- putting an empty declaration for it in a ` .d.ts` file in your project like below.

```ts
declare module 'some-untyped-module';
```

### Add variable to globle object

```ts
declare global {
  interface Window {
    Hello: {
      moneyFormatter: (price: string) => string;
    };
  }
}
```

### re export

```ts
// index.d.ts
export { default as CSSTransition } from './CSSTransition';
```

### nodemon for ts

- nodemon.json
- execute with only `nodemon`

```json
{
  "watch": ["server"],
  "ext": "ts",
  "exec": "ts-node index.ts"
}
```

## tsconfig.json

- default generated by `tsc --init`

```jsonp
{
  "compilerOptions": {
    /* Visit https://aka.ms/tsconfig.json to read more about this file */

    /* Basic Options */
    // "incremental": true,                   /* Enable incremental compilation */
    "target": "es5" /* Specify ECMAScript target version: 'ES3' (default), 'ES5', 'ES2015', 'ES2016', 'ES2017', 'ES2018', 'ES2019', 'ES2020', or 'ESNEXT'. */,
    "module": "commonjs" /* Specify module code generation: 'none', 'commonjs', 'amd', 'system', 'umd', 'es2015', 'es2020', or 'ESNext'. */,
    // "lib": [],                             /* Specify library files to be included in the compilation. */
    // "allowJs": true,                       /* Allow javascript files to be compiled. */
    // "checkJs": true,                       /* Report errors in .js files. */
    // "jsx": "preserve",                     /* Specify JSX code generation: 'preserve', 'react-native', or 'react'. */
    // "declaration": true,                   /* Generates corresponding '.d.ts' file. */
    // "declarationMap": true,                /* Generates a sourcemap for each corresponding '.d.ts' file. */
    // "sourceMap": true,                     /* Generates corresponding '.map' file. */
    // "outFile": "./",                       /* Concatenate and emit output to single file. */
    // "outDir": "./",                        /* Redirect output structure to the directory. */
    // "rootDir": "./",                       /* Specify the root directory of input files. Use to control the output directory structure with --outDir. */
    // "composite": true,                     /* Enable project compilation */
    // "tsBuildInfoFile": "./",               /* Specify file to store incremental compilation information */
    // "removeComments": true,                /* Do not emit comments to output. */
    // "noEmit": true,                        /* Do not emit outputs. */
    // "importHelpers": true,                 /* Import emit helpers from 'tslib'. */
    // "downlevelIteration": true,            /* Provide full support for iterables in 'for-of', spread, and destructuring when targeting 'ES5' or 'ES3'. */
    // "isolatedModules": true,               /* Transpile each file as a separate module (similar to 'ts.transpileModule'). */

    /* Strict Type-Checking Options */
    "strict": true /* Enable all strict type-checking options. */,
    // "noImplicitAny": true,                 /* Raise error on expressions and declarations with an implied 'any' type. */
    // "strictNullChecks": true,              /* Enable strict null checks. */
    // "strictFunctionTypes": true,           /* Enable strict checking of function types. */
    // "strictBindCallApply": true,           /* Enable strict 'bind', 'call', and 'apply' methods on functions. */
    // "strictPropertyInitialization": true,  /* Enable strict checking of property initialization in classes. */
    // "noImplicitThis": true,                /* Raise error on 'this' expressions with an implied 'any' type. */
    // "alwaysStrict": true,                  /* Parse in strict mode and emit "use strict" for each source file. */

    /* Additional Checks */
    // "noUnusedLocals": true,                /* Report errors on unused locals. */
    // "noUnusedParameters": true,            /* Report errors on unused parameters. */
    // "noImplicitReturns": true,             /* Report error when not all code paths in function return a value. */
    // "noFallthroughCasesInSwitch": true,    /* Report errors for fallthrough cases in switch statement. */

    /* Module Resolution Options */
    // "moduleResolution": "node",            /* Specify module resolution strategy: 'node' (Node.js) or 'classic' (TypeScript pre-1.6). */
    // "baseUrl": "./",                       /* Base directory to resolve non-absolute module names. */
    // "paths": {},                           /* A series of entries which re-map imports to lookup locations relative to the 'baseUrl'. */
    // "rootDirs": [],                        /* List of root folders whose combined content represents the structure of the project at runtime. */
    // "typeRoots": [],                       /* List of folders to include type definitions from. */
    // "types": [],                           /* Type declaration files to be included in compilation. */
    // "allowSyntheticDefaultImports": true,  /* Allow default imports from modules with no default export. This does not affect code emit, just typechecking. */
    "esModuleInterop": true /* Enables emit interoperability between CommonJS and ES Modules via creation of namespace objects for all imports. Implies 'allowSyntheticDefaultImports'. */,
    // "preserveSymlinks": true,              /* Do not resolve the real path of symlinks. */
    // "allowUmdGlobalAccess": true,          /* Allow accessing UMD globals from modules. */

    /* Source Map Options */
    // "sourceRoot": "",                      /* Specify the location where debugger should locate TypeScript files instead of source locations. */
    // "mapRoot": "",                         /* Specify the location where debugger should locate map files instead of generated locations. */
    // "inlineSourceMap": true,               /* Emit a single file with source maps instead of having a separate file. */
    // "inlineSources": true,                 /* Emit the source alongside the sourcemaps within a single file; requires '--inlineSourceMap' or '--sourceMap' to be set. */

    /* Experimental Options */
    // "experimentalDecorators": true,        /* Enables experimental support for ES7 decorators. */
    // "emitDecoratorMetadata": true,         /* Enables experimental support for emitting type metadata for decorators. */

    /* Advanced Options */
    "skipLibCheck": true /* Skip type checking of declaration files. */,
    "forceConsistentCasingInFileNames": true /* Disallow inconsistently-cased references to the same file. */
  }
}
```

- react

```json
{
  "compilerOptions": {
    "baseUrl": ".",
    "outDir": "build/dist",
    "module": "esnext",
    "target": "es5",
    "lib": ["es6", "dom", "esnext.asynciterable"],
    "sourceMap": true,
    "allowJs": true,
    "jsx": "react",
    "moduleResolution": "node",
    "rootDir": "src",
    "forceConsistentCasingInFileNames": true,
    "noImplicitReturns": true,
    "noImplicitThis": true,
    "noImplicitAny": true,
    "strictNullChecks": true,
    "suppressImplicitAnyIndexErrors": true,
    "noUnusedLocals": true,
    "skipLibCheck": true
  },
  "exclude": [
    "node_modules",
    "build",
    "scripts",
    "acceptance-tests",
    "webpack",
    "jest",
    "src/setupTests.ts"
  ]
}
```

- node

```json
{
  "compilerOptions": {
    "target": "es6",
    "module": "commonjs",
    "lib": ["dom", "es6", "es2017", "esnext.asynciterable"],
    "sourceMap": true,
    "outDir": "./dist",
    "moduleResolution": "node",
    "removeComments": true,
    "noImplicitAny": true,
    "strictNullChecks": true,
    "strictFunctionTypes": true,
    "noImplicitThis": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noImplicitReturns": true,
    "noFallthroughCasesInSwitch": true,
    "allowSyntheticDefaultImports": true,
    "esModuleInterop": true,
    "emitDecoratorMetadata": true,
    "experimentalDecorators": true,
    "resolveJsonModule": true,
    "baseUrl": "."
  },
  "exclude": ["node_modules"],
  "include": ["./src/**/*.tsx", "./src/**/*.ts"]
}
```

- react native

```json
{
  "compilerOptions": {
    "target": "es5",
    "module": "commonjs",
    "jsx": "react-native",
    "lib": ["es6", "esnext.asynciterable"],
    "strict": true,
    "esModuleInterop": true,
    "rootDir": "src",
    "forceConsistentCasingInFileNames": true,
    "noImplicitReturns": true,
    "noImplicitThis": true,
    "noImplicitAny": true,
    "strictNullChecks": true,
    "suppressImplicitAnyIndexErrors": true,
    "noUnusedLocals": true,
    "skipLibCheck": true,
    "baseUrl": "."
  }
}
```
