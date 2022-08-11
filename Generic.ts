// Generic Class
class KeyValuePair<K, V> {
  constructor(public key: K, public value: V) {}
}

let pair = new KeyValuePair<string, string>('two', 'one');

// Generic Functions
class ArrayUtils {
  static wrapInArray<T>(value: T) {
    return [value];
  }
}

ArrayUtils.wrapInArray(1);

// Generic Interfaces
interface Result<T> {
  data: T | null;
  error: string | null;
}

function fetchData<T>(url: string): Result<T> {
  return { data: null, error: null };
}

interface User {
  username: string;
}

interface Product {
  title: string;
}

const result = fetchData<Product>('url');
result.data?.title;

// Generic Constraints
class Person {
  constructor(public name: string) {}
}

class Customer extends Person {}
interface PersonInterface {
  lastname: string;
}
function echo<T extends PersonInterface | Person>(value: T): T {
  return value;
}

echo({ lastname: '1' });
echo(new Customer('1'));

// Extending Generic Classes
interface Product {
  name: string;
  price: number;
}

class Store<T> {
  protected _objects: T[] = [];

  add(object: T): void {
    this._objects.push(object);
  }

  find(property: keyof T, value: unknown): T | undefined {
    return this._objects.find((o) => o[property] === value);
  }
}

let store = new Store<Product>();
store.add({ name: '1', price: 1, title: '1' });
store.find('name', '1');
store.find('title', '1');

class CompressibleStore<T> extends Store<T> {
  // pass on the generic type parameter
  compress(): void {}
}

let cStore = new CompressibleStore<Product>();
cStore.compress();

class SearchableStore<T extends { name: string }> extends Store<T> {
  // restrict the generic type parameter
  override find(name: keyof T): T | undefined {
    return this._objects.find((o) => o.name === name);
  }
}

class ProductStore extends Store<Product> {
  // Fix the generic type parameter
  filterByCategory(category: string): Product[] {
    return [];
  }
}

// Type Mappings
type ReadOnly<T> = {
  readonly [Property in keyof T]: T[Property];
};

let product: ReadOnly<Product> = {
  name: '1',
  price: 1,
  title: '1',
};
// product.name = '2';

