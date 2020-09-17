## JS

### Basic

#### Iterator

```js
let i = [1, 2, 3, 4];

let iterator = i[Symbol.iterator]();

console.log(iterator.next());
console.log(iterator.next());
console.log(iterator.next());
console.log(iterator.next());
```

#### Generator

```js
function* generator() {
  yield 1;
  yield 2;
  yield 3;
  yield 4;
}

let iterator = generator();

console.log(iterator.next());
console.log(iterator.next());
console.log(iterator.next());
console.log(iterator.next());
console.log(iterator.next());
```

#### Map

- Map
  > The Map object holds key-value pairs and **remembers** the original insertion order of the keys. Any value (both **objects** and primitive values) may be used as either a key or a value.

```js
const a = {};
const b = { num: 1 };

const map = new Map();

map.set(a, 'a').set(b, 'b').set(a, 'c');
// map.delete(b);
// map.get(a)

for (let [key, value] of map.entries()) {
  console.log(key, value);
}

const arr = [...map]; // convert to array

console.log(map);
```

- WeakMap

> [info](https://javascript.info/weakmap-weakset)

> The first difference from Map is that WeakMap keys must be **objects**, not primitive values

> Key won't be garbage-collected

#### Set

- Set

> The Set object lets you store **unique** values of any type, whether primitive values or object references.

> Set objects are collections of values. You can iterate through the elements of a set in **insertion order**. A value in the Set may **only occur once**; it is unique in the Set's collection.

```js
const set = new Set([1, 2, 4]);
set.add(1).add(2);
// set has no order
for (let val of set) {
  console.log(val);
}
console.log(set);
```

```js
const arr = [1, 2, 2, 3];
const newArr = [...new Set[arr]()]; //[1,2,3]
const newArr = [Array.from(new Set[arr]())]; //[1,2,3]
```

- WeakSet

```js
const ws = new WeakSet([{ a: 1 }, { b: 2 }]);

for (let val of ws) {
  console.log(val); // ERROR => WeakSet is not iterable.
}
```

### Nodejs

- `Ctrl + L` clean screen, similar to `process.stdout.write("\u001b[2J\u001b[0;0H");`

#### Generate zip file from string

```ts
import jszip from 'jszip';

const createZip = async (code: string): Promise<void> => {
  const zip = new jszip();
  zip.file('index.js', code);
  const content = zip.generateAsync({ type: 'string' });
  console.log(content);
  zip
    .generateNodeStream({ type: 'nodebuffer', streamFiles: true })
    .pipe(fs.createWriteStream('./dist/out.zip'))
    .on('finish', function () {
      console.log('out.zip written.');
    });
};
```

#### TS decorator

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

### Download file

```js
const filename = '1.txt';
const a = document.createElement('a');
const blob = new Blob(['hello world!']);
a.download = filename;
a.href = URL.createObjectURL(blob);
a.click();
URL.revokeObjectURL(blob);
```

### Cors

```
[withCredentials=true] => [Access-Control-Allow-Origin] must be the address

[withCredentials=true] => [Access-Control-Allow-Credentials] must be configured

if request header is added, [Access-Control-Allow-Headers] must be allowed

```
