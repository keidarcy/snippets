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