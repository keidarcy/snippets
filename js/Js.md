# Javascript Basic

- [javascript runtime simulation](http://latentflip.com/loupe)

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

## Set

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
