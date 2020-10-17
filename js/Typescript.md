# Typescript basic

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
