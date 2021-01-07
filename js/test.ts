type IsNumber<T> = T extends number ? 'number' : 'other';

type WithNumber = IsNumber<number>;
type WithOther = IsNumber<boolean[]>;

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

type IsArray<T> = T extends Array<infer O> ? O : T;

type ItIsArray = IsArray<string[]>;
type ItIsNotArray = IsArray<boolean>;

const d = function neverExcute() {
  throw 'nothing';
};

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
  // const _bugTester: never = s;
}

const MyArray = [
  { name: 'Alice', age: 15 },
  { name: 'Bob', age: 23 },
  { name: 'Eve', age: 38 }
] as const;

type Ages = typeof MyArray[number]['age'];
