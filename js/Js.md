# Javascript Basic

- [javascript runtime simulation](http://latentflip.com/loupe)
- [Concept explanation](https://www.javascripttutorial.net/)
- [chinese resource](https://github.com/coffe1891/frontend-hard-mode-interview)

- [Javascript Basic](#javascript-basic)
  - [Proxy](#proxy)
  - [With](#with)
  - [Iterator](#iterator)
  - [Generator](#generator)
  - [Map](#map)
  - [Set](#set)
  - [Async javascript](#async-javascript)
    - [Callback](#callback)
    - [Promise](#promise)
      - [Promise.all, Promise.race](#promiseall-promiserace)
    - [async/await](#asyncawait)
  - [IIFE(Immediately Invoked Function Expression)](#iifeimmediately-invoked-function-expression)
  - [this](#this)
    - [bind, call, apply](#bind-call-apply)
    - [constructor](#constructor)
  - [`this`](#this-1)
    - [Global context](#global-context)
    - [Function context](#function-context)
    - [Function invocation](#function-invocation)
    - [Method invocation](#method-invocation)
    - [Constructor invocation](#constructor-invocation)
    - [Indirect invocation](#indirect-invocation)
    - [Arrow functions](#arrow-functions)
  - [Prototype, **proto**, prototypal inheritance](#prototype-proto-prototypal-inheritance)
  - [Linked List](#linked-list)
  - [React useState](#react-usestate)
  - [`slice` `splice` `splite`](#slice-splice-splite)
  - [Generator and Iterator](#generator-and-iterator)
      - [loop though array object](#loop-though-array-object)
      - [iterator](#iterator-1)
      - [generator](#generator-1)
      - [use generator to create own Object.entries function](#use-generator-to-create-own-objectentries-function)

## Proxy

- [MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Proxy)

> The Proxy object enables you to create a proxy for another object, which can intercept and redefine fundamental operations
> for that object.

```js
let data = { count: 1 };

let proxy = new Proxy(data, {
  get(target, key) {
    console.log('getting...');
    return target[key];
  },
  set(target, key, value) {
    console.log('setting...');
    target[key] = value;
    return true;
  }
});
```

## With

- [MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/with)
  > The with statement extends the scope chain for a statement.

```js
let person = {
  name: 'JANE',
  age: 25
};

let expression = '`${name} is ${age} years old`';

with (person) {
  console.log(eval(expression));
}
```

## Iterator

```js
let i = [1, 2, 3, 4];

let iterator = i[Symbol.iterator]();

console.log(iterator.next());
console.log(iterator.next());
console.log(iterator.next());
console.log(iterator.next());
```

## Generator

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

## Map

- Map

The Map object holds key-value pairs and **remembers** the original insertion order of the keys. Any value (both **objects** and primitive values) may be used as either a key or a value.

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

1. The first difference from Map is that WeakMap keys must be **objects**, not primitive values
2. if we use an object as the key in it, and there are no other references to that object – it will be removed from memory (and from the map) automatically
3. WeakMap does not support iteration and methods keys(), values(), entries(), so there’s no way to get all keys or values from it.

```js
let weakMap = new WeakMap();

let obj = {};

weakMap.set(obj, "ok"); // works fine (object key)

// can't use a string as the key
weakMap.set("test", "Whoops"); // Error, because "test" is not an object
```


```js
let john = { name: "John" };

let weakMap = new WeakMap();
weakMap.set(john, "...");

john = null; // overwrite the reference

// john is removed from memory!
```

## Set

- Set

1. The Set object lets you store **unique** values of any type, whether primitive values or object references.

2. Set objects are collections of values. You can iterate through the elements of a set in **insertion order**. A value in the Set may **only occur once**; it is unique in the Set's collection.

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
const newArr = [...new Set(arr)]; //[1,2,3]
const newArr = [Array.from(new Set[arr]())]; //[1,2,3]
```

- WeakSet

> [info](https://javascript.info/weakmap-weakset)

1. It is analogous to `Set`, but we may only add objects to `WeakSet` (not primitives).
2. An object exists in the set while it is reachable from somewhere else.
3. Like `Set`, it supports `add`, `has` and `delete`, but not `size`, `keys()` and no iterations.

```js
const ws = new WeakSet([{ a: 1 }, { b: 2 }]);

for (let val of ws) {
  console.log(val); // ERROR => WeakSet is not iterable.
}
```

## Async javascript

### Callback

```js
let greeting = (name) => cosnole.log(`Hello ${name}!`);

const userInfo = (firstName, lastName, callback) => {
  const fullName = `${firstName} ${lastName}`;
  callback(fullName);
};

userInfo('John', 'Doe', greeting);
```

### Promise

```js
const hasMeeting = false;
const meeting = new Promise((resolve, reject) => {
  if (!hasMeeting) {
    const meetingDetails = {
      name: 'Maketing Meeting',
      location: 'Skype',
      time: '1:00 pm'
    };
    resolve(meetingDetails);
  } else {
    reject(new Error('Meeting already scheduled'));
  }
});

meeting.then((res) => console.log(res)).catch((err) => console.error(err));
```

- chain promise

```js
const addToCalendar = (meetingDetails) => {
  const calendar = `${meetingDetails.name} is scheduled at ${meetingDetails.time} on ${meetingDetails.location}`;
  return Promise.resolve(calendar);
};

meeting
  .then(addToCalendar)
  .then((res) => console.log(res))
  .catch((err) => console.error(err));
```

#### Promise.all, Promise.race

```js
const promise1 = Promise.resolve('Promise 1 complete');
const promise2 = Promise.resolve('Promise 2 complete');

promise1.then((res) => console.log(res));
promise2.then((res) => console.log(res));
// Promise 1 complete
// Promise 2 complete

// excute when all promise fulfilled
Promise.all([promise1, promise2]).then((res) => console.log(res));
// ["Promise 1 complete", "Promise 2 complete"]

// excute when first promise fulfilled
Promise.race([promise1, promise2]).then((res) => console.log(res));
// Promise 1 complete
```

### async/await

```js
const myMeeting = async () => {
  try {
    const meetingDetails = await meeting;
    const message = await addToCalendar(meetingDetails);
    console.log(message);
  } catch (err) {
    console.log(err.message);
  }
};
myMeeting();
```

## IIFE(Immediately Invoked Function Expression)

- normal function

```js
function multiply(a, b) {
  return a * b;
}
console.log(mutiply(2, 5));
```

- IIFE function

```js
(function (a, b) {
  return a * b;
})(2, 5);
```

- IIFE arrow function

```js
((a, b) => {
  return a * b;
})(2, 5);
```

- Variable scope with IIFE

```js
var name = 'JOE'(function (a, b) {
  var name = 'BOB';
  console.log(name); // BOB
})(2, 5);
console.log(name); // JOE
```

- Compare with normal situation and es6

```js
var name = 'JOE';
{
  var name = 'BOB';
  console.log(name); // BOB
}
console.log(name); // BOB
```

```js
let name = 'JOE';
{
  let name = 'BOB';
  console.log(name); // BOB
}
console.log(name); // JOE
```

## this

- [this](https://web.dev/javascript-this/)

- `this` === current execution context

```js
const user = {
  name: 'john',
  whodis: function () {
    console.log(this); // user
  },
  butWhoAmI: () => console.log(this) // global
};
```

### bind, call, apply

```js
function showFace() {
  return this.face;
}
const user = {
  face: 'smile'
};

const showUserFace = showFace.bind(user);

console.log(showUserFace());
console.log(showFace.call(user, 1, 2, 3));
console.log(showFace.call(user, ...[1, 2, 3]));
console.log(showFace.apply(user, [1, 2, 3]));
```

### constructor

```js
function Horse(name) {
  this.name = name;
  this.voice = function () {
    console.log('yoyo');
    return this;
  };
}

const myHorse = new Horse('yuyu');
myHorse.voice().voice();
```

## `this`

### Global context

```js
console.log(this === globalThis); // true
```

### Function context

- Function invocation
- Method invocation
- Constructor invocation
- Indirect invocation

### Function invocation

- non-strict mode

```js
function a() {
  console.log(this === globalThis);
}

a(); // true
```

- strict mode

```js
'use strict';
function a() {
  console.log(this);
}

a(); // undefine
```

### Method invocation

```js
const Car = {
  brand: 'Toyota',
  getBrand: function () {
    return this.brand;
  }
};

console.log(Car.getBrand()); // Toyota

const car = Car.getBrand;

console.log(car()); // undefine
```

- to solve this problem, use `bind()` method of `Function.prototype`

```js
const b = Car.getBrand.bind();
```

- however it's able to bind other object to this

```js
const Car = {
  brand: 'Toyota',
  getBrand: function () {
    return this.brand;
  }
};

const Mobile = {
  brand: 'apple'
};

console.log(Car.getBrand());

const car = Car.getBrand.bind(Mobile);

console.log(car());
```

### Constructor invocation

```js
function Car(brand) {
  this.brand = brand;
}

Car.prototype.getBrand = function () {
  return this.brand;
};

var car = new Car('Honda');
console.log(car.getBrand());
```

```js
var bmw = Car('BMW');
console.log(bmw.brand);
// => TypeError: Cannot read property 'brand' of undefined
```

```js
function Car(brand) {
  // if (!(this instanceof Car)) {
  //     throw Error('Must use the new operator to call the function');
  // }
  if (!new.target) {
    throw Error('Must use the new operator to call the function');
  }
  this.brand = brand;
}
```

### Indirect invocation

```js
function getBrand(prefix) {
  console.log(prefix + this.brand);
}

let honda = {
  brand: 'Honda'
};
let audi = {
  brand: 'Audi'
};

getBrand.call(honda, "It's a "); // It's a Honda
getBrand.call(audi, "It's an "); // It's an Audi
```

```js
getBrand.apply(honda, ["It's a "]); // "It's a Honda"
getBrand.apply(audi, ["It's an "]); // "It's a Audi"
```

### Arrow functions

```js
let getThis = () => this;
console.log(getThis() === window); // true
```

```js
function Car() {
  this.speed = 120;
}

Car.prototype.getSpeed = () => {
  return this.speed;
};

var car = new Car();
car.getSpeed(); // TypeError
```

## Prototype, **proto**, prototypal inheritance

`.prototype` only exist on function

```js
var a = function () {};
var b = [1, 2, 3];

console.log(a.prototype); //>> function(){}
console.log(b.prototype); //>> undefined
```

```js
var a = function () {};
var b = [1, 2, 3];

console.log(a.__proto__ === Function.prototype); //>> true
console.log(b.__proto__ === Array.prototype); //>> true
console.log(a.__proto__ === a.__proto__.constructor.prototype); // true
console.log(b.__proto__ === b.__proto__.constructor.prototype); //true

console.log(a.__proto__.__proto__ === Object.prototype); //>> true
console.log(b.__proto__.__proto__ === Object.prototype); //>> true

console.log(new Object().__proto__.__proto__); //>> null
console.log(Object.prototype.__proto__); //>> null
```

`__proto__` of a value is the prototype of constructor. //一个对象的原型就是它的构造函数的 prototype 属性的值

```js
let person = {
  name: 'John Doe',
  greet: function () {
    return "Hi, I'm " + this.name;
  }
};

console.log(person instanceof Object); // true
console.log(person.toString()); // [object Object]
```

`person` no toString() method, so find in prototype chaine, then excute this `Object.prototype.toString()`

```js
let person = {
  name: 'John Doe',
  greet: function () {
    return "Hi, I'm " + this.name;
  }
};

let teacher = {};

teacher.__proto__ = person;

console.log(teacher.name);
console.log(Object.getPrototypeOf(teacher) === person); //true
```

```js
let person = {
  name: 'John Doe',
  greet: function () {
    return "Hi, I'm " + this.name;
  }
};

let teacher = Object.create(person);

console.log(Object.getPrototypeOf(teacher) === person); //true
```

`class`

```js
class A {}

A === A.prototype.constructor;
```

```js
class A {}

class B extends A {}

console.log(B.__proto___ === A); // true
B.prototype.__proto__ === A.prototype; //>> true
```

## Linked List

```js
class Node {
  constructor(data, next = null) {
    this.data = data;
    this.next = next;
  }
}

class LinkedList {
  constructor() {
    this.head = null;
    this.size = 0;
  }

  // Insert first node
  insertFirst(data) {
    this.head = new Node(data, this.head);
    this.size++;
  }

  // Insert last node
  insertLast(data) {
    let node = new Node(data);
    let current;
    if (!this.head) {
      this.head = node;
    } else {
      current = this.head;
      while (current.next) {
        current = current.next;
      }
      current.next = node;
    }

    this.size++;
  }

  // Insert at index
  insertAt(data, index) {
    // If index is out of range
    if (index > 0 && index > this.size) return;

    if (index === 0) {
      this.insertFirst(data);
    }

    const node = new Node(data);
    let current, previous;
    current = this.head;

    let count = 0;
    while (count < index) {
      previous = current; // node before index
      current = current.next;
      count++;
    }
    node.next = current;
    previous.next = node;
    this.size++;
  }

  // Get at index

  getIndex(index) {
    let current = this.head;
    let count = 0;
    while (current) {
      if (count === index) {
        console.log({ index: `index ${index}: ${current.data}` });
      }
      current = current.next;
      count++;
    }
    return null;
  }

  // Remove at index
  removeAt(index) {
    if (index > this.size) return;

    let current = this.head;
    let previous;
    let count = 0;

    if (!index) {
      this.head = current;
    } else {
      while (count < index) {
        previous = current;
        current = current.next;
        count++;
      }
      previous.next = current.next;
    }
    this.size--;
  }

  // Clear list
  clearList() {
    this.head = null;
    this.size = 0;
  }

  // Print list data
  printListData() {
    let current = this.head;
    while (current) {
      console.log(current.data);
      current = current.next;
    }
    if (!this.size) console.log('nothing');
  }
}

const ll = new LinkedList();

ll.insertFirst(400);
ll.insertFirst(300);
ll.insertFirst(200);
ll.insertFirst(100);
ll.removeAt(2);
ll.printListData();
ll.clearList();
ll.printListData();
```

## React useState

```js
let isMount = true;
let workInProgressHook = null;

const fiber = {
  stateNode: App,
  memoizedState: null
};

function schedule() {
  workInProgressHook = fiber.memoizedState;
  const app = fiber.stateNode();
  isMount = false;
  return app;
}

function useState(initialState) {
  let hook;
  if (isMount) {
    hook = {
      memoizedState: initialState,
      next: null,
      queue: {
        pending: null
      }
    };
    if (!fiber.memoizedState) {
      fiber.memoizedState = hook;
    } else {
      workInProgressHook.next = hook;
    }
    workInProgressHook = hook;
  } else {
    hook = workInProgressHook;
    workInProgressHook = workInProgressHook.next;
  }
  let baseState = hook.memoizedState;

  if (hook.queue.pending) {
    let firstUpdate = hook.queue.pending.next;
    do {
      const action = firstUpdate.action;
      baseState = action(baseState);
      firstUpdate = firstUpdate.next;
    } while (firstUpdate !== hook.queue.pending.next);
    hook.queue.pending = null;
  }
  hook.memoizedState = baseState;
  return [baseState, dispatchAction.bind(null, hook.queue)];
}

function dispatchAction(queue, action) {
  const update = {
    action,
    next: null
  };

  if (queue.pending === null) {
    update.next = update;
  } else {
    update.next = queue.pending.next;
    queue.pending.next = update;
  }
  queue.pending = update;
  schedule();
}

function App() {
  const [num, updateNum] = useState(0);
  console.log({ isMount, num });
  return {
    onClick() {
      updateNum((a) => a + 1);
    }
  };
}

window.app = schedule;
```

## `slice` `splice` `splite`

- `slice()`

  - Copies elements from an array
  - Returns them as a new array
  - Doesn’t change the original array
  - Starts slicing from … until given index: array.slice (from, until)
  - Slice doesn’t include “until” index parameter
  - Can be used both for arrays and strings

- `Splice()`

  - Used for adding/removing elements from array
  - Returns an array of removed elements
  - Changes the array
  - For adding elements: array.splice (index, number of elements, element)
  - For removing elements: array.splice (index, number of elements)
  - Can only be used for arrays

- `Split()`

  - Divides a string into substrings
  - Returns them in an array
  - Takes 2 parameters, both are optional: string.split(separator, limit)
  - Doesn’t change the original string
  - Can only be used for strings

## Generator and Iterator

#### loop though array object

```js
const person = { name: 'eriii', phone: '123-2312' };
for (const [key, value] of Object.entries(person)) {
  console.log(`key: ${key}is ${value}`);
}
// > "key: nameis eriii"
// > "key: phoneis 123-2312"
```

#### iterator

```js
const names = ['john', 'joe', 'mm'];

for (const name in names) {
  console.log(name);
}
// 1 2 3

for (const name of names) {
  console.log(name);
}
// john joe mm

const iterator = names[Symbol.iterator]();
iterator.next(); // {value: "john", done: false}
iterator.next(); // {value: "joe", done: false}
iterator.next(); // {value: "mm", done: false}
iterator.next(); // {value: undefined, done: true}
```

#### generator

```js
const person = {name: 'eriii', phone: '123-2312'}
// person[Sy]

person[Symbol.iterator] = function* (){
  yield Object.keys(this)
}

[...person] // [['name', 'phone']]

for(const b of person){
  b // [name, phone]
}

person[Symbol.iterator] = function* (){
  yield* Object.keys(this)
}


[...person] // ['name', 'phone']

for(const b of person){
  b // name, phone
}
```

```js
function* generatorThings() {
  yield 'hello';
  const yo = yield 'hey';
  console.log(yo); //> yoyo
  yield 'hi';
  return 'return';
}

const g = generatorThings();

g.next(); // > {value: "hello", done: false}
g.next('yo'); // > {value: "hey", done: false}
g.next('yoyo'); // > {value: "hi", done: false}
const m = g.next(); // > {value: undefined, done: true}
m.done;
```
#### use generator to create own Object.entries function

```js

const person = { name: 'eriii', phone: '123-2312' };

function* oentries(obj) {
  const keys = Object.keys(obj)
  for (const key of keys){
    yield [key, obj[key]]
  }
}

for(const [key, value] of oentries(person)){
  console.log(`key: ${key}is ${value}`);
}
// key: nameis eriii
// key: phoneis 123-2312
```
