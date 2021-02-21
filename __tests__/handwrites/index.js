const Promise = require('./promise');

const p = new Promise((resolve, reject) => {
  // throw new Error('nice');
  // setTimeout(() => {
  //   reject(100);
  // }, 500);
  setTimeout(() => {
    resolve(100);
    throw 111;
  }, 100);
})
  .then((data) => {
    console.log({ data1: data });
    return data + 1;
  })
  .then((data) => {
    console.log({ data2: data });
    return data;
  })
  .catch((err) => console.log({ err }));
