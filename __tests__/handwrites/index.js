const Promise = require('./promise');

const p = new Promise((resolve, reject) => {
  // throw new Error('nice');
  setTimeout(() => {
    reject(100);
  }, 500);
})
  .then(
    (data) => {
      console.log({ data1: data });
      return data + 1;
    },
    (reason) => {
      console.log({ reason });
      throw 2000;
      return reason;
    }
  )
  .then(
    (data) => {
      console.log({ data2: data });
      return data;
    },
    (reason) => {
      console.log({ reason });
    }
  );
