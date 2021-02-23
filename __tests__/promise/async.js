const URL = 'https://jsonplaceholder.typicode.com/todos/1';

fetch(URL)
  .then((response) => response.json())
  .then((data) => console.log(data));

async function asyncCall() {
  const rawRes = await fetch(URL);
  const res = await rawRes.json();
  console.log(res);
  // expected output: "resolved"
}
try {
  asyncCall();
} catch (ex) {
  console.error(ex);
}

const async = (generator) => {
  const g = generator();
  (function next(value) {
    const n = g.next(value);
    if (n.done) return;
    n.value.then(next);
  })();
};

async(function* () {
  const response = yield fetch(URL);
  const data = yield response.json();
  console.log(data);
});

const fs = require('fs').promises;
const path = require('path');

async(function* () {
  const data = yield fs.readFile(path.resolve(__dirname, 'index.js'), 'utf-8');
  console.log({ data });
});
